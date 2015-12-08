package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var indent string
var help bool

func init() {
	flag.StringVar(&indent, "indent", "\t", "indent value: defaults to \t")
	flag.StringVar(&indent, "i", "\t", "short flag for -indent")
	flag.BoolVar(&help, "help", false, "indenter help")
	flag.BoolVar(&help, "h", false, "short flag for -help")
}

func main() {
	os.Exit(realmain())
}

func realmain() int {
	flag.Parse()
	args := flag.Args()
	if help {
		Help()
		return 0
	}
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "\nindenter error: filename required")
	}

	for _, fname := range args {
		// read the file
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		// unmarshal the json
		var j interface{}
		err = json.Unmarshal(b, &j)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		// marshal indent the json
		b, err = json.MarshalIndent(j, "", indent)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		// write the json
		err = ioutil.WriteFile(fname, b, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		fmt.Printf("%s successfully fmt'd.\n", fname)
	}
	return 0
}

func Help() {
	fmt.Println("Help is not yet implemented.")
}
