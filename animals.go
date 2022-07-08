package main

import "github.com/brianvoe/gofakeit/v6"

func fakeAnimalsOnFarm(maxAnimalsOnFarm int) []animals {

	animalsOnFarm := make([]animals, 0, maxAnimalsOnFarm)

	quantityDogOnFarm := gofakeit.Number(1, 10)
	quantityCatOnFarm := gofakeit.Number(1, 10)
	quantityCowOnFarm := gofakeit.Number(1, 10)

	dogs := fillingDog(quantityDogOnFarm)
	cats := fillingCat(quantityCatOnFarm)
	cows := fillingCow(quantityCowOnFarm)

	animalsOnFarm = append(animalsOnFarm, dogs...)
	animalsOnFarm = append(animalsOnFarm, cats...)
	animalsOnFarm = append(animalsOnFarm, cows...)

	return animalsOnFarm

}

func fillingDog(quantityAnimals int) []animals {
	dogs := make([]animals, 0, quantityAnimals)
	for i := 0; i < quantityAnimals; i++ {
		name := gofakeit.PetName()
		weight := gofakeit.Number(5, 50)
		d := newDog(name, weight)
		dogs = append(dogs, d)
	}
	return dogs
}

func fillingCat(quantityAnimals int) []animals {
	cats := make([]animals, 0, quantityAnimals)
	for i := 0; i < quantityAnimals; i++ {
		name := gofakeit.PetName()
		weight := gofakeit.Number(5, 20)
		c := newCat(name, weight)
		cats = append(cats, c)
	}
	return cats
}

func fillingCow(quantityAnimals int) []animals {
	cows := make([]animals, 0, quantityAnimals)
	for i := 0; i < quantityAnimals; i++ {
		name := gofakeit.PetName()
		weight := gofakeit.Number(15, 100)
		cw := newCow(name, weight)
		cows = append(cows, cw)
	}
	return cows
}
