package main

import (
	"log"
	"runtime"
)

const info = `The golang version is %s`

func init() {
	log.Println(`Starting...`)
}

func main() {
	ver := runtime.Version()
	log.Printf(info, ver)
}
