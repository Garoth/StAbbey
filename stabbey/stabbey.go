/* Barebones server */
package stabbey

import (
    "fmt"
    "http"
    "io/ioutil"
)

const (
    UI_FILE = "ui.html"
    MAX_PLAYERS = 10
)

var NUM_PLAYERS = 0;
var PLAYERS = make([]Player, MAX_PLAYERS)

func init() {
    http.HandleFunc("/", initSetup)
    http.HandleFunc("/connect", playerSetup)
}

func initSetup(w http.ResponseWriter, r *http.Request) {
    s, e := ioutil.ReadFile(UI_FILE)
    if e != nil {
        fmt.Println(e)
        return
    }
    fmt.Fprint(w, string(s))
}

func playerSetup(w http.ResponseWriter, r *http.Request) {
    id := -1

    if (NUM_PLAYERS < len(PLAYERS)) {
        id = NUM_PLAYERS;
        NUM_PLAYERS++;
    }

    PLAYERS[id] = Player{id}
    fmt.Println(PLAYERS[id].id)

    fmt.Fprintf(w, "{ \"id\": \"%d\" }", PLAYERS[0].id)
}
