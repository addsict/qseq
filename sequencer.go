package qseq

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type SequenceRequest struct {
	fh *os.File
	incr uint64
}

type Sequencer struct {
	ReqChan chan *SequenceRequest
	ResChan chan uint64
}

func NewSequencer() (*Sequencer, error) {
	return &Sequencer{
		ReqChan: make(chan *SequenceRequest, 100),
		ResChan: make(chan uint64, 100),
	}, nil
}

func (s *Sequencer) Run() {
	for {
		req := <-s.ReqChan
		s.ResChan <- GetNextSequence(req.fh, req.incr)
	}
}

func GetNextSequence(fh *os.File, incr uint64) uint64 {
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

	nextSeq := seq + incr

	fh.Seek(0, 0)
	_, err = fh.WriteString(fmt.Sprintf("%d", nextSeq))
	if err != nil {
		log.Fatal(err)
	}

	return nextSeq
}
