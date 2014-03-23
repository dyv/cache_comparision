package main

/* cache_comp: is an algorithm for minimizing the memory cost of
   memory, when each access to an element, has an associated memory
   cost of O(log(l)) where l is the last time it was used. It employs
   a divide and conquer strategy for the main algorithm, and a divide
   and conquer strategy for the conquer step itself.
*/

import (
	"fmt"
	"math"
)

var cache []int
var time float64
var num_items int

func Find(x int) int {
	for i := 0; i < len(cache); i++ {
		if cache[len(cache)-i-1] == x {
			//fmt.Println("Cache Hit")
			return i/2 + 1
		}
	}
	//fmt.Println("Cache Miss:", len(cache), num_items)
	if len(cache) < 2*num_items {
		return num_items
	}
	return len(cache) / 2
}

func comp(x, y int) {
	//fmt.Println("comp: ", x, y)
	if x == y {
		return
	}
	xi := Find(x)
	yi := Find(y)
	time += math.Log2(float64(xi))
	time += math.Log2(float64(yi))
	cache = append(cache, x)
	cache = append(cache, y)
	//fmt.Println("comp dist: ", xi, yi)
}

func conquer(a, b []int) {
	if len(a) <= 1 || len(b) <= 1 {
		if len(a) == 1 && len(b) == 1 {
			comp(a[0], b[0])
		}
		return
	}
	conquer(a[:len(a)/2], b[:len(b)/2])
	conquer(a[len(a)/2:], b[:len(b)/2])
	conquer(a[:len(a)/2], b[len(b)/2:])
	conquer(a[len(a)/2:], b[len(b)/2:])
}

func compare(a, b []int) {
	//fmt.Println("comparing:", a, b)
	if len(a) <= 1 || len(b) <= 1 {
		if len(a) == 1 && len(b) == 1 {
			comp(a[0], b[0])
		}
		return
	}
	compare(a[:len(a)/2], a[len(a)/2:])
	compare(b[:len(b)/2], b[len(b)/2:])
	conquer(a, b)
}
func main() {
	iterations := 13
	results := make([]float64, iterations)
	for i := 0; i < iterations; i++ {
		time = 0
		num_items = int(math.Pow(2, float64(i)))
		cache = make([]int, 0, 2*num_items*num_items)
		a := make([]int, num_items)
		for i := range a {
			a[i] = i
		}
		//fmt.Println(num_items, cache, a)
		compare(a[:len(a)/2], a[len(a)/2:])
		fmt.Println(num_items, ", ", time)
		results[i] = time
	}
}
