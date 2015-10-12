package main

import (
    // "fmt"
    // "net/http"
    // "bufio"
    // "strconv"
    // "io"
    // "strings"
    // "runtime/pprof"
    "log"
    "flag"
    "github.com/addsict/qseq"
)

func main() {
    datadir := flag.String("datadir", "/tmp", "data directory")
    port    := flag.Int("port", 8080, "port number")
    flag.Parse()

    log.Printf("data directory: %s\n", *datadir)
    log.Printf("port number: %d\n", *port)

    h, _ := qseq.NewHandler(*datadir)

    d, _ := qseq.NewDispatcher(h)
    d.Run(*port)

    // http.HandleFunc("/sequence", func(w http.ResponseWriter, r *http.Request) {
        // g.ReqChan <- 1
        // nextSeq := <-g.ResChan
        // fmt.Fprintf(w, strconv.FormatUint(nextSeq, 10))
    // })

    // http.ListenAndServe(":9000", nil)
}
