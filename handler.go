package qseq

import (
    "os"
    "fmt"
)

type Handler struct {
    seqFiles map[string]*os.File
    generator *Generator
}

func NewHandler(g *Generator) (*Handler, error) {
    return &Handler {
        seqFiles: map[string]*os.File{},
        generator: g,
    }, nil;
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
