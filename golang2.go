package main

import "fmt"

type The struct {
	best string
}

type Are struct {
	the The
}

type We struct {
	are Are
}

func main() {
	we := We{
		are: Are{
			the: The{
				best: "koda",
			},
		},
	}

	fmt.Println(we.are.the.best)
}
