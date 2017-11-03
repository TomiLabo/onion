package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/maxmellon/onion"
	"github.com/maxmellon/onion/printer"
)

func main() {
	args := len(os.Args)
	if args == 0 || args == 1 {
		fmt.Printf(`
Usage:
  $ tmng-mail-lint [filename]
`)
	}

	var wg sync.WaitGroup
	for _, filename := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			data := onion.ReadFile(filename)
			parsedData := onion.Parse(data)
			lintedData := onion.Linting(parsedData)
			printer.SimplePrintResult(filename, lintedData)
			defer wg.Done()
		}(filename)
	}

	wg.Wait()
	return
}
