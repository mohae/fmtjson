package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	indent = "\t"
	spaces int
	help   bool
)

func init() {
	flag.IntVar(&spaces, "spaces", 0, "number of spaces to indent; by default, if not specified, or 0, '\t' is used for indents")
	flag.IntVar(&spaces, "s", 0, "the short flag for -spaces")
	flag.BoolVar(&help, "help", false, "indenter help")
	flag.BoolVar(&help, "h", false, "the short flag for -help")
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
	if spaces > 0 {
		for i := 0; i < spaces; i++ {
			indent += " "
		}
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "\nfmtjson error: filename required")
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
	helpText := `
Usage: fmtjson [options] <filename...>

fmtjson will format the json in the received list of files.
If more than one file is specified, they are separated by
spaces.  The formatted json will replace the content of the
input file.

    $ fmtjson file.json
    $ fmtjson file1.json file2.json

The indentation is configurable using the '-s' or '-spaces'
flag.  A value > 0 must be specified.  Tab, '\t', is the
default indent.

    $ fmtjson -s 2 file.json

Flags:

Short  Flag      Type   Default
-------------------------------
-s     -spaces   int
-h     -help     bool   false
`
	fmt.Println(helpText)
}
