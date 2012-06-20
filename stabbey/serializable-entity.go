package stabbey

type SerializableEntity struct {
    EntityId, BoardId, X, Y int
    Name, Type string
}

func NewSerializableEntity(e Entity) *SerializableEntity {
    se := &SerializableEntity{}
    se.EntityId = e.GetEntityId();
    se.Name = e.GetName();
    se.Type = e.GetType();
    se.BoardId, se.X, se.Y = e.GetPosition();
    return se;
}
