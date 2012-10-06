package monsters

import (
    "strconv"

    "stabbey/entity"
    "stabbey/uidgenerator"
    "stabbey/interfaces"
)

var uidg = uidgenerator.New();

type Monster struct {
    entity.Entity
    MonsterId int
}

/* Generates a particular type of monster using a monster builder */
func New(monsterBuilder func(*Monster)) *Monster {
    monster := newGeneric()
    monsterBuilder(monster)
    return monster
}

/* Basic initiation of a monster class */
func newGeneric() *Monster {
    me := &Monster{}

    /* Player stuff */
    me.MonsterId = uidg.NextUid();

    /* Entity stuff */
    me.SetPosition(0, 8, 6)
    me.SetEntityId(entity.UIDG.NextUid())
    me.SetType(interfaces.ENTITY_TYPE_MONSTER)
    me.SetName("Monster " + strconv.Itoa(me.MonsterId))

    return me
}

func (me *Monster) GetMonsterId() int {
    return me.MonsterId
}

func (me *Monster) SetMonsterId(id int) {
    me.MonsterId = id
}
