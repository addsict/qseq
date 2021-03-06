package qseq

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Handler struct {
	datadir   string
	seqFiles  map[string]*os.File
	sequencer *Sequencer
}

func NewHandler(datadir string) (*Handler, error) {
	s, _ := NewSequencer()
	go s.Run()

	return &Handler{
		datadir:   datadir,
		seqFiles:  map[string]*os.File{},
		sequencer: s,
	}, nil
}

func (h *Handler) GetSequence(key string, incr uint64) (uint64, error) {
	fh := h.seqFiles[key]
	if fh == nil {
		var err error
		abspath := h.getAbsPath(key)
		log.Printf("open the file: %s\n", abspath)
		fh, err = os.OpenFile(abspath, os.O_RDWR, 0666)
		if err != nil {
			return 0, fmt.Errorf("sequence %s not found", key)
		}

		// keep fh to serve later request
		h.seqFiles[key] = fh
	}

	seqReq := &SequenceRequest{
		fh: fh,
		incr: incr,
	}
	h.sequencer.ReqChan <- seqReq

	seq := <-h.sequencer.ResChan

	return seq, nil
}

func (h *Handler) PutSequence(key string, value uint64) (uint64, error) {
	absPath := h.getAbsPath(key)
	log.Printf("create a file: %s\n", absPath)

	fh, err := os.Create(absPath)
	if err != nil {
		return 0, err
	}

	_, err = fh.WriteString(fmt.Sprintf("%d", value))
	if err != nil {
		return 0, err
	}

	// keep fh to serve later request
	h.seqFiles[key] = fh

	return value, nil
}

func (h *Handler) DeleteSequence(key string) error {
	absPath := h.getAbsPath(key)
	log.Printf("remove the file: %s\n", absPath)
	err := os.Remove(absPath)
	if err != nil {
		return err
	}

	delete(h.seqFiles, key)

	return nil
}

func (h *Handler) getAbsPath(path string) string {
	absdir, _ := filepath.Abs(h.datadir)
	return filepath.Join(absdir, path)
}
