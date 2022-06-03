package main

import (
	"flag"
	"fmt"
)

type Params struct {
	ForceSearch bool
	Info        string
	Download    bool
	SearchBy    string
}

var params Params

func main_t() {
	flag.BoolVar(&params.ForceSearch, "a", false, "Print all matches")
	flag.StringVar(&params.SearchBy, "by", "", "Specify field to search by")

	flag.StringVar(&params.Info, "i", "", "Request info about a package provided as argument")
	flag.Parse()
	if params.ForceSearch {
		fmt.Println("Search forced...")
	}
	fmt.Println(flag.Args())
	fmt.Println()
}
