package qseq

import (
    "os"
    "fmt"
)

type Handler struct {
    seqFiles map[string]*os.File
    generator *Generator
}

func NewHandler() (*Handler, error) {
    g, _ := NewGenerator()
    go g.Run()

    return &Handler {
        seqFiles: map[string]*os.File{},
        generator: g,
    }, nil
}

func (h *Handler) HandleGetSequence(key string, incr uint32) (uint64, error) {
    fh := h.seqFiles[key]
    if fh == nil {
        var err error
        fh, err = os.OpenFile(key, os.O_RDWR, 0666)
        if err != nil {
            return 0, fmt.Errorf("key %s not found", key)
        }
        // defer fh.Close()

        h.seqFiles[key] = fh
    }

    h.generator.ReqChan <- fh
    seq := <-h.generator.ResChan
    return seq, nil
}

func (h *Handler) PutSequence(key string, value uint64) (uint64, error) {
    fh, err := os.Create(key)
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
