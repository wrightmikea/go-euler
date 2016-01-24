// euler1 - sum subset of first 1000 numbers which are multiples of 3 or 5
package main

import (
	"fmt"
	"runtime"
)

const max = 1000

func main() {
	fmt.Printf("euler1 NumCPU()=%v\n", runtime.NumCPU())
	fmt.Printf("Euler1a sum=%d\n", Euler1a(max))
	fmt.Printf("Euler1b sum=%d\n", Euler1b(max))
	fmt.Printf("Euler1c sum=%d\n", Euler1c(max))
	fmt.Printf("Euler1d sum=%d\n", Euler1d(max))
	fmt.Printf("Euler1e sum=%d\n", Euler1e(max))
	fmt.Printf("Euler1f sum=%d\n", Euler1f(max))
	fmt.Printf("Euler1g sum=%d\n", Euler1g(max))
	fmt.Printf("Euler1h sum=%d\n", Euler1h(max))
}

// ****************************************************************
func Euler1a(max int) (sum int) {
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			//fmt.Printf("i=%d (sum=%d); ", i, sum)
		}
	}
	return
}

// ****************************************************************
func Euler1b(max int) (sum int) {
	sum += sum3s(max)
	sum += sum5s(max)
	sum -= sum15s(max)
	return
}
func sum3s(max int) (sum int) {
	for i := 0; i < max; i += 3 {
		sum += i
	}
	return
}
func sum5s(max int) (sum int) {
	for i := 0; i < max; i += 5 {
		sum += i
	}
	return
}
func sum15s(max int) (sum int) {
	for i := 0; i < max; i += 15 {
		sum += i
	}
	return
}

// ****************************************************************
func Euler1c(max int) (sum int) {
	sum += sumMultiplesToMax(3, max)
	sum += sumMultiplesToMax(5, max)
	sum -= sumMultiplesToMax(3*5, max)
	return
}
func sumMultiplesToMax(mult int, max int) (sum int) {
	for i := 0; i < max; i += mult {
		sum += i
	}
	return
}

// ****************************************************************
func Euler1d(max int) (sum int) {
	sum = sumTwoMultiplesToMax(3, 5, max)
	return
}
func sumTwoMultiplesToMax(first int, second int, max int) (sum int) {
	for i1 := 0; i1 < max; i1 += first {
		sum += i1
	}
	for i2 := 0; i2 < max; i2 += second {
		sum += i2
	}
	product := first * second
	for dup1 := 0; dup1 < max; dup1 += product {
		sum -= dup1
	}
	return
}

// ****************************************************************
func Euler1e(max int) (sum int) {
	sum = sumGenerators(3, 5, max)
	return
}
func sumGenerators(first int, second int, max int) (sum int) {
	var nextFirst, nextSecond, nextDup = 0, 0, 0
	product := first * second
	gen := func() (next int) {
		nextFirst += first
		if nextFirst < max {
			next = nextFirst
			return
		}
		nextSecond += second
		if nextSecond < max {
			next = nextSecond
			return
		}
		nextDup += product
		if nextDup < max {
			next = -nextDup
			return
		}
		return 0
	}
	next := gen()
	for next != 0 {
		sum += next
		next = gen()
	}
	return
}

// ****************************************************************
func Euler1f(max int) (sum int) {
	for i := range foo(3, 5, max) {
		sum += i
	}
	return
}

func foo(first int, second int, max int) chan int {
	c := make(chan int, 50)
	product := first * second
	go func() {
		for i := 0; i < max; i += first {
			c <- i
		}
		for j := 0; j < max; j += second {
			c <- j
		}
		for k := 0; k < max; k += product {
			c <- -k
		}
		close(c)
	}()
	return c
}

// ****************************************************************
func Euler1g(max int) (sum int) {
	for i := range bar(3, 5, max)[0] {
		sum += i
	}
	for i := range bar(3, 5, max)[1] {
		sum += i
	}
	for i := range bar(3, 5, max)[2] {
		sum += i
	}
	return
}

func bar(first int, second int, max int) []chan int {
	var c = []chan int{
		make(chan int, 50),
		make(chan int, 50),
		make(chan int, 50),
	}
	product := first * second
	go func() {
		for i := 0; i < max; i += first {
			c[0] <- i
		}
		close(c[0])
	}()
	go func() {
		for j := 0; j < max; j += second {
			c[1] <- j
		}
		close(c[1])
	}()
	go func() {
		for k := 0; k < max; k += product {
			c[2] <- -k
		}
		close(c[2])
	}()
	return c
}

//****************************************************************
type E struct {
	First, Second, Product, Sum int
}

func (e *E) update(first, second, product int, max int) {
	if e.First+first < max {
		e.First += first
		e.Sum += e.First
	}
	if e.Second+second < max {
		e.Second += second
		e.Sum += e.Second
	}
	if e.Product+product < max {
		e.Product += product
		e.Sum -= e.Product
	}
}

func Euler1h(max int) (sum int) {
	var e E
	for i := 0; i < max; i++ {
		e.update(3, 5, 3*5, max)
	}
	sum = e.Sum
	return
}

// ****************************************************************
