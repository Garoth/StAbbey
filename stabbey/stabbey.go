package stabbey

import (
    "fmt"
    "net/http"
    "html/template"
    "appengine"
    "appengine/user"
)

type MainTemplate struct {
    Me string
    Token string
    Gamekey string
}

/* Configures what virtual urls map to what functions */
func init() {
    http.HandleFunc("/", initSetup)
    http.HandleFunc("/connect", connectSetup)
    http.HandleFunc("/update", updateRequest)
}

/* Serves the new game / connect page */
func initSetup(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)

    setupTemplate, _ := template.ParseFiles("setup.html")
    if err := setupTemplate.Execute(w, map[string]string{}); err != nil {
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
    player  := NewPlayer(user.ID)

    if newgame {
        fmt.Println("Making new game,", gamekey)
        gamekey = user.ID
        game.AddPlayer(player)
        board := NewBoard(0)
        board.MakeTestBoard()
        board.Save(context, gamekey)
        game.AddBoard(board)
        game.Save(context, gamekey)
    } else {
        fmt.Println("Game exists, adding player", user.ID)
        game.Load(context, gamekey)
        game.AddPlayer(player)
        game.Save(context, gamekey)
    }

    player.Save(context, gamekey)
    tok, _ := player.OpenChannel(context, gamekey)

    mainTemplate, _ := template.ParseFiles("main.html")
    err := mainTemplate.Execute(w, MainTemplate{user.ID, tok, gamekey})

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
        p.SendGamestate(context, gamekey, game)
    }
}
