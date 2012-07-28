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
    "stabbey/constants"
    "stabbey/game"
    "stabbey/interfaces"
    "stabbey/order"
    "stabbey/player"
    "stabbey/signalhandlers"
)

var ADDR = flag.String("addr", ":8080", "http service address")
/* TODO ok it's not safe to have a global like this */
var GAME *game.Game

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
    http.HandleFunc(constants.HTTP_ROOT,    InitSetup)
    http.HandleFunc(constants.HTTP_CONNECT, ConnectSetup)
    http.HandleFunc(constants.HTTP_TEST,    RunTests)
    http.Handle(constants.HTTP_WEBSOCKET,   websocket.Handler(WebSocketConnect))

    log.Println("Starting Server")
    if err := http.ListenAndServe(*ADDR, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

/* Send the basic page that just has the buttons to start/join games */
func InitSetup(w http.ResponseWriter, req *http.Request) {
    if tmpl, e := template.ParseFiles(constants.FILE_SETUP_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, nil)
    }
}

/* Create the game, add players as they join */
func ConnectSetup(w http.ResponseWriter, r *http.Request) {
    gamekey := r.FormValue(constants.FORMVAL_GAMEKEY)
    var curPlayer *player.Player

    /* New game! */
    if gamekey == "" && GAME == nil {
        gamekey = "0"
        GAME = game.NewGame(gamekey)
        GAME.AddBoard(board.New(0))
        GAME.Run();
        go ProcessOrders();
    }

    curPlayer = player.New()
    GAME.AddPlayer(curPlayer, curPlayer)
    log.Printf("Added player %v to game", curPlayer.GetPlayerId())

    if tmpl, e := template.ParseFiles(constants.FILE_MAIN_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, MainPageTemplate{curPlayer.GetPlayerId(), gamekey,
            "ws://" + r.Host + "/ws"})
    }
}

/* Initialize a websocket connection and pair it with the handshaking player */
func WebSocketConnect(ws *websocket.Conn) {
    websocketConnection := struct {
        CommandCode int
        Gamekey string
        Player string
    }{}

    var message string;
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
            ORDER_STREAM <- o
        }
    }
    log.Printf("Closing socket for %v", p.GetPlayerId())
    ws.Close()
}

/* Runs unit tests (lolno) */
func RunTests(w http.ResponseWriter, req *http.Request) {
}

/* Serve javascript files manually in order to set the content type */
func JavascriptHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/javascript");
    http.ServeFile(w, r, r.URL.Path[1:])
}

/* Serve javascript files manually in order to set the content type */
func CssHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css");
    http.ServeFile(w, r, r.URL.Path[1:])
}
