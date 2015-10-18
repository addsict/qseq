package qseq

import (
	"io/ioutil"
	"testing"
	"fmt"
	"github.com/addsict/qseq"
	"github.com/lestrrat/go-tcptest"
	"time"
	"net/http"
	"strings"
)

func TestServeHTTP(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Failed to create a temporary data directory")
	}

	h, _ := qseq.NewHandler(tempdir)
	d, _ := qseq.NewDispatcher(h)

	server, err := tcptest.Start2(func(tt *tcptest.TCPTest) {
		port := tt.Port()
		d.Run(port)
	}, time.Second * 30)
	if err != nil {
		t.Fatalf("Failed to start the server")
	}

	t.Logf("Starting server... port = %d", server.Port())

	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d/sequences/foo", server.Port()), nil)
	res, err := client.Do(req)
	if res.StatusCode != 404 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:%d/sequences/foo", server.Port()), nil)
	res, err = client.Do(req)
	if res.StatusCode != 200 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body);
	if err != nil {
		t.Errorf("Invalid response")
	}
	if respBody := string(b); respBody != "0" {
		t.Errorf("Invalid response: %s", respBody)
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:%d/sequences/foo", server.Port()), strings.NewReader("100"))
	res, err = client.Do(req)
	if res.StatusCode != 200 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}
	b, err = ioutil.ReadAll(res.Body);
	if err != nil {
		t.Errorf("Invalid response")
	}
	if respBody := string(b); respBody != "100" {
		t.Errorf("Invalid response: %s", respBody)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d/sequences/foo", server.Port()), nil)
	res, err = client.Do(req)
	if res.StatusCode != 200 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}
	b, err = ioutil.ReadAll(res.Body);
	if err != nil {
		t.Errorf("Invalid response")
	}
	if respBody := string(b); respBody != "101" {
		t.Errorf("Invalid response: %s", respBody)
	}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:%d/sequences/foo", server.Port()), nil)
	res, err = client.Do(req)
	if res.StatusCode != 204 {
		t.Errorf("Invalid status code: %d", res.StatusCode)
	}
}
