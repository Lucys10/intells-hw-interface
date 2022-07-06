package main

const (
	feedPerWightDog      = 10
	feedPerWightCat      = 7
	feedPerWightCow      = 25
	dogFeedPerWeightOnKg = 5
)

type dog struct {
	name         string
	feedPerWight int
	weight       int
}

func newDog(name string, weight int) dog {
	return dog{
		name:         name,
		feedPerWight: feedPerWightDog,
		weight:       weight,
	}
}

func (d dog) feedPerMonth() int {
	return d.feedPerWight * d.weight / dogFeedPerWeightOnKg
}

func (d dog) String() string {
	return "Dog - " + d.name
}

func (d dog) getWeight() int {
	return d.weight
}

type cat struct {
	name         string
	feedPerWight int
	weight       int
}

func newCat(name string, weight int) cat {
	return cat{
		name:         name,
		feedPerWight: feedPerWightCat,
		weight:       weight,
	}
}

func (c cat) feedPerMonth() int {
	return c.feedPerWight * c.weight
}

func (c cat) String() string {
	return "Cat - " + c.name
}

func (c cat) getWeight() int {
	return c.weight
}

type cow struct {
	name         string
	feedPerWight int
	weight       int
}

func newCow(name string, weight int) cow {
	return cow{
		name:         name,
		feedPerWight: feedPerWightCow,
		weight:       weight,
	}
}

func (cw cow) feedPerMonth() int {
	return cw.feedPerWight * cw.weight
}

func (cw cow) String() string {
	return "Cow - " + cw.name
}

func (cw cow) getWeight() int {
	return cw.weight
}
