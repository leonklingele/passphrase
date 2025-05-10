package passphrase_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/leonklingele/passphrase"
)

func TestPassphraseLength(t *testing.T) {
	t.Parallel()

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

//nolint:paralleltest // Make this parallel once we no longer modify a global variable of the package
func TestPassphraseLengthCustomSeparator(t *testing.T) {
	passphrase.Separator = "_" //nolint:reassign // TODO: Get rid of global variable as it makes everything racy. Make this test parallel then.
	TestPassphraseLength(t)
}

func TestPassphraseInvalidLength(t *testing.T) {
	t.Parallel()

	ls := []int{-1, 0}

	for _, l := range ls {
		if _, err := passphrase.Generate(l); !errors.Is(err, passphrase.ErrInvalidLengthSpecified) {
			t.Errorf("unexpected error for length %d: %v", l, err)
		}
	}
}
