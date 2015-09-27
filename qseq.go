package main

import (
    "fmt"
    "net/http"
    "os"
    // "bufio"
    "strconv"
    "log"
    "io"
    // "strings"
    // "runtime/pprof"
)

// type Server struct {
    // generatorChan chan string
// }

// func NewServer() *Server {
    // g := 
    // s := &Server{
    // }
    // return s
// }

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

func main() {

    requestChan := make(chan int, 100)
    responseChan := make(chan uint64, 100)

    go func() {
        fh, err := os.OpenFile("seq_foo", os.O_RDWR, 0666)
        if err != nil {
            log.Fatal(err)
        }
        defer fh.Close()

        for ;; {
            <-requestChan
            responseChan <- GetNextSequence(fh)
        }
    }()

    http.HandleFunc("/sequence", func(w http.ResponseWriter, r *http.Request) {
        requestChan <- 1
        nextSeq := <-responseChan
        fmt.Fprintf(w, strconv.FormatUint(nextSeq, 10))
    })

    http.ListenAndServe(":9000", nil)
}
