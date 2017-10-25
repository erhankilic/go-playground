/*

Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

package main

import "fmt"

func fibonacci() func() int {
	arr := []int{1, 1}
	return func() int {
		length := len(arr)
		a := arr[length-1]
		b := arr[length-2]
		sum := a + b
		arr = append(arr, sum)
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
