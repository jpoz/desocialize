package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpoz/desocialize"
)

const Version = "1.0.0"

func main() {
	var effect string

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] input.jpg output.jpg\nv%s\n\n",
			os.Args[0],
			Version,
		)
		flag.PrintDefaults()
	}

	flag.StringVar(&effect, "e", "", "Effect to put on faces (default: )")
	flag.Parse()

	in := flag.Arg(0)

	dsclze := &desocialize.Desocalizer{
		Effect: effect,
	}

	dsclze.Desocialize(in)
}

func fail(s string, e error) {
	fmt.Fprintf(os.Stderr, s)
	flag.Usage()
	check(e)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
