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
	Separator = " "
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
	max := big.NewInt(int64(len(wordList)))

	for i := 0; i < l; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buf[i] = wordList[index]
	}

	return strings.Join(buf, Separator), nil
}

func randomInt(max *big.Int) (int, error) {
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(i.Int64()), nil
}
