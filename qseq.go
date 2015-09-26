package main

import (
    "fmt"
    "net/http"
    "os"
    "bufio"
    "strconv"
    "log"
    "io"
    "strings"
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
    fh.Seek(0, 0)
    reader := bufio.NewReader(fh)
    line, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        log.Fatal(err)
    }

    seq, err := strconv.ParseUint(strings.Trim(line, "\n"), 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    nextSeq := seq + 1

    fh.Truncate(0)
    fh.Seek(0, 0)
    writer := bufio.NewWriter(fh)
    _, err = writer.WriteString(fmt.Sprintf("%d\n", nextSeq))
    if err != nil {
        log.Fatal(err)
    }
    writer.Flush()

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
