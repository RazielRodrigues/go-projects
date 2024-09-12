package main

import "fmt"

type item struct {
	itemId int
	price  float64
	name   string
}

func (i *item) changePrice(newValue float64) {
	i.price = newValue
}

type cart struct {
	userId int
	items  []item
}

func (c cart) getTotal() float64 {
	total := 0.0
	for _, v := range c.items {
		total += v.price
	}
	return total
}

func main() {

	c1 := cart{
		1, []item{
			item{1, 10.2, "presunto"},
			item{2, 10.3, "queijo"},
		},
	}

	fmt.Println(c1.getTotal())

	itemChanged := item{1, 2.2, "watch"}
	fmt.Println(itemChanged.price)

	itemChanged.changePrice(125.0)
	fmt.Println(itemChanged.price)

}
