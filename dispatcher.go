package qseq

import (
    "net/http"
    "fmt"
    // "log"
    "regexp"
    "strconv"
)

type Dispatcher struct {
    handler *Handler
    routing map[string]*regexp.Regexp
}

func NewDispatcher(h *Handler) (*Dispatcher, error) {
    routing := map[string]*regexp.Regexp {
        "get": regexp.MustCompile("^/sequences/([^/]+)$"),
    }
    return &Dispatcher{
        handler: h,
        routing: routing,
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
    matched := d.routing["get"].FindStringSubmatch(r.RequestURI)
    if len(matched) > 1 {
        // incr := r.FormValue("increment")
        nextSeq, err := d.handler.HandleGetSequence(matched[1], 1)
        if err != nil {
            http.Error(w, err.Error(), 404)
            return
        }
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
