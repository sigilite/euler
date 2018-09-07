package main

import (
	"fmt"
)

func main() {
	a := []string{"JACQUES", "STEVE", "EMILY", "ANDY", "SCOTT", "DAVE", "JENN", "ANTHONY", "STEVEN"}
	fmt.Println(sortNames(a))
}

func sortNames(a []string) []string {
	fmt.Printf("Sorting: %v\n", a)
	if len(a) < 2 {
		return a
	}
	firsthalf := sortNames(a[:len(a)/2])
	secondhalf := sortNames(a[len(a)/2:])
	answer := []string{}
	for len(firsthalf) > 0 && len(secondhalf) > 0 {
		done := false
		firstname := []rune(firsthalf[0])
		secondname := []rune(secondhalf[0])
		for letter := 0; letter < min(len(firstname), len(secondname)); letter++ {

			if firstname[letter] == secondname[letter] {
				continue
			}
			if firstname[letter] < secondname[letter] {
				firsthalf, answer = popOnto(firsthalf, answer)
				done = true
				break
			} else if firstname[letter] > secondname[letter] {
				secondhalf, answer = popOnto(secondhalf, answer)
				done = true
				break
			}

		}
		if done == false {
			if len(firstname) > len(secondname) {
				secondhalf, answer = popOnto(secondhalf, answer)
			} else {
				firsthalf, answer = popOnto(firsthalf, answer)
			}
		}
	}
	if len(firsthalf) == 0 {
		answer = append(answer, secondhalf...)
	} else if len(secondhalf) == 0 {
		answer = append(answer, firsthalf...)
	}
	return answer
}

func popOnto(source, target []string) ([]string, []string) {
	target = append(target, source[0])
	source = source[1:]
	return source, target
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
