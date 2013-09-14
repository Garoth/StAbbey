/* Functions to configure the basics of different entity types */
package entity

import (
    "strconv"
    "stabbey/interfaces"
)

/* These are AI-controlled entities */
func newBasicMonster(g interfaces.Game) *Entity {
    me := New(UIDG.NextUid(), g)
    me.SetType(interfaces.ENTITY_TYPE_MONSTER)
    me.SetName("Monster " + strconv.Itoa(me.GetEntityId()))
    return me
}

/* These are things you step on, like traps, pickups, etc */
func newBasicTrigger(g interfaces.Game) *Entity {
    me := New(UIDG.NextUid(), g)
    me.SetType(interfaces.ENTITY_TYPE_TRIGGER)
    me.SetName("Trigger " + strconv.Itoa(me.GetEntityId()))
    me.SetTangible(false)
    return me
}

/* Random inert entities that do nothing */
func newBasicInert(g interfaces.Game) *Entity {
    me := New(UIDG.NextUid(), g)
    me.SetType(interfaces.ENTITY_TYPE_INERT)
    me.SetName("Inert " + strconv.Itoa(me.GetEntityId()))
    return me
}
