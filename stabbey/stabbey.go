/* Barebones server */
package stabbey

import (
    "fmt"
    "http"
    "old/template"
    "appengine"
    "appengine/user"
)

var SETUP_TEMPLATE = template.MustParseFile("setup.html", nil)
var MAIN_TEMPLATE  = template.MustParseFile("main.html", nil)

/* Configures what virtual urls map to what functions */
func init() {
    http.HandleFunc("/", initSetup)
    http.HandleFunc("/connect", connectSetup)
    http.HandleFunc("/update", updateRequest)
}

/* Serves the new game / connect page */
func initSetup(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)

    if err := SETUP_TEMPLATE.Execute(w, map[string]string{}); err != nil {
        context.Errorf("Error executing setup template: %v", err)
    }
}

/* Creates game if necessary, connects players to it */
func connectSetup(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    user    := user.Current(context)
    gamekey := r.FormValue("gamekey")
    newgame := gamekey == ""
    game    := NewGame()
    player  := NewPlayer(user.Id)

    if newgame {
        fmt.Println("Making new game,", gamekey)
        gamekey = user.Id
        game.AddPlayer(player)
        board := NewBoard(0)
        board.MakeTestBoard()
        board.Save(context, gamekey)
        game.AddBoard(board)
        game.Save(context, gamekey)
    } else {
        fmt.Println("Game exists, adding player", user.Id)
        game.Load(context, gamekey)
        game.AddPlayer(player)
        game.Save(context, gamekey)
    }

    player.Save(context, gamekey)
    tok, _ := player.OpenChannel(context, gamekey)

    err := MAIN_TEMPLATE.Execute(w, map[string]string{
        "token"   : tok,
        "me"      : user.Id,
        "gamekey" : gamekey,
    })

    if err != nil {
        context.Errorf("Error executing main template: %v", err)
    }
}

/* Replies to the clients' update requests by resending the game state */
func updateRequest(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    gamekey := r.FormValue("gamekey")

    game := new(Game)
    game.Load(context, gamekey)

    // Send the game state to both clients.
    for _, Id := range game.Players {
        p := NewPlayer(Id)
        p.Load(context, gamekey)
        p.SendJSON(context, gamekey, game)
    }
}
