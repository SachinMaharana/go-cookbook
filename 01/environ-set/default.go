package main

import (
	"log"
	"os"
)

func main() {
	key := "DB_CONN"

	os.Setenv(key, "Custom Value")
	val := GetEnvDefault(key, "Default Value")

	log.Println("The value is", val)

	os.Unsetenv(key)

	val = GetEnvDefault(key, "Another defualt value")

	log.Println("The value is", val)

}

//GetEnvDefault ..
func GetEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return defVal
	}

	return val
}
