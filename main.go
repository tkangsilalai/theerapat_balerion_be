package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/tkangsilalai/theerapat_balerion_be/bahttext"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Thai Baht Text Converter")
	fmt.Println("Enter a number (or type 'exit' to quit)")

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye.")
			return
		}

		dec, err := decimal.NewFromString(input)
		if err != nil {
			fmt.Println("Invalid number format")
			continue
		}

		result, err := bahttext.ToThaiBahtText(dec)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}
