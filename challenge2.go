package main

import "fmt"

func main() {
	i:=0
	for i<5; i++{
		fmt.Println("Nilai i=%d\n",i)
	}

	for j:=0; j<=10; j++{
		if j==5{
			var chr = []rune{}
			chr = append(chr, 'C','A','B','A','P','B','O')
			position := 0
			for _, value := range chr{
				fmt.Println("Character %U'%c' starts at byte position %d\n", value, value, position)
				position += 2
			}
		}else{
			fmt.Println("Nilai j= ", j)
		}
	}
}
