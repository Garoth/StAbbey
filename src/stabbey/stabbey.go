package main

import (
    "flag"
    "log"
    "net/http"
    "text/template"

    "stabbey/constants"
    "stabbey/signalhandlers"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
    flag.Parse()

    go signalhandlers.Interrupt()
    go signalhandlers.Quit()

    http.HandleFunc(constants.HTTP_ROOT,    InitSetup)
    http.HandleFunc(constants.HTTP_CONNECT, ConnectSetup)
    http.HandleFunc(constants.HTTP_UPDATE,  UpdateRequest)
    http.HandleFunc(constants.HTTP_TEST,    RunTests)

    log.Println("Starting Server")
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
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

    if newgame {
        gamekey = "0"
        game := NewGame(gamekey)
        game.Run();
    } else {
    }
}

func UpdateRequest(w http.ResponseWriter, req *http.Request) {
}

func RunTests(w http.ResponseWriter, req *http.Request) {
}
