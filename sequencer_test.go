package qseq

import (
	"io/ioutil"
	"testing"
	"math"
	"fmt"
	"github.com/addsict/qseq"
)

var nextSeqTests = []struct {
	current uint64
	increment uint64
	next uint64
}{
	{0, 1, 1},
	{100, 1, 101},
	{uint64(math.Pow(2, 31)) - 1, 1, uint64(math.Pow(2, 31))},
	{uint64(math.Pow(2, 32)) - 1, 1, uint64(math.Pow(2, 32))},
	{uint64(math.Pow(2, 63)) - 1, 1, uint64(math.Pow(2, 63))},
}

func TestGetNextSequence(t *testing.T) {
	tempfh, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Failed to create a temporary file")
	}

	for _, s := range nextSeqTests {
		err := tempfh.Truncate(0)
		if err != nil {
			t.Fatalf("Failed to truncate a temporary file")
		}
		tempfh.Seek(0, 0)

		_, err = tempfh.WriteString(fmt.Sprintf("%d", s.current))
		if err != nil {
			t.Fatalf("Failed to write a content to the temporary file")
		}

		seq := qseq.GetNextSequence(tempfh, s.increment)
		if seq != s.next {
			t.Errorf("Next sequence is invalid: expected = %d, got = %d", s.next, seq)
		}
	}
}
