package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
const path = "input.txt"

func main() {
	pipes := parseInput()
	x := countConnections(pipes)
	fmt.Println(x)
}
func countConnections(pipes [][]string) (c int) {
	m := make(map[string]bool)
	changed := true
	for i := range pipes {
		m[pipes[i][0]] = false
	}
	m["0"] = true
	//loop through until nothing is changed
	for changed {
		changed = false
		for i := range pipes {
			for j := range pipes[i] {
				//unknown if connected
				if !m[pipes[i][0]] {
					break
				}
				//skip "<->"
				if pipes[i][j] == "<->" {
					continue
				} else if !m[pipes[i][j]]{ //if not already listed
					m[pipes[i][j]] = true
					changed = true
				}

			}
		}
	}
	for i := range m {
		if m[i] {
			c++
		}
	}
	//fmt.Println(m)
	return 
}

func parseInput() (pipes [][]string) {
	file,_ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//parse each line and append to array
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for i := range line {
			line[i] = strings.Trim(line[i],",")
		}
		pipes = append(pipes,line)
		
	}
	return
}