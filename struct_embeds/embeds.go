package structembeds

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num:%d\n", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func Run() {
	co := container{
		str: "some str",
		base: base{
			num: 10,
		},
	}

	fmt.Printf("co: %v\n", co)
	fmt.Printf("co.base: %v\n", co.base)

	fmt.Printf("co.base.num: %v\n", co.base.num)
	fmt.Printf("co.num: %v\n", co.num)
	fmt.Printf("co.base.describe(): %v\n", co.base.describe())
	fmt.Printf("co.describe(): %v\n", co.describe())

	var d describer = co
	fmt.Printf("d.describe(): %v\n", d.describe())
}
