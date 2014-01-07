package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

const (
	defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

var (
	charsetString    string
	verbose          bool
	charset          []byte
	key, originalKey []byte
	candidates       = make(map[int][]byte)
)

// Recursively iterates over all key permutations
func permute(i int) {
	if i == len(key) {
		fmt.Printf("%s\n", key)
		return
	}
	if originalKey[i] != '?' {
		permute(i + 1)
		return
	}
	chars := getCandidates(i)
	for j := 0; j < len(chars); j++ {
		key[i] = chars[j]
		permute(i + 1)
	}
}

// Returns slice of candidate chars for index i
func getCandidates(i int) []byte {
	chars, ok := candidates[i]
	if ok {
		return chars
	}
	return charset
}

// Processes and stores candidates from key pattern. Also does some crude error
// checking.
func processKey() error {
	open := -1     // Stripped index
	openOrig := -1 // Original index
	for i, j := 0, 0; i < len(originalKey); i++ {
		c := originalKey[i]
		switch c {
		case '{':
			if open != -1 {
				return fmt.Errorf("nested brace (%d)", i)
			}
			open = j
			openOrig = i
			continue
		case '}':
			if open == -1 {
				return fmt.Errorf("unmatched closing brace (%d)", i)
			}
			if openOrig == i-1 {
				return fmt.Errorf("empty brace statement (%d)", openOrig)
			}
			open = -1
			key = append(key, '?')
			j++
		default:
			if open == -1 {
				key = append(key, c)
				j++
			} else {
				// Only check candidates for illegal chars
				if bytes.IndexByte(charset, c) == -1 {
					return fmt.Errorf("'%c' is not contained by charset (%d)", c, i)
				}
				candidates[open] = append(candidates[open], c)
			}
		}
	}

	if open >= 0 {
		return fmt.Errorf("unmatched opening brace (%d)", openOrig)
	}

	// fmt.Println("------")
	// for k, v := range candidates {
	// 	fmt.Printf("%d => %s\n", k, v)
	// }
	// fmt.Println("------")

	return nil
}

// Set up flag parameters
func init() {
	const usage = "set of permissible key characters"
	flag.StringVar(&charsetString, "c", defaultCharset, usage)
	flag.BoolVar(&verbose, "v", false, "be verbose")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-c charset] [-v] pattern\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		// fmt.Fprintln(os.Stderr, "missing key parameter")
		flag.Usage()
		os.Exit(1)
	}
	charset = []byte(charsetString)
	originalKey = []byte(flag.Arg(0))
	if err := processKey(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "# Original: %s\n", originalKey)
		fmt.Fprintf(os.Stderr, "# Stripped: %s\n\n", key)
	}
	originalKey = append([]byte(nil), key...)
	permute(0)
}
