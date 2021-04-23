package kanye

import (
	"testing"
)

func TestKanyeAPI(t *testing.T) {
	c := NewClient()
	q, err := c.GetQuote()

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if q == "" {
		t.Errorf("got empty quote: %v", q)
	}
}
