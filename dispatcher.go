package qseq

import (
    "net/http"
    "fmt"
    "log"
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
        "get": regexp.MustCompile(`^/sequences/([0-9a-zA-Z]+)$`),
        "put": regexp.MustCompile(`^/sequences/([0-9a-zA-Z]+)$`),
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

        var value uint64 = 0
        var err error
        if valStr != "" {
            value, err = strconv.ParseUint(valStr, 10, 64)
            if err != nil {
                http.Error(w, err.Error(), 400)
                return
            }
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

func (d *Dispatcher) Run(port int) {
    s := &http.Server{
        Addr: fmt.Sprintf(":%d", port),
        Handler: d,
    }

    log.Fatal(s.ListenAndServe())
}
