package stabbey

import (
    "strconv"
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
    http.HandleFunc("/", InitSetup)
    http.HandleFunc("/connect", ConnectSetup)
    http.HandleFunc("/update", UpdateRequest)
}

/* Serves the new game / connect page */
func InitSetup(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)

    setupTemplate, _ := template.ParseFiles("setup.html")
    if err := setupTemplate.Execute(w, map[string]string{}); err != nil {
        context.Errorf("Error executing setup template: %v", err)
    }
}

/* Creates game if necessary, connects players to it */
func ConnectSetup(w http.ResponseWriter, r *http.Request) {
    c       := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    newgame := c.Gamekey == ""
    game    := NewGame()
    player  := NewPlayer(user.Current(c.GAEContext).ID)

    if newgame {
        c.GAEContext.Infof("Making new game, %v", c.Gamekey)
        c.Gamekey = player.Id
        game.AddPlayer(player)
        board := NewBoard(c, "0")
        game.AddBoard(board)
        game.Save(c)
    } else {
        c.GAEContext.Infof("Game exists, adding player %v", player.Id)
        game := LoadGame(c);
        game.AddPlayer(player)
        game.Save(c)
    }

    player.Save(c)
    tok, _ := player.ChannelOpen(c)

    mainTemplate, _ := template.ParseFiles("main.html")
    mainTemplate.Execute(w, MainTemplate{player.Id, tok, c.Gamekey})
}

/* Replies to the clients' update requests by resending the game state */
// TODO validate client input here
func UpdateRequest(w http.ResponseWriter, r *http.Request) {
    c := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    ticknum, _ := strconv.Atoi(r.FormValue("ticknum"))
    PlayerUpdateLastTick(c, r.FormValue("player"), ticknum)
    RunGame(c)
}
