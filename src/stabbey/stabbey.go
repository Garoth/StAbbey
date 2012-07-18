package main

import (
    "encoding/json"
    "flag"
    "log"
    "net/http"
    "strconv"
    "text/template"

    "code.google.com/p/go.net/websocket"

    "stabbey/constants"
    "stabbey/game"
    "stabbey/interfaces"
    "stabbey/player"
    "stabbey/signalhandlers"
)

var ADDR = flag.String("addr", ":8080", "http service address")
var GAME *game.Game

type MainPageTemplate struct {
    Me int
    Gamekey string
    Host string
}

func main() {
    flag.Parse()

    go signalhandlers.Interrupt()
    go signalhandlers.Quit()

    http.HandleFunc("/resources/js/",       JavascriptHandler)
    http.HandleFunc(constants.HTTP_ROOT,    InitSetup)
    http.HandleFunc(constants.HTTP_CONNECT, ConnectSetup)
    http.HandleFunc(constants.HTTP_TEST,    RunTests)
    http.Handle(constants.HTTP_WEBSOCKET,   websocket.Handler(WebSocketConnect))

    log.Println("Starting Server")
    if err := http.ListenAndServe(*ADDR, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func JavascriptHandler(w http.ResponseWriter, r *http.Request) {
    /* Serve javascript files manually in order to set the content type */
    w.Header().Set("Content-Type", "text/javascript");
    http.ServeFile(w, r, r.URL.Path[1:])
}

func InitSetup(w http.ResponseWriter, req *http.Request) {
    if tmpl, e := template.ParseFiles(constants.FILE_SETUP_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, nil)
    }
}

func ConnectSetup(w http.ResponseWriter, r *http.Request) {
    gamekey := r.FormValue(constants.FORMVAL_GAMEKEY)
    newgame := gamekey == ""
    var curPlayer interfaces.Player

    if newgame {
        gamekey = "0"
        GAME = game.NewGame(gamekey)
        GAME.Run();
        curPlayer = player.New()
        GAME.AddPlayer(curPlayer)
    } else {
    }

    if tmpl, e := template.ParseFiles(constants.FILE_MAIN_HTML); e != nil {
        log.Fatal("Parse error:", e)
    } else {
        tmpl.Execute(w, MainPageTemplate{curPlayer.GetPlayerId(), gamekey,
            "ws://" + r.Host + "/ws"})
    }
}

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
            log.Printf("Decode Successful: %+v", websocketConnection)
            playerId, _ := strconv.Atoi(websocketConnection.Player)
            p := GAME.GetPlayer(playerId)
            p.SetWebSocketConnection(ws)
        }
    }
}

func RunTests(w http.ResponseWriter, req *http.Request) {
}
