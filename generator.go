package passphrase

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	minWordListLength = 1000
)

var (
	// ErrInvalidLengthSpecified is returned when the length specified is invalid
	ErrInvalidLengthSpecified = errors.New("invalid passphrase length specified")
	// ErrInvalidWordListLength is returned when the length of the word list is invalid
	ErrInvalidWordListLength = errors.New("invalid word list length specified")

	// Separator specifies the delimiter to join passphrase words with
	// Example (separator is a dash symbol): correct-horse-battery-staple
	Separator = " " //nolint:gochecknoglobals // TODO: Get rid of global variable as it makes everything racy
)

// Generate generates a passphrase with length 'l' words each separated by passphrase.Separator
func Generate(l int) (string, error) {
	// Make sure the list is long enough
	if len(wordList) < minWordListLength {
		return "", ErrInvalidWordListLength
	}

	// Length needs to be in range [1, 1<<31-1]
	if l <= 0 || l > 1<<31-1 {
		return "", ErrInvalidLengthSpecified
	}

	buf := make([]string, l)
	maxlen := big.NewInt(int64(len(wordList)))

	for i := range l {
		index, err := randomInt(maxlen)
		if err != nil {
			return "", err
		}

		buf[i] = wordList[index]
	}

	return strings.Join(buf, Separator), nil
}

func randomInt(maxlen *big.Int) (int, error) {
	i, err := rand.Int(rand.Reader, maxlen)
	if err != nil {
		return 0, err //nolint:wrapcheck // Fine to not wrap this here
	}

	return int(i.Int64()), nil
}
