package main

/**
 * A command like cat(1), but slow.
 *
 *    $ cat foo.txt
 *    $ cat foo.txt bar.txt
 *    $ cat < foo.txt
 */

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		var err error = nil
		for err == nil {
			_, err = io.CopyN(os.Stdout, os.Stdin, 1)
			time.Sleep(12_000_000)
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			fh, err := os.Open(fname)
			if err != nil {
				log.Fatal(err)
			}

			for err == nil {
				_, err = io.CopyN(os.Stdout, fh, 1)
				time.Sleep(12_000_000)
			}

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}
	}
}
