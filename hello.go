package main


import "fmt"


func hello(s string) string {
	if(len(s) == 0) {
		return "Hello, my dude!"
	} else {
		return fmt.Sprintf("Hello, %v!", s)
	}
}
