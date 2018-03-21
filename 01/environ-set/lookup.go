package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	key := "DB_CONN"
	connStr, ok := os.LookupEnv(key)

	if !ok {
		log.Printf("the env varibale is %s is not set. \n", key)
	}
	fmt.Println(connStr)
}

//unset DB_CONN && go run lookup.go
