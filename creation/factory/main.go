package main

import (
	"design-pattern/creation/factory/publication"
	"fmt"
)

func main() {
	magazine1, _ := publication.New(publication.Magazine, "Magento", "The Magento", 60)
	magazine2, _ := publication.New(publication.Magazine, "Suki", "Suki Inc", 40)

	news1, _ := publication.New(publication.Newspaper, "The standard", "The standard", 10)
	news2, _ := publication.New(publication.Newspaper, "The time", "The time", 20)

	fmt.Println(magazine1.String())
	fmt.Println(magazine2.String())

	fmt.Println(news1.String())
	fmt.Println(news2.String())
}
