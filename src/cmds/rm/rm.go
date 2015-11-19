// Copyright 2013 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"flag"
)

var (
	recursive_flag = flag.Bool("R", false, "Remove file hierarchies.")
	recursive_alias = flag.Bool("r", false, "Equivalent to -R.")
	verbose = flag.Bool("v", false, "Verbose mode.")
	cmd = struct { name, flags string } {
		"rm",
		"[-Rrv] file...",
	}
)

// rm function 
func rm(files []string, do_recursive bool, verbose bool) error {
	f := os.Remove
	if do_recursive {
		f = os.RemoveAll
	}

	// looping for removing files and folders
	for _, file := range(files) {
		err := f(file)
		if err != nil {
			fmt.Printf("%v: %v\n", file, err)
			return err
		}

		if verbose {
			fmt.Printf("Deleting: %v\n", file)
		}
	}
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:", cmd.name, cmd.flags)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	recursive := *recursive_flag || *recursive_alias
	flag.Parse()


	if flag.NArg() < 1 {
		usage()
	}

	rm(flag.Args(), recursive , *verbose)
}
