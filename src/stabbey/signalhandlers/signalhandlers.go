package signalhandlers

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
)

func Interrupt() {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT)
    <-ch
    fmt.Println("")
    log.Println("CTRL-C (SIGINT); exiting")
    os.Exit(0)
}

func Quit() {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGQUIT)
    <-ch
    fmt.Println("")
    log.Println("CTRL-\\ (SIGQUIT); exiting")
    os.Exit(1)
}

