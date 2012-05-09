package stabbey

import (
    "appengine"
)

type Context struct {
    GAEContext appengine.Context
    Gamekey string
}

func NewContext(context appengine.Context, gamekey string) *Context {
    return &Context{context, gamekey}
}
