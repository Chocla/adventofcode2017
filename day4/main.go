package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)


func main(){
	file,_ := os.Open("pass.txt")
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		valid := true
		m := make(map[string]int) 
		s := strings.Split(scanner.Text(), " ")
		for i := range s {
			m[s[i]]++
		}
		for i := range m {
			if m[i] > 1 {
				valid = false
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println(count)
}