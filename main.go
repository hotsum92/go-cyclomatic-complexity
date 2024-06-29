package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func ifelse(i int) bool {
	if i%2 == 1 {
		return true
	} else {
		return false
	}
}

func switchcase(month int) string {
	switch month {
	case 1:
		return "January"
	case 2:
		return "February"
	default:
		return ""
	}
}

func complex() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		fmt.Println("infinite loop")
		break
	}

	// loop through list
	list := []int{1, 2, 3, 4, 5}
	for i, v := range list {
		fmt.Println(i, v)
	}

	// loop through map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// defer
	defer fmt.Println("defer")
	fmt.Println("first")

	// panic
	panic("panic")
}

func complex2(i int, j int) {
	if i == 1 {
		if j == 1 {
			fmt.Println("i and j are 1")
		} else {
			fmt.Println("i is 1 but j is not 1")
		}
	} else {
		if j > 3 {
			fmt.Println("i is not 1 but j is 1")
			if i == 2 && j < 5 {
				fmt.Println("i is 2 and j is less than 5")
			}
		} else {
			fmt.Println("i and j are not 1")
		}
	}
}
