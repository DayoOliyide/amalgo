package main

import (
	flags "github.com/jessevdk/go-flags"
	"os"
	"fmt"
	"github.com/DayoOliyide/amalgo/core"
)

const (
	Version = "0.0.1"
	FullVersion = "Amalgo-" + Version
)

var opts struct{
	Version bool `short:"v" long:"version" description:"Print Version Info"`
}

func main()  {
	args, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println("Version", FullVersion)
		os.Exit(0)
	}

	for _, fileName := range args {
		fmt.Println(core.OutfileName(fileName))
	}
}
