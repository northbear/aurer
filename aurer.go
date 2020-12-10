package main

import (
	"fmt"
	"log"
	"os"
)

func main() {	
	cliopt, err := GetActionConfig(os.Args)
	if err == ErrorNoParams {
		fmt.Println(Usage())
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CLI params: ", cliopt)
	fmt.Println("CLI param out of index: ", os.Args[2])
	// fmt.Println("Hello, world!", cliopt.SearchPattern, "...")
}
