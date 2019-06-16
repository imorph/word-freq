package main

import (
	"os"
	"bufio"
	"fmt"
	"log"

	"github.com/imorph/word-freq/lib/freq"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nlet's analyse!\nPress any time CTRL-C to exit")
	fmt.Println("")
	
	for {
		fmt.Print("Please enter text for analysis: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Entered string in slice form: ", freq.Top10(StatMapFromSlice(StringToCleanSlice(text))))
			fmt.Println("")
		}
	}
}