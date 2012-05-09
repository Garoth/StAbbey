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
    c       := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    newgame := c.Gamekey == ""
    game    := NewGame()
    player  := NewPlayer(user.Current(c.GAEContext).ID)

    if newgame {
        fmt.Println("Making new game,", c.Gamekey)
        c.Gamekey = player.Id
        game.AddPlayer(player)
        board := NewBoard(0)
        board.Save(c)
        game.AddBoard(board)
        game.Save(c)
    } else {
        fmt.Println("Game exists, adding player", player.Id)
        game := LoadGame(c);
        game.AddPlayer(player)
        game.Save(c)
    }

    player.Save(c)
    tok, _ := player.OpenChannel(c)

    mainTemplate, _ := template.ParseFiles("main.html")
    mainTemplate.Execute(w, MainTemplate{player.Id, tok, c.Gamekey})
}

/* Replies to the clients' update requests by resending the game state */
func updateRequest(w http.ResponseWriter, r *http.Request) {
    c := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    game := LoadGame(c)

    for _, Id := range game.Players {
        LoadPlayer(c, Id).SendGamestate(c, game)
    }
}
