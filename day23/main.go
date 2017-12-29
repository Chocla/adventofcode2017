package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type instruction struct {
	op string
	reg string
	val string
}
const inputPath = "input.txt"
func main() {
	reg := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
	}
	ins := parseInstructions(inputPath)
	count := execInstructions(ins,reg)
	fmt.Println(count)
}

func execInstructions(ins []instruction, reg map[string]int) (c int) {
	for i := 0; i < len(ins);i++ {
		switch ins[i].op {
		case "set":
			reg[ins[i].reg] = convert(ins[i].val,reg)
			break
		case "sub":
			reg[ins[i].reg] -= convert(ins[i].val,reg)
			break
		case "mul":
			reg[ins[i].reg] *= convert(ins[i].val,reg)
			c++
			break
		case "jnz":
			if convert(ins[i].reg,reg) != 0 {
				i += convert(ins[i].val,reg) - 1
			}
			break
		default:
			panic("Error: Unknown Command")
		}
	}
	return 
}
func parseInstructions(path string) (ins []instruction) {
	file,_ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 3 {
			ins = append(ins, instruction{parts[0],parts[1], parts[2]})
		} else {
			ins = append(ins, instruction{parts[0],parts[1], ""})
		}
	}
	return
}
func convert(value string, reg map[string]int) (int) {
	numVal,err := strconv.Atoi(value)
	if err != nil {
		return reg[value] //value is another register
	} else {
		return numVal
	}
}