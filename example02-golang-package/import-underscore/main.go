package main

import (
	"fmt"
	"os"

	// "github.com/joho/godotenv/autoload"
	"strconv"
)

func main() {
	project := os.Getenv("GOLANG_PROJECT")
	fmt.Println(project)

	const age = 21
	var name string = "abc"
	name = "nobody"

	fmt.Println(name)
	fmt.Println("Please input your name")

	fmt.Scanln(&name)
	fmt.Println(name + "," + strconv.Itoa(age))
}
