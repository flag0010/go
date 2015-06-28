package main

import (
	"fmt"
	"time"
	"strconv"
	"math"
	"common"
)

type Vertex struct {
	X int
	Y int
}

func greet() string{
	return time.Now().String()
}
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	//fmt.Println(time.Now().String())
	fmt.Println(greet())
	k, _ := strconv.Atof64("453")
	fmt.Println(k*k)

	duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}

 	printslice(duplicate)

 	dup_map := dup_count(duplicate)

 	//fmt.Println(dup_map)

 	for k, v := range dup_map {
 		fmt.Printf("Item : %s , Count : %d\n", k, v)
 	}
}
