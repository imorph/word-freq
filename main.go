package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
			fmt.Println("10 most used words (high to low): ", freq.Top10(freq.StatMapFromSlice(freq.StringToCleanSlice(text))))
			fmt.Println("")
		}
	}
}
