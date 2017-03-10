package passphrase_test

import (
	"strings"
	"testing"

	"github.com/leonklingele/passphrase"
)

func TestPassphraseLength(t *testing.T) {
	ls := []int{1, 5, 10, 50, 100}

	for _, l := range ls {
		s, err := passphrase.Generate(l)
		if err != nil {
			t.Error(err)
		}

		split := strings.Split(s, passphrase.Separator)
		if len(split) != l {
			t.Errorf("unexpected length of passphrase string: want %d, got %d", l, len(split))
		}
	}
}

func TestPassphraseLengthCustomSeparator(t *testing.T) {
	passphrase.Separator = "_"
	TestPassphraseLength(t)
}

func TestPassphraseInvalidLength(t *testing.T) {
	ls := []int{-1, 0}

	for _, l := range ls {
		if _, err := passphrase.Generate(l); err != passphrase.ErrInvalidLengthSpecified {
			t.Errorf("unexpected error for length %d: %v", l, err)
		}
	}
}
