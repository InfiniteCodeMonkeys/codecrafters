package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func main() {
	input, err := os.Stdin.Stat()
	var output []rune

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if input.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is not being piped input.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, input)
	}

	print(output)
}
