package stabbey

import (
    "fmt"
)

func RunPlayerTests(c *Context) {
    testId := 101010101

    fmt.Println("-> Basic save/load")
    p1 := NewPlayer(c, testId)
    p2 := LoadPlayer(c, testId)
    if p1.Equals(p2) {
        fmt.Println("   success")
    } else {
        fmt.Println("   failure")
    }
}
