package main


import "fmt"

func main() {
	outputName := hello("Benny Bwah")
	fmt.Println(outputName)

	outputEmpty := hello("")
	fmt.Println(outputEmpty)
}
