package qseq

import (
    "net/http"
    "fmt"
    // "log"
    "regexp"
    "strconv"
    "bytes"
)

type Dispatcher struct {
    handler *Handler
    routing map[string]*regexp.Regexp
}

func NewDispatcher(h *Handler) (*Dispatcher, error) {
    routing := map[string]*regexp.Regexp {
        "get": regexp.MustCompile("^/sequences/([^/]+)$"),
        "put": regexp.MustCompile("^/sequences/([^/]+)$"),
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

    case "PUT":
        d.HandlePut(w, r)
    
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

func (d *Dispatcher) HandlePut(w http.ResponseWriter, r *http.Request) {
    matched := d.routing["put"].FindStringSubmatch(r.RequestURI)
    if len(matched) > 1 {
        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        valStr := buf.String()

        value, err := strconv.ParseUint(valStr, 10, 64)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }

        seq, err := d.handler.PutSequence(matched[1], value)
        if err != nil {
            http.Error(w, err.Error(), 409)
            return
        }
        fmt.Fprintf(w, strconv.FormatUint(seq, 10))
        return
    }

    http.Error(w, "Invalid Parameter", 400)
}

func (d *Dispatcher) Run() {
    s := &http.Server{
        Addr: ":9000",
        Handler: d,
    }

    s.ListenAndServe()
}
