package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("DBU Database Utility")
	fmt.Println("1 - View DBF File")
	fmt.Println("   ESC - Exit")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if string(char) == "1" {
			fmt.Println("Load file")

			fmt.Print("Enter a line of text: ")
			text, _ := reader.ReadString('\n') // Read until newline character
			// Remove trailing newline
			text = text[:len(text)-1]

			fmt.Println("You entered:", text)
		}
		// fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		}
	}
}
