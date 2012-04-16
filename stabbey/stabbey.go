/* Barebones server */
package stabbey

import (
    "os"
    "fmt"
    "http"
    "old/template"
    "appengine"
    "appengine/user"
    "appengine/datastore"
    "appengine/channel"
)

var SETUP_TEMPLATE = template.MustParseFile("setup.html", nil)
var MAIN_TEMPLATE  = template.MustParseFile("main.html", nil)

func init() {
    http.HandleFunc("/", initSetup)
    http.HandleFunc("/connect", connectSetup)
    http.HandleFunc("/update", updateRequest)
}

func initSetup(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)

    err := SETUP_TEMPLATE.Execute(w, map[string]string{})

    if err != nil {
        context.Errorf("setupTemplate: %v", err)
    }
}

func connectSetup(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Starting up server")

    context := appengine.NewContext(r)
    user := user.Current(context)
    gamekey := r.FormValue("gamekey")
    newgame := gamekey == ""
    game := NewGame()

    if newgame {
        gamekey = user.Id
        fmt.Println("Making new game,", gamekey)
        game.AddPlayer(NewPlayer(user.Id))
        game.Save(context, gamekey)
    } else {
        game.Load(context, gamekey)
        fmt.Println("Game exists, adding player", user.Id)
        game.AddPlayer(NewPlayer(user.Id))
        game.Save(context, gamekey)
    }

    fmt.Println("Making channel of:", user.Id + gamekey)
    tok, err := channel.Create(context, user.Id + gamekey)
    if err != nil {
        http.Error(w, "Couldn't create Channel", http.StatusInternalServerError)
        context.Errorf("channel.Create: %v", err)
        return
    }

    err = MAIN_TEMPLATE.Execute(w, map[string]string{
        "token"   : tok,
        "me"      : user.Id,
        "gamekey" : gamekey,
    })

    if err != nil {
        context.Errorf("mainTemplate: %v", err)
    }
}

func updateRequest(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    //user := user.Current(context)
    gamekey := r.FormValue("gamekey")
    fmt.Println("updateRequest gamekey:", gamekey);

    game := new(Game)
    err := datastore.RunInTransaction(context, func(context appengine.Context) os.Error {

        /* Retrieve the game from the Datastore */
        key := datastore.NewKey(context, "Game", gamekey, 0, nil)
        if err := datastore.Get(context, key, game); err != nil {
            return err
        }

        /* TODO do something with the game */

        /* Update the Datastore. */
        _, err := datastore.Put(context, key, game)
        return err
    }, nil)

    if err != nil {
        http.Error(w, "Couldn't handle update", http.StatusInternalServerError)
        context.Errorf("updateRequest: %v", err)
        return
    }

    // Send the game state to both clients.
    for _, Id := range game.Players {
        p := NewPlayer(Id)
        e := p.Load(context, gamekey)

        if e != nil {
            fmt.Println("Error in getting player", e,
            "Id:", Id,
            "context:", context)
        } else {
            fmt.Println("Player", Id, "loaded successfully!")
        }

        fmt.Println("Sending msg to user", p.Id, "channel is", p.Id + gamekey)
        err = channel.SendJSON(context, p.Id + gamekey, game)
        if err != nil {
            context.Errorf("sending Game: %v", err)
        }
        fmt.Println("done")
    }
}
