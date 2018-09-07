package mathpack

import (
	"strconv"
)

func SumArray(arr []int) int {
	answer := 0
	for _, num := range arr {
		answer += num
	}
	return answer
}

func ProdArray(arr []int) int {
	answer := 1
	for _, num := range arr {
		answer *= num
	}
	return answer
}

func IsPrime(x int) bool {
	if x < 2 {
		return false
	} else if x == 2 {
		return true
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func IsPal(x int) bool {
	s := strconv.Itoa(x)
	for len(s) > 1 {
		if s[0] == s[len(s)-1] {
			s = s[1 : len(s)-1]
		} else {
			return false
		}
	}
	return true
}

func DivisorSum(x int) int {
	// takes a number greater than 1, returns the sum of all divisors STRICTLY SMALLER than it, including 1
	factorList := []int{1}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			factorList = append(factorList, i)
			if i*i == x {
				break
			} else {
				factorList = append(factorList, x/i)
			}
		}
	}
	sum := 0
	for _, n := range factorList {
		sum += n
	}
	return (sum)
}

func PrimeFactors(x int) map[int]int {
	// takes a number and factors it, returning a hashmap of primefactor --> how many copies
	answer := make(map[int]int)
	if x < 2 {
		return nil
	}
	for !IsPrime(x) {
		for i := 2; i < x; i++ {
			if IsPrime(i) && x%i == 0 {
				answer[i] += 1
				x = x / i
			}
		}
	}
	answer[x] += 1
	return answer
}
