package main

import (
    "log"
    "flag"
    "github.com/addsict/qseq"
	"io/ioutil"
)

func main() {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	// command line options
    datadir := flag.String("datadir", tmpdir, "data directory")
    port    := flag.Int("port", 8080, "port number")
    flag.Parse()

    log.Printf("data directory: %s\n", *datadir)
    log.Printf("port number: %d\n", *port)

    h, _ := qseq.NewHandler(*datadir)

    d, _ := qseq.NewDispatcher(h)
    d.Run(*port)
}
