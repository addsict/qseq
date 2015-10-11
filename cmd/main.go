package main

import (
    // "fmt"
    // "net/http"
    // "bufio"
    // "strconv"
    // "io"
    // "strings"
    // "runtime/pprof"
    "github.com/addsict/qseq"
)

func main() {
    g, _ := qseq.NewGenerator()
    go g.Run()

    d, _ := qseq.NewDispatcher(g)
    d.Run()

    // http.HandleFunc("/sequence", func(w http.ResponseWriter, r *http.Request) {
        // g.ReqChan <- 1
        // nextSeq := <-g.ResChan
        // fmt.Fprintf(w, strconv.FormatUint(nextSeq, 10))
    // })

    // http.ListenAndServe(":9000", nil)
}
