package interfaces

type Action interface {
    /* Get the action command string (ex. ml for move left) */
    ActionString() string
    /* Execute the action, whatever it is */
    Act(e Entity, g Game)
}
