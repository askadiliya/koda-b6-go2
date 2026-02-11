package main

import "fmt"

// menampilkan koda
type The struct {
	best string
}

type Are struct {
	the The
}

type We struct {
	are Are
}

//menampilkan koda
type Hello struct {
	world string
}

// menampilkan Tech Achademy

type Tech struct {
	academy string
}

type Man struct {
	tech Tech
}

type StrItem struct {
	man []Man
}


type Obj struct {
	str [][][]StrItem
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

	obj := Obj{
		str: [][][]StrItem{
			{}, 
			{}, 
			{}, 
			{ 
				{}, 
				{ 
					{}, 
					{}, 
					{ 
						man: []Man{
							{
								tech: Tech{
									academy: "Tech Academy",
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(we.are.the.best)

	fmt.Println(hello.world)

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

}
