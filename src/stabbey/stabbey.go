package main

import (
    "encoding/json"
    "flag"
    "log"
    "net/http"
    "strconv"
    "text/template"

    "code.google.com/p/go.net/websocket"

    "stabbey/board"
    "stabbey/game"
    "stabbey/interfaces"
    "stabbey/order"
    "stabbey/player"
    "stabbey/signalhandlers"
    "stabbey/runtime"
)

const FILE_SETUP_HTML string = "resources/html/setup.html"
const FILE_MAIN_HTML string  = "resources/html/main.html"
const HTTP_ROOT string       = "/"
const HTTP_CONNECT string    = "/connect"
const HTTP_WEBSOCKET string  = "/ws"
const FORMVAL_GAMEKEY string = "gamekey"

var ADDR = flag.String("addr", ":8080", "http service address")
var GAME *game.Game
var RUNTIME *runtime.Runtime

type MainPageTemplate struct {
    Me int
    Gamekey string
    Host string
}

/* Start the server and connect url paths to functions */
func main() {
    flag.Parse()

    go signalhandlers.Interrupt()
    go signalhandlers.Quit()

    http.HandleFunc("/resources/js/",       JavascriptHandler)
    http.HandleFunc("/resources/css/",      CssHandler)
    http.HandleFunc("/resources/img/",      ImgHandler)
    http.HandleFunc(HTTP_ROOT,    InitSetup)
    http.HandleFunc(HTTP_CONNECT, ConnectSetup)
    http.Handle(HTTP_WEBSOCKET,   websocket.Handler(WebSocketConnect))

    log.Println("Starting Server")
    if err := http.ListenAndServe(*ADDR, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

/* Send the basic page that just has the buttons to start/join games */
func InitSetup(w http.ResponseWriter, req *http.Request) {
    if tmpl, e := template.ParseFiles(FILE_SETUP_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, nil)
    }
}

/* Create the game, add players as they join */
func ConnectSetup(w http.ResponseWriter, r *http.Request) {
    gamekey := r.FormValue(FORMVAL_GAMEKEY)
    var curPlayer *player.Player

    /* New game! */
    if gamekey == "" && GAME == nil {
        gamekey = "0"
        GAME = game.NewGame(gamekey)
        GAME.AddBoard(board.New(0))
        GAME.Run()
    }

    curPlayer = player.New()
    pX, pY := GAME.GetBoard(0).GetRandomSpawnPoint()
    curPlayer.SetPosition(0, pX, pY)
    GAME.AddPlayer(curPlayer, curPlayer)
    log.Printf("Added player %v to game at %v, %v ", curPlayer.GetPlayerId(),
        pX, pY)

    if tmpl, e := template.ParseFiles(FILE_MAIN_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, MainPageTemplate{curPlayer.GetPlayerId(), gamekey,
            "ws://" + r.Host + "/ws"})
    }

    if RUNTIME == nil {
        RUNTIME = runtime.New(GAME)
    }
}

/* Initialize a websocket connection and pair it with the handshaking player */
func WebSocketConnect(ws *websocket.Conn) {
    websocketConnection := struct {
        CommandCode int
        Gamekey string
        Player string
    }{}

    var message string
    if err := websocket.Message.Receive(ws, &message); err != nil {
        log.Fatal("Reading Socket Error:", err)
    } else {
        err := json.Unmarshal([]byte(message), &websocketConnection)
        if err != nil {
            log.Fatal("Decoding Message Error:", err)
        } else {
            log.Printf("WebSocket Connected: %+v", websocketConnection)
            playerId, _ := strconv.Atoi(websocketConnection.Player)
            p := GAME.GetPlayer(playerId)
            p.SetWebSocketConnection(ws)
            KeepReading(p, ws)
        }
    }
}

/* Keep the player's websocket alive and continue reading from it forever */
func KeepReading(p interfaces.Player, ws *websocket.Conn) {
    for {
        var message string
        playerOrder := struct {
            CommandCode int
            TickNum int
            Queue []string
        }{}

        if e := websocket.Message.Receive(ws, &message); e != nil {
            break
        } else {
            err := json.Unmarshal([]byte(message), &playerOrder)
            if err != nil {
                log.Fatalf("Error Decoding Order: %v", e)
                break
            }

            o := order.NewOrder(playerOrder.CommandCode, playerOrder.TickNum,
                playerOrder.Queue, p)
            RUNTIME.AddOrder(o)
        }
    }
    log.Printf("Closing socket for %v", p.GetPlayerId())
    ws.Close()
}

/* Serve javascript files manually in order to set the content type */
func JavascriptHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/javascript")
    http.ServeFile(w, r, r.URL.Path[1:])
}

/* Serve javascript files manually in order to set the content type */
func CssHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css")
    http.ServeFile(w, r, r.URL.Path[1:])
}

/* Serve image files manually in order to set the content type */
func ImgHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/png")
    http.ServeFile(w, r, r.URL.Path[1:])
}
