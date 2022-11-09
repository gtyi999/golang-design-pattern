package main

import (
	"fmt"
	simplefactory "golang-design-pattern/00_simple_factory"
)

func main()  {
	api := simplefactory.NewAPI(1)
	fmt.Println(api.Say("Tom"))
	api = simplefactory.NewAPI(2)
	fmt.Println(api.Say("Tom"))
}
