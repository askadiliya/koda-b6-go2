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

// menampilkan koda
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

// menampilkan apel

type Fruit struct {
	is string
}

type Favourite struct {
	fruit Fruit
}

type My struct {
	favourite []Favourite
}

// menampilkan pertambahan

type Num struct {
	first  []int
	second []int
}

func tasks() {
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

	my := []My{
		{
			favourite: []Favourite{
				{},
				{},
				{},
				{
					fruit: Fruit{
						is: "Apple",
					},
				},
			},
		},
	}

	num := Num{
		first:  []int{5, 10, 15},
		second: []int{7, 9, 22},
	}

	fmt.Println(we.are.the.best)

	fmt.Println(hello.world)

	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

	fmt.Println(my[0].favourite[3].fruit.is)

	fmt.Println(num.first[1] + num.second[2])

}
