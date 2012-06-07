package stabbey

import (
    "appengine"
    "appengine/datastore"
)

type UidGenerator struct {
    /* Unique name for the UID Generator, for database storage */
    Name string
    /* The last UId this generator has handed out */
    LastUid int
}

func NewUidGenerator(c *Context, name string) *UidGenerator {
    uidg := &UidGenerator{name, 0}
    uidg.Save(c);
    return uidg
}

func GetUidGeneratorKey(c *Context, name string) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "UidGen", name, 0, nil)
}

func (uidg *UidGenerator) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetUidGeneratorKey(c, uidg.Name), uidg)

    if e != nil {
        c.GAEContext.Errorf("Error saving UID Generator: %v", e)
    } else {
        c.GAEContext.Infof("Successfully saved UID Generator")
    }

    return e;
}

func LoadUidGenerator(c *Context, name string) (*UidGenerator, error) {
    uidg := &UidGenerator{}
    e := datastore.Get(c.GAEContext, GetUidGeneratorKey(c, name), uidg)

    if e != nil {
        return uidg, e
    }

    c.GAEContext.Infof("Successfully loaded UID Generator")
    return uidg, nil
}

/* Tries to load the given uid generator (by name), but if it doesn't exist
 * in the database, make a new one. This is useful for situations where you
 * don't know whether a UID generator has existed before. For example, if
 * you're starting a new game, and want a unique gamekey, it could be that
 * a gamekey generator has never existed, so you want a new one. But most
 * likely, one has already been put in the database, so you want to get that.
 */
func GetUidGenerator(c *Context, name string) *UidGenerator {
    uidg, e := LoadUidGenerator(c, name)
    if e == datastore.ErrNoSuchEntity {
        return NewUidGenerator(c, name)
    } else if e != nil {
        c.GAEContext.Errorf("Serious datastore error on uid generator: %v", e)
    }
    return uidg
}

/* Hands out a new UID and updates the database */
func UidGeneratorGetUid(c *Context, name string) int {
    var ret int

    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        uidg, e := LoadUidGenerator(c, name)
        uidg.LastUid += 1
        ret = uidg.LastUid
        uidg.Save(c)
        return e
    }, nil)

    return ret
}
