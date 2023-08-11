package main

import (
	"encoding/hex"
	"flag"
	"io"
	"os"
	"strings"
)

var (
	decode = flag.Bool("d", false, "Decode incoming Hex stream into binary data.")
)

func main() {
	flag.Usage = func() {
		os.Stderr.WriteString("Usage: hex [-d] [string]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	var r io.Reader = os.Stdin
	args := flag.Args()
	if len(args) > 0 {
		r = strings.NewReader(args[0])
	}
	if *decode {
		io.Copy(os.Stdout, hex.NewDecoder(r))
	} else {
		io.Copy(hex.NewEncoder(os.Stdout), r)
	}
}
