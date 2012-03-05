/* Barebones server */
package stabbey

import (
    "os"
    "fmt"
    "http"
    "old/template"
    "appengine"
//    "appengine/user"
    "appengine/datastore"
    "appengine/channel"
)

var MAIN_TEMPLATE = template.MustParseFile("main.html", nil)

func init() {
    http.HandleFunc("/", initSetup)
    http.HandleFunc("/connect", connectSetup)
}

func initSetup(w http.ResponseWriter, r *http.Request) {
}

func connectSetup(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Starting up server")

    context := appengine.NewContext(r)
    //user := user.Current(context)
    gamekey := r.FormValue("gamekey")

    err := datastore.RunInTransaction(context, func(c appengine.Context) os.Error {
        key := datastore.NewKey(context, "Game", gamekey, 0, nil)
        gamekey = key.String();

        game := new(Game)
        /* TODO game detection here */
        if true {
            // No game specified.
            game.p1 = Player{1}
        } else {
            // Game key specified, load it from the Datastore.
            if err := datastore.Get(context, key, game); err != nil {
                return err
            }
            if game.p2.id == 0 {
                game.p2 = Player{2}
            }
        }
        // Store the created or updated Game to the Datastore.
        _, err := datastore.Put(context, key, game)
        return err
    }, nil)

    if err != nil {
        http.Error(w, "Couldn't load Game", http.StatusInternalServerError)
        context.Errorf("setting up: %v", err)
        return
    }

    tok, err := channel.Create(context, "1" + gamekey)
    if err != nil {
        http.Error(w, "Couldn't create Channel", http.StatusInternalServerError)
        context.Errorf("channel.Create: %v", err)
        return
    }

    fmt.Println("gk: ", gamekey);
    err = MAIN_TEMPLATE.Execute(w, map[string]string{
        "token"   : tok,
        "me"      : "1",
        "gamekey" : gamekey,
    })
    if err != nil {
        context.Errorf("mainTemplate: %v", err)
    }
}
