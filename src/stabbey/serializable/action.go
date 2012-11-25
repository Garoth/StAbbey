package serializable

import (
    "stabbey/interfaces"
)

type Action struct {
    ActionString string
    AvailableDirections [5]bool
    ShortDescription string
    LongDescription string
}

func NewAction(a interfaces.Action) *Action {
    me := &Action{}

    /* Actions handed to the client are generally action "types", and not
     * specific instances to be used. Thus, we only supply the first character,
     * which is the type of action -- and not any secondary characters, which
     * are generally the direction the action is being used in */
    me.ActionString = string(a.ActionString()[0])
    me.AvailableDirections = a.AvailableDirections()
    me.ShortDescription = a.ShortDescription()
    me.LongDescription = a.LongDescription()

    return me
}
