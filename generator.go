package qseq

import (
    "fmt"
    "os"
    "strconv"
    "log"
    "io"
)

type Generator struct {
    ReqChan chan int
    ResChan chan uint64
}

func NewGenerator() (*Generator, error) {
    return &Generator{
        ReqChan: make(chan int, 100),
        ResChan: make(chan uint64, 100),
    }, nil
}

func (g *Generator) Run() {
    fh, err := os.OpenFile("seq_foo", os.O_RDWR, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer fh.Close()

    for {
        <-g.ReqChan
        g.ResChan <- GetNextSequence(fh)
    }
}

func GetNextSequence(fh *os.File) uint64 {
    b := make([]byte, 32)

    fh.Seek(0, 0)
    n, err := fh.Read(b)
    if err != nil && err != io.EOF {
        log.Fatal(err)
    }

    seq, err := strconv.ParseUint(string(b[:n]), 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    nextSeq := seq + 1

    fh.Seek(0, 0)
    _, err = fh.WriteString(fmt.Sprintf("%d", nextSeq))
    if err != nil {
        log.Fatal(err)
    }

    return nextSeq
}
