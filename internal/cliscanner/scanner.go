package cliscanner

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	// "log"
	"os"
)

var cliScanner bool

func init () {
	flag.BoolVar(&cliScanner, "cli-scanner", false, "Run CLI scanner.")
}

func scanner() (tokens []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			if string(token) == "quit" {
				return 0, nil, io.EOF
				// err = fmt.Errorf(" quiting...")
			}
		}
		return})
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	// if err := scanner.Err(); err != nil {
	// 	return
	// 	// log.Fatalf(" %v\n", err)
	// }
	return
}


func CliScanner() int {
	if !cliScanner {
		return 0
	}
	fmt.Printf(" Commands: %v\n", scanner())
	return 1
}
