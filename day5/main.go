package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	f,_ := os.Open("path.txt")
	scanner := bufio.NewScanner(f)
	path := make([]int,0)
	for scanner.Scan() {
		tmp := scanner.Text()
		num,_ := strconv.Atoi(tmp)
		path = append(path,num)
	}
	fmt.Println(followPath(path))
	
}

func followPath(path []int)(steps int) {
	for i := 0; i >= 0 && i < len(path); {	
		path[i]++
		i += path[i]-1
		steps++
	}
	return
}