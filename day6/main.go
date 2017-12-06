package main

import (
	"fmt"
)

func main(){
	//test := []int{0,2,7,0}
	input := []int{5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6}
	var seenArrays [][]int
	seenArrays = append(seenArrays,[]int{5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6})
	seen := false
	count := 0
	
	for !seen {
		count++
		max,index := max(input)
		for input[index] = 0; max > 0; max--{
			index = (index + 1) % len(input)
			input[index]++
		}
		for i := range seenArrays {
			if equal(input,seenArrays[i]) {
				seen = true
			}
		}
		a := make([]int, len(input))
		copy(a,input)
		seenArrays = append(seenArrays,a)
	}
	fmt.Println(count)

}
func equal(a,b []int)(bool) {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func max(a []int)(m,j int){
	m = a[0]
	for i := range a {
		if m < a[i] {
			m = a[i]
			j = i
		}
	}
	return
}