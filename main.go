package main

import (
	"fmt"
)

func main() {
	task1 := "Learn Go lang"
	task2 := "Build Api In Go Lang"
	task3 := "condition"
	task4 := "stop"

	tasks := []string{task1, task2, task3, task4}

	names := []string{"John", "Doe", "Smily", "John", "Wick"}

	names = append(names, "John Wick")

	loops(names)

	loops(tasks)

	nums := []int{1, 2, 3}

	println(cap(nums)) // Cap refer to capacity

	println(nums == nil)

	animal := make([]string, 3)

	animal[0] = "dog"
	animal[1] = "cat"
	animal[2] = "lion"

	println(len(animal))
	println(cap(animal))

	println(animal[0], animal[1], animal[2])

	// include, exclude

	values := []int{1, 2, 3, 4}

	// exclude values after specific index
	fmt.Println(values[:2])

	// include values after specific index
	fmt.Println(values[2:])

	doubleArray := [][]string{{"john", "doe"}, {"john", "smith"}}

	fmt.Println(doubleArray[0])
	fmt.Println(doubleArray[1])

	// Maps in Go

	// maps -> object, dict , hash

	m := make(map[string]string)

	m["name"] = "john doe"
	m["area"] = "backend"

	fmt.Println(m["name"])
	fmt.Println(m["area"])

	// If key does not exist in map then it returns zero or boolean value for string type value

	n := make(map[string]int)

	n["age"] = 20
	n["phone"] = 1098765432

	// delete(n, "phone") delete an item from map based on key

	// map conditions

	v, ok := n["phone"]

	fmt.Println(v) // get the value

	if ok {
		fmt.Println("good to go")
	} else {
		fmt.Println("something is missing here")
	}

	for index, value := range "motu" {
		fmt.Println(index, string(value))
	}

	fmt.Println("Hello World")
}

func loops(items []string, condition ...string) {
	searchCondition := ""

	if len(condition) > 0 {
		searchCondition = condition[0]
	}

	for index, item := range items {
		if searchCondition != "" && item == searchCondition {
			fmt.Printf("%d. %s \n", index+1, item)
			break
		} else {
			fmt.Printf("%d. %s \n", index+1, item)
		}

	}
}

// We can slices instead of arrays slice helps in memory management it manage memory dynamically provinding better memory management
