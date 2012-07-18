package uidgenerator

import (
)

type GeneratorType struct {
    Type string
}

func PlayerType() *GeneratorType {
    return &GeneratorType{"player"}
}

func (gt *GeneratorType) String() string {
    return gt.Type
}

type UidGenerator struct {
    /* Unique name for the UID Generator, for database storage */
    Type *GeneratorType
    /* Communications Channel */
    channel chan int
}

func New(gentype *GeneratorType) *UidGenerator {
    uidg := &UidGenerator{gentype, make(chan int)}
    go func() {
        for i := 0; ; i++ {
            uidg.channel <- i
        }
    }()
    return uidg
}

func (uidg *UidGenerator) NextUid() int {
    return <-uidg.channel
}
