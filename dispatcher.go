package qseq

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Dispatcher struct {
	handler *Handler
	routing map[string]*regexp.Regexp
}

func NewDispatcher(h *Handler) (*Dispatcher, error) {
	routing := map[string]*regexp.Regexp{
		"get":    regexp.MustCompile(`^/sequences/([0-9a-zA-Z]+)`),
		"put":    regexp.MustCompile(`^/sequences/([0-9a-zA-Z]+)`),
		"delete": regexp.MustCompile(`^/sequences/([0-9a-zA-Z]+)`),
	}
	return &Dispatcher{
		handler: h,
		routing: routing,
	}, nil
}

func (d *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		d.HandleGet(w, r)

	case "PUT":
		d.HandlePut(w, r)

	case "DELETE":
		d.HandleDelete(w, r)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func (d *Dispatcher) HandleGet(w http.ResponseWriter, r *http.Request) {
	matched := d.routing["get"].FindStringSubmatch(r.RequestURI)
	if len(matched) > 1 {
		incrStr := r.FormValue("increment")
		var incr uint64 = 1
		var err error
		if incrStr != "" {
			incr, err = strconv.ParseUint(incrStr, 10, 64)
			if err != nil {
				http.Error(w, "invalid parameter", 400)
				return
			}
		}
		nextSeq, err := d.handler.GetSequence(matched[1], incr)
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

func (d *Dispatcher) HandleDelete(w http.ResponseWriter, r *http.Request) {
	matched := d.routing["delete"].FindStringSubmatch(r.RequestURI)
	if len(matched) > 1 {
		err := d.handler.DeleteSequence(matched[1])
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(204)
		return
	}

	http.Error(w, "Invalid Parameter", 400)
}

func (d *Dispatcher) Run(port int) {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: d,
	}

	log.Fatal(s.ListenAndServe())
}
