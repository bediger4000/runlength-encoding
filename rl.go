package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	decodeFlag := flag.Bool("d", false, "decode run-length encoded string")
	flag.Parse()

	if flag.NArg() == 0 {
		return
	}

	for _, str := range flag.Args() {
		var output string
		if *decodeFlag {
			output = runLengthDecode(str)
			fmt.Printf("Decoded: %q\n", output)
			continue
		}
		output = runLengthEncode(str)
		fmt.Printf("Encoded: %q\n", output)
	}
}

func runLengthDecode(encoded string) string {
	fmt.Printf("Decode %q\n", encoded)
	var output []rune
	var run int
	var err error
	runes := []rune(encoded)
	for i, r := range runes {
		if unicode.IsDigit(r) {
			run, err = strconv.Atoi(string(runes[i : i+1]))
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		for j := 0; j < run; j++ {
			output = append(output, r)
		}
	}
	return string(output)
}

func runLengthEncode(cleartext string) string {
	var output string
	last := ' '
	count := 0
	for _, r := range []rune(cleartext) {
		if r != last || count == 9 {
			if count > 0 {
				output += fmt.Sprintf("%d%c", count, last)
			}
			last = r
			count = 0
		}
		count++
	}
	if count > 0 {
		output += fmt.Sprintf("%d%c", count, last)
	}
	return string(output)
}
