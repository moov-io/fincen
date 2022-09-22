// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moov-io/fincen"
)

var (
	programName = filepath.Base(os.Args[0])
	summaryCmd  = "summary"
	validateCmd = "validate"
	reformatCmd = "reformat"
)

func main() {

	versionFlag := flag.Bool("version", false, "show version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Work seamlessly with Fincen BSA form from the command line.\n\nUsage:\n  %s <command> [flags]\n\n", programName)
		fmt.Fprintf(os.Stdout, "Available commands:\n")
		fmt.Fprintf(os.Stdout, "  %s: display form summary \n", summaryCmd)
		fmt.Fprintf(os.Stdout, "  %s: validate financial report form\n", validateCmd)
		fmt.Fprintf(os.Stdout, "  %s: reformat financial report form\n", reformatCmd)
		fmt.Fprintf(os.Stdout, "\n")
	}

	summaryCommand := flag.NewFlagSet(summaryCmd, flag.ExitOnError)
	summaryCommand.Usage = func() {
		fmt.Fprintf(os.Stdout, "Display financial report form summary.\n\nUsage:\n  %s %s <files> \n\n", programName, summaryCmd)
		summaryCommand.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	validateCommand := flag.NewFlagSet(validateCmd, flag.ExitOnError)
	validateCommand.Usage = func() {
		fmt.Fprintf(os.Stdout, "Validate financial report form.\n\nUsage:\n  %s %s <file> \n\n", programName, validateCmd)
		validateCommand.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	reformatCommand := flag.NewFlagSet(reformatCmd, flag.ExitOnError)
	reformatCommand.Usage = func() {
		fmt.Fprintf(os.Stdout, "Reformat financial report form.\n\nUsage:\n  %s %s [flags] <file> \n\n", programName, reformatCmd)
		fmt.Fprintf(os.Stdout, "Flags: \n")
		reformatCommand.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	generateAttrs := reformatCommand.Bool("generate-attrs", false, "specify to regenerate attributes")
	formatType := reformatCommand.String("format", "xml", "specify format type (xml or json)")

	flag.Parse()

	if *versionFlag {
		fmt.Fprintf(os.Stdout, "Version: %s\n\n", fincen.Version)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)

	switch command {
	case summaryCmd:
		inputArgs := os.Args[2:]
		if len(inputArgs) == 0 {
			summaryCommand.Usage()
			os.Exit(1)
		}

		summaryCommand.Parse(inputArgs)

		err := Summary(summaryCommand.Args())
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error summary form files: %s\n", err)
			os.Exit(1)
		}

	case validateCmd:
		inputArgs := os.Args[2:]
		if len(inputArgs) == 0 {
			validateCommand.Usage()
			os.Exit(1)
		}

		validateCommand.Parse(inputArgs)

		err := Validate(validateCommand.Args())
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error validating file: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stdout, "Valid report file: %s\n", inputArgs[0])
		}

	case reformatCmd:
		inputArgs := os.Args[2:]
		if len(inputArgs) == 0 {
			reformatCommand.Usage()
			os.Exit(1)
		}

		reformatCommand.Parse(inputArgs)

		buf, err := Reformat(reformatCommand.Args(), *generateAttrs, *formatType)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed formatting file: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", string(buf))
		}

	default:
		fmt.Fprintf(os.Stdout, "Uknown command: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}

	commands := []*flag.FlagSet{summaryCommand, validateCommand, reformatCommand}
	for _, cmd := range commands {
		if cmd.Parsed() {
			files := cmd.Args()
			if len(files) == 0 {
				cmd.Usage()
				os.Exit(1)
			}
		}
	}
}
