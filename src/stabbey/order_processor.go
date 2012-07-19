package main

import (
    "log"
    "stabbey/interfaces"
)

var OrderStream = make(chan interfaces.Order)

func ProcessOrders() {
    for {
        order := <-OrderStream
        log.Printf("Found Order: command:%v tick:%v actions:%v player:%v",
            order.GetCommandCode(), order.GetTickNumber(), order.GetActions(),
            order.GetPlayer().GetPlayerId())
    }
}
