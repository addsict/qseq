package qseq

import (
    "os"
    "fmt"
    "log"
    "path/filepath"
)

type Handler struct {
    datadir string
    seqFiles map[string]*os.File
    generator *Generator
}

func NewHandler(datadir string) (*Handler, error) {
    g, _ := NewGenerator()
    go g.Run()

    return &Handler {
        datadir: datadir,
        seqFiles: map[string]*os.File{},
        generator: g,
    }, nil
}

func (h *Handler) HandleGetSequence(key string, incr uint32) (uint64, error) {
    fh := h.seqFiles[key]
    if fh == nil {
        var err error
        abspath := h.getAbsPath(key)
        log.Printf("open file: %s\n", abspath)
        fh, err = os.OpenFile(abspath, os.O_RDWR, 0666)
        if err != nil {
            return 0, fmt.Errorf("sequence %s not found", key)
        }
        // defer fh.Close()

        h.seqFiles[key] = fh
    }

    h.generator.ReqChan <- fh
    seq := <-h.generator.ResChan
    return seq, nil
}

// func (h *Handler) GetAllSequences() ([]string, error) {
// }

func (h *Handler) PutSequence(key string, value uint64) (uint64, error) {
    absPath := h.getAbsPath(key)
    log.Printf("create file: %s\n", absPath)

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
    log.Printf("remove file: %s\n", absPath)
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
