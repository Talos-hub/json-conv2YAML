package flags

import (
	"flag"
	"fmt"
	"os"
)

// All flags
type Flag struct {
	Input   string
	Output  string
	Help    bool
	Version bool
}

// Parse seting and retruning our flags
func Parse() *Flag {

	//set flag
	input := flag.String(FlagInput, "", DescInput)
	output := flag.String(FlagOutput, "", DescOutput)
	help := flag.Bool(FlagHelp, false, DescHelp)
	vers := flag.Bool(FlagVersion, false, DescVersion)

	//parsing flag of cmd
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if *vers {
		fmt.Println(version)
		os.Exit(1)
	}

	return &Flag{Input: *input, Output: *output, Help: *help, Version: *vers}

}
