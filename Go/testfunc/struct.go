package main

import "fmt"

type AnimalCategory struct {
	kingdom string
	species string
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s", ac.kingdom, ac.species)
}
