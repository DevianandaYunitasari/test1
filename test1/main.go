package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func triangularNumbers(n int) []int {
	result := make([]int, n)
	currentSum := 0

	for i := 1; i <= n; i++ {
		currentSum += i
		result[i-1] = currentSum
	}

	return result
}

func main() {
	for {
		fmt.Print("Silahkan masukkan nilai n atau ketik 'exit' untuk keluar: ")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error, silahkan masukan nilai yang sesuai:", err)
			return
		}
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "exit" {
			fmt.Println("Terima kasih, Program telah berhenti.")
			return
		}
		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input harus berupa bilangan bulat atau 'exit' untuk keluar.")
			continue
		}

		if n < 1 {
			fmt.Println("Input harus berupa bilangan bulat positif atau 'exit' untuk keluar.")
			continue
		}

		output := triangularNumbers(n)
		fmt.Printf("Output: %v\n", output)
	}
}
