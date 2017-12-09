package main

import (
	"fmt"
	"io/ioutil"
)
//iterate through characters
	//if = {
		//increment level
	//if = <
		//trash flag = true
	//if = !
		//skip next character
	//if >
		//trash flag = false
	// if  }
		//add level to score
		//decement level


func main() {
	test,_ := ioutil.ReadFile("in.txt")
	fmt.Println(string(test))
	trash,skip := false,false
	score := 0

	level := 0
	for i := range test {

		
		char := string(test[i])
		if skip {
			skip = false
			continue
		}

		switch char {
		case "{": 
			if trash {
				break
			}
			level++
			break
		case "}": 
			if trash {
				break
			}
			score+= level
			level--
			break
		case "<": 
			trash = true
			break
		case ">":
			trash = false
			break
		case "!":
			skip = true
			break
		default:
			break
		}
	}
	fmt.Println(score)
}