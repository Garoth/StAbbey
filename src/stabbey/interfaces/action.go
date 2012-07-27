package interfaces

type Action interface {
    /* Get the action type, relates to some const defines */
    ActionType() string
}
