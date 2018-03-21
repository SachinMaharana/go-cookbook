package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	fmt.Println("Version with which binary is run: ", runtime.Version())
	if len(os.Args) == 1 {
		fmt.Println(`This binary expects non-zero arguments..`)
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	binary := args[0]
	fmt.Println(`Binary is :`, binary)
	restArgs := args[1:]
	fmt.Println(`Arguments given to this binary are: `)
	for i, e := range restArgs {
		fmt.Printf("Arg %d : %s\n", i+1, e)
	}
}
