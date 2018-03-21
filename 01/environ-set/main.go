package main

import (
	"fmt"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	fmt.Println(`Connection string: `, connStr)
}

//DB_CONN=db:/user@example && go run main.go
