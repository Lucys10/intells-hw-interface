package main

import "fmt"

type animals interface {
	feedPerMonth() int
	getWeight() int
	fmt.Stringer
}

func main() {

	maxAnimalsOnFarm := 30

	animalsOnFarm := fakeAnimalsOnFarm(maxAnimalsOnFarm)

	sumFeedOnFarm := infoAnimalsOnFarm(animalsOnFarm)

	fmt.Printf("Need sum feed on farm - %v kg.", sumFeedOnFarm)
}

func infoAnimalsOnFarm(animalsOnFarm []animals) int {
	sumFeedOnFarm := 0

	for _, animal := range animalsOnFarm {
		name := animal.String()
		weight := animal.getWeight()
		feedPerMonth := animal.feedPerMonth()

		fmt.Printf("%v. Weight - %v kg. Feed per month - %v kg. \n", name, weight, feedPerMonth)

		sumFeedOnFarm += feedPerMonth
	}

	return sumFeedOnFarm
}
