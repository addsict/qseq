package qseq

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Sequencer struct {
	ReqChan chan *os.File
	ResChan chan uint64
}

func NewSequencer() (*Sequencer, error) {
	return &Sequencer{
		ReqChan: make(chan *os.File, 100),
		ResChan: make(chan uint64, 100),
	}, nil
}

func (s *Sequencer) Run() {
	for {
		fh := <-s.ReqChan
		s.ResChan <- GetNextSequence(fh)
	}
}

func GetNextSequence(fh *os.File) uint64 {
	b := make([]byte, 32)

	fh.Seek(0, 0)
	n, err := fh.Read(b)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	seq, err := strconv.ParseUint(string(b[:n]), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	nextSeq := seq + 1

	fh.Seek(0, 0)
	_, err = fh.WriteString(fmt.Sprintf("%d", nextSeq))
	if err != nil {
		log.Fatal(err)
	}

	return nextSeq
}
