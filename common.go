package main

import (
	"fmt"
	"strconv"
	"sort"
)

func Int2str(x int64) string {
	k := strconv.FormatInt(x, 10)
	return k
}

func Float2str(x float64) string {
	k := strconv.FormatFloat(x, 'g', -1, 64)
	return k
}

func Str2int(x string) int64 {
	k, _ := strconv.ParseInt(x, 10, 64)
	return k
}

//use, range is (-9223372036854775808 to 9223372036854775807)
func Str2float(x string) float64 {
	k, _ := strconv.ParseFloat(x, 64)
	return k
}

func Count_all(list []string, freq int) map[string]float64 {
	//if freq=0, count all items in list and return map of item => count
	//if freq is set to non-zero value instead you get item => freq
	duplicate_frequency := make(map[string]float64)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1.0 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1.0 // else start counting from 1
		}
	}
	if freq == 0 {
		return duplicate_frequency
	} else {
		duplicate_frequency2 := make(map[string]float64)
		sum := 0.0
		for _, v := range duplicate_frequency {
			sum += v
		}
		for k, v := range duplicate_frequency {
			duplicate_frequency2[k] = v / sum
		}
		return duplicate_frequency2
	}
}

func Rev(list []int) []int {
	//reverse a slice of ints
	rplace := len(list)
	out := make([]int, rplace, rplace)
	for i := 0; i < rplace; i++ {
		out[rplace-1-i] = list[i]
	}
	return out
}

func main() {
	duplicate := []string{"You", "ARE", "ARE", "awesome!"}
	fmt.Println(duplicate)
	dup_map := Count_all(duplicate, 0)
	fmt.Println(dup_map)
	//for k, v := range dup_map {
	//	fmt.Printf("Item : %s , Count : %d\n", k, v)
	//}
	duplicate = []string{Int2str(2343), Int2str(2343), Float2str(32.323), Int2str(234)}
	dup_map = Count_all(duplicate, 1)
	fmt.Println(dup_map)
        example := []int{1, 25, 3, 5, 4}
	sort.Ints(example)
	example = Rev(example)
	fmt.Println(example)
}
