package main

import (
	"fmt"
	"log"

	"github.com/imorph/word-freq/lib/freq"
)


func main() {
	fmt.Println("\nlet's analyse!\nPress any time CTRL-C to exit")
	fmt.Println("")

	var s string
	
	for {
		fmt.Print("Please enter text for analysis: ")

		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Entered string in slice form: ", freq.StringToCleanSlice(s))
			fmt.Println("")
		}
	}
}