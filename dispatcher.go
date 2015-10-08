package qseq

import (
    "net/http"
    "fmt"
    // "log"
    "regexp"
)

type Dispatcher struct {
}

func NewDispatcher() (*Dispatcher, error) {

    return &Dispatcher{
    }, nil;
}

func (d *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        // r.URL
        rgx := regexp.MustCompile("^/sequences/([^/]+)$") // あくまでも GET /sequences{/id} にmatchするかどうかをチェックする
        matched := rgx.FindStringSubmatch(r.RequestURI)
        if len(matched) > 1 {
            fmt.Fprintf(w, matched[1])
        } else {
            fmt.Fprintf(w, "hello")
        }
        // log.Println(r.RequestURI)
    
    default:
        http.Error(w, "Method Not Allowed", 405)
    }
}

func (d *Dispatcher) Run() {
    s := &http.Server{
        Addr: ":9000",
        Handler: d,
    }

    s.ListenAndServe()
}
