package stabbey

import (
    "fmt"
    "appengine"
    "appengine/datastore"
)

func Move(c *Context, entity Entity, command string) {
    bid, x, y := entity.GetPosition()
    if command == "mr" {
        entity.SetPosition(bid, x + 1, y)
    } else if command == "ml" {
        entity.SetPosition(bid, x - 1, y)
    } else if command == "mu" {
        entity.SetPosition(bid, x, y - 1)
    } else if command == "md" {
        entity.SetPosition(bid, x, y + 1)
    }
}

/* THIS IS THE MAIN ENTRY POINT FOR THE RUNNING GAME
 * This function will update the game state as per the
 * game mechanics, update objects, tell clients about updates, etc.
 */
func RunGame(c *Context, commandcode, playerId, ticknum int, queue []string) {

    if commandcode == 1 {
        PlayerUpdateLastTick(c, playerId, ticknum)
    } else if commandcode == 2 {
        // TODO figure out how to run this in a transaction well
        p := LoadPlayer(c, playerId)
        Move(c, p, queue[0])
        p.EntityStruct.Save(c);
        p.Save(c);
    }

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
            GameUpdateLastTick(c)
            game := LoadGame(c)
            fmt.Println("Everyone's up to date, send next tick", game.LastTick)
            for _, playerID := range game.Players {
                LoadPlayer(c, playerID).ChannelSendGame(c, game);
            }
        }

        return nil // TODO
    }, nil)
}
