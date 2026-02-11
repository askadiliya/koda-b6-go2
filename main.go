package main

import "fmt"

func main() {

	scores := []int{50, 75, 66, 20, 32, 90}

	fmt.Println("Slice awal:", scores)

	scores = append(scores[:3], append([]int{88}, scores[3:]...)...)

	fmt.Println("Setelah insert 88:", scores)

	fmt.Println("\nTampilkan per score:")
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	scores[2], scores[3] = scores[3], scores[2]

	fmt.Println( scores)
}
