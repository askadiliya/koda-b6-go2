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

type Hello struct {
	world string
}

func main() {
	we := We{
		are: Are{
			the: The{
				best: "koda",
			},
		},
	}

	hello := Hello{
		world: "Hello World",
	}

	fmt.Println(we.are.the.best)

	fmt.Println(hello.world)
}
