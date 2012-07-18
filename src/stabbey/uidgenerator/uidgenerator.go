package uidgenerator

type UidGenerator struct {
    /* Communications Channel */
    channel chan int
}

func New() *UidGenerator {
    uidg := &UidGenerator{make(chan int)}
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
