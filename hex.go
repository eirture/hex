package main

import (
	"encoding/hex"
	"flag"
	"io"
	"os"
)

var (
	decode = flag.Bool("d", false, "Decode incoming Hex stream into binary data.")
)

func main() {
	flag.Parse()

	var r io.Reader = os.Stdin
	if *decode {
		io.Copy(os.Stdout, hex.NewDecoder(r))
	} else {
		io.Copy(hex.NewEncoder(os.Stdout), r)
	}
}
