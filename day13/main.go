package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const filePath = "input.txt"
const length = 85

func main(){
	firewall := parseInput()
	severity := 0
	v := make([]bool, length)
	for i := 0; i < length; i++ {
		if len(firewall[i]) > 0 && firewall[i][0] == 1 {
			severity += (i * len(firewall[i]))
		}
		stepFirewall(firewall,v)
	}

	fmt.Println(severity)
}
func stepFirewall(fw [][]int, v []bool) {
	for i := range fw {
		//empty column
		if len(fw[i]) == 0{
			continue 
		}
		for j := range fw[i] {
			if fw[i][j] == 1 {
				if j == 0 {
					v[i] = false
				} else if j == (len(fw[i])-1) {
					v[i] = true
				}
				if !v[i] {
					
					fw[i][j],fw[i][j+1] = 0,1
					break
				} else {
					fw[i][j],fw[i][j-1] = 0,1
					break
				}
			}
		}
	}
}
func parseInput() ([][]int){
	firewall := make([][]int,length)

	file,_ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		index,_ := strconv.Atoi(line[0])
		depth,_ := strconv.Atoi(line[1])
		firewall[index] = make([]int, depth)
		firewall[index][0] = 1 //init security scanner
	}
	return firewall
}