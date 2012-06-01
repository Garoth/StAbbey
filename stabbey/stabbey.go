package stabbey

import (
    "strconv"
    "net/http"
    "html/template"
    "appengine"
    "appengine/user"
)

type MainTemplate struct {
    Me int
    Token string
    Gamekey string
}

/* Configures what virtual urls map to what functions */
func init() {
    http.HandleFunc("/", InitSetup)
    http.HandleFunc("/connect", ConnectSetup)
    http.HandleFunc("/update", UpdateRequest)
    http.HandleFunc("/tests", RunTests)
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
    c := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    newgame := c.Gamekey == ""
    playerId, _ := strconv.Atoi(user.Current(c.GAEContext).ID[0:9])

    if newgame {
        c.Gamekey = strconv.Itoa(playerId)
        player := NewPlayer(c, playerId)
        NewGame(c)
        GameAddPlayer(c, player)
        GameAddBoard(c, NewBoard(c, 0))
        tok, _ := player.ChannelOpen(c)
        mainTemplate, _ := template.ParseFiles("main.html")
        mainTemplate.Execute(w, MainTemplate{player.Id, tok, c.Gamekey})
    } else {
        player := NewPlayer(c, playerId)
        c.GAEContext.Infof("Game exists, adding player %v", player.Id)
        GameAddPlayer(c, player)
        tok, _ := player.ChannelOpen(c)
        mainTemplate, _ := template.ParseFiles("main.html")
        mainTemplate.Execute(w, MainTemplate{player.Id, tok, c.Gamekey})
    }
}

/* Replies to the clients' update requests by resending the game state */
// TODO validate client input here
func UpdateRequest(w http.ResponseWriter, r *http.Request) {
    c := NewContext(appengine.NewContext(r), r.FormValue("gamekey"))
    ticknum, _ := strconv.Atoi(r.FormValue("ticknum"))
    playerId, _ := strconv.Atoi(r.FormValue("player"))
    PlayerUpdateLastTick(c, playerId, ticknum)
    RunGame(c)
}

/* Runs unit tests */
func RunTests(w http.ResponseWriter, r *http.Request) {
    c := NewContext(appengine.NewContext(r), "TEST")
    RunPlayerTests(c)
}
