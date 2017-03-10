#!/bin/bash

LIST_NAME=${1:-"eff-large"}
WORDLIST=$(cat "wordlist-$LIST_NAME.txt")

cat > wordlist.go <<EOF
package passphrase

import (
	"strings"
)

var wordList []string

func init() {
	wordList = strings.Split(\`$WORDLIST\`, "\n")
}
EOF
