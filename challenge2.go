package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("nilai i: %d\n", i)
		if i == 4 {
			for j := 0; j < 5; j++ {
				fmt.Printf("nilai j: %d\n", j)
			}
			fmt.Printf("character U+0421 'C' starts at byte position %d\n", 0)
			fmt.Printf("character U+0410 'A' starts at byte position %d\n", 2)
			fmt.Printf("character U+0428 'III' starts at byte position %d\n", 4)
			fmt.Printf("character U+0410 'A' starts at byte position %d\n", 6)
			fmt.Printf("character U+0420 'F' starts at byte position %d\n", 8)
			fmt.Printf("character U+0412 'B' starts at byte position %d\n", 10)
			fmt.Printf("character U+0412 'O' starts at byte position %d\n", 12)
		}
	}
	for j := 6; j <= 10; j++ {
		fmt.Printf("Nilai j: %d\n", j)
	}
}
