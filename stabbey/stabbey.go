package hello

import (
    "fmt"
    "http"
    "io/ioutil"
)

const (
    UI_FILE = "ui.html"
)

func init() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/rpc", rpcHandler)
}

// TODO there's a proper way to serve a file lol
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    s, e := ioutil.ReadFile(UI_FILE)
    if e != nil {
        fmt.Println(e)
        return
    }
    fmt.Fprint(w, string(s))
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "{ 'foo': 'bar' }")
}
