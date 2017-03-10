# Passphrase generator in Go

[![Build Status](https://travis-ci.org/leonklingele/passphrase.svg?branch=master)](https://travis-ci.org/leonklingele/passphrase)

A simple library to generate passphrases.

## tl;dr

```sh
# Install library
go get -u github.com/leonklingele/passphrase
```

```go
// .. and use it. Generate a 7-word long passphrase
s, err := passphrase.Generate(7)
// Example output: discover thinner scared escalate security glitch getting
```

```go
// You can even specify your own word-separator
passphrase.Separator = "-"
s, err := passphrase.Generate(7)
// Example output (notice the dash!): discover-thinner-scared-escalate-security-glitch-getting
```

## Recommended number of words

Make sure to generate passphrases with at least 7 words. This yields around 90 bits of entropy ([ld(7776^7)](https://www.wolframalpha.com/input/?i=ld(7776%5E7))) which is more than enough.

## Wordlist

This library uses [EFF's large wordlist](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt).
