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

    if newgame {
        gamekey = user.Id
        fmt.Println("Making new game,", gamekey)
    }

    err := datastore.RunInTransaction(context, func(context appengine.Context) os.Error {
        key := datastore.NewKey(context, "Game", gamekey, 0, nil)

        game := new(Game)
        if newgame {
            game.P1 = user.Id
        } else {
            /* Load game, or error if we can't find it */
            if err := datastore.Get(context, key, game); err != nil {
                return err
            }

            /* Both players are in the game already */
            if game.P2 != "" {
                return nil
            }

            /* This is second player connecting, add them */
            if game.P1 != user.Id {
                game.P2 = user.Id
            }
        }

        /* Store the created or updated Game to the Datastore */
        _, err := datastore.Put(context, key, game)
        return err
    }, nil)

    if err != nil {
        http.Error(w, "Couldn't load Game", http.StatusInternalServerError)
        context.Errorf("setting up: %v", err)
        return
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
    for _, uId := range []string{game.P1, game.P2} {
        fmt.Println("Sending msg to user", uId, "channel is", uId + gamekey)
        err = channel.SendJSON(context, uId + gamekey, game)
        if err != nil {
            context.Errorf("sending Game: %v", err)
        }
        fmt.Println("done")
    }
}
