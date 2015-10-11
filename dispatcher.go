package qseq

import (
    "net/http"
    "fmt"
    // "log"
    "regexp"
    "strconv"
)

type Dispatcher struct {
    generator *Generator
}

func NewDispatcher(g *Generator) (*Dispatcher, error) {
    return &Dispatcher{
        generator: g,
    }, nil;
}

func (d *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        d.HandleGet(w, r)
    
    default:
        http.Error(w, "Method Not Allowed", 405)
    }
}

func (d *Dispatcher) HandleGet(w http.ResponseWriter, r *http.Request) {
    rgx := regexp.MustCompile("^/sequences/([^/]+)$")
    matched := rgx.FindStringSubmatch(r.RequestURI)
    if len(matched) > 1 {
        d.generator.ReqChan <- matched[1]
        nextSeq := <-d.generator.ResChan
        fmt.Fprintf(w, strconv.FormatUint(nextSeq, 10))
        return
    }

    http.Error(w, "Not Found", 404)
}

func (d *Dispatcher) Run() {
    s := &http.Server{
        Addr: ":9000",
        Handler: d,
    }

    s.ListenAndServe()
}
