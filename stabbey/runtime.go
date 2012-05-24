package stabbey

import (
    "fmt"
    "appengine"
    "appengine/datastore"
)

/* THIS IS THE MAIN ENTRY POINT FOR THE RUNNING GAME
 * This function will update the game state as per the
 * game mechanics, update objects, tell clients about updates, etc.
 */
func RunGame(c *Context) {
    everyone_up_to_date := true

    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        game := LoadGame(c)

        for _, playerID := range game.Players {
            if (LoadPlayer(c, playerID).LastTick < game.LastTick + 1) {
                fmt.Println("Players aren't ready yet.")
                everyone_up_to_date = false
            }
        }

        return nil // TODO
    }, nil)

    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        if everyone_up_to_date {
            fmt.Println("Everyone's up to date, send next tick")
            GameUpdateLastTick(c)
            game := LoadGame(c)
            for _, playerID := range game.Players {
                LoadPlayer(c, playerID).ChannelSendGame(c, game);
            }
        }

        return nil // TODO
    }, nil)
}
