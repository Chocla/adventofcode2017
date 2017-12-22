package main

import (
	"fmt"
	"time"
)
const (
	testStepSize = 3
	inputStepSize = 349
	testIter = 10
	inputIter = 2018
)

func findAnswer(lock []int) (int) {
	for i := range lock {
		if lock[i] == 2017 {
			return lock[i+1]
		}
	}
	return 0
}
func spinLock(step,iter int) (lock []int) {
	position := 0
	lock = append(lock,0)

	for i := 1; i < iter; i++ {
		position = (position + step) % len(lock) + 1
		lock = updateLock(position,i,lock)
		//fmt.Println(lock,position)
	}

	return
}

func updateLock(pos, i int, lock []int) ([]int) {
	tmp := make([]int, len(lock) + 1)
	for j := 0; j < pos; j++ {
		tmp[j] = lock[j]
	}
	tmp[pos] = i
	for j := pos + 1; j < len(tmp);j++ {
		tmp[j] = lock[j-1]
	}

	return tmp
}
func main() {
	t0 := time.Now()
	answer := findAnswer(spinLock(inputStepSize,inputIter))
	fmt.Println("Answer: ", answer, ". Time: ",time.Since(t0))
}