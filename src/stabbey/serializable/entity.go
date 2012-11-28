package serializable

import (
    "stabbey/interfaces"
)

type Entity struct {
    EntityId, BoardId, X, Y, MaxArdour, Ardour int
    Name, Type string
    ActionQueue []string
}

func NewEntity(e interfaces.Entity) *Entity {
    se := &Entity{}
    se.EntityId = e.GetEntityId()
    se.Name = e.GetName()
    se.Type = e.GetType()
    se.MaxArdour = e.GetMaxArdour()
    se.Ardour = e.GetArdour()
    se.BoardId, se.X, se.Y = e.GetPosition()
    se.ActionQueue = e.GetStringActionQueue()
    return se
}
