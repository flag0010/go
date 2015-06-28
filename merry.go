package main

import (
	"fmt"
	"github.com/gosimple/listdict"
)

func printslice(slice []string) {
	fmt.Println("slice = ", slice)
}

func main() {
	duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}
    printslice(duplicate)
	d := listdict.Dict{"adfa": 1, "ddd": 44}
	d["fasdf3"] = 77
	fmt.Println(d.Items())
    l := d.Keys()
    fmt.Println(len(l))
}
