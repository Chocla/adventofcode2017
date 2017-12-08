package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
type operation struct {
	reg string
	inc bool //+= if true, -= if false
	incVal, operVal int
	operReg, oper string
}

const INPUT_PATH = "input.txt"

func main(){
	reg := make(map[string]int)
	operationSlice := make([]operation,0)
	max := 0

	file,_ := os.Open(INPUT_PATH)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	//get list of operations from input file
	for scanner.Scan() {
		line := scanner.Text()
		operationSlice = append(operationSlice, parseOperation(line))
	}

	//init map values for all registries
	for i := range operationSlice {
		reg[operationSlice[i].reg] = 0
	}

	//evalute operations
	for i := range operationSlice {
		reg = evaluate(reg, operationSlice[i])

		//part two: maximum registry value at any point in computation
		for i := range reg {
			if reg[i] > max {
				 max = reg[i]
			}
		}
	}

	fmt.Println(max)
}

func evaluate(reg map[string]int, o operation) (map[string]int) {
	var expression bool
	left := reg[o.operReg]
	right := o.operVal

	switch o.oper {
	case "<":  expression = left <  right 
		break
	case ">":  expression = left >  right
		break
	case ">=": expression = left >= right
		break
	case "<=": expression = left <= right
		break
	case "==": expression = left == right
		break
	case "!=": expression = left != right
		break

	}
	if expression {
		if o.inc {
			reg[o.reg] += o.incVal
		} else {
			reg[o.reg] -= o.incVal
		}
	}

	return reg
}

func parseOperation(input string)(o operation) {
	inputSlice := strings.Split(input, " ")
	
	o.reg, o.operReg, o.oper = inputSlice[0],inputSlice[4],inputSlice[5]
	
	if inputSlice[1] == "inc" {
		o.inc = true
	}
	o.incVal ,_ = strconv.Atoi(inputSlice[2])
	o.operVal,_ = strconv.Atoi(inputSlice[6])

	return
}

//instructions
	//0 - reg
	//1 - inc or dec
	//2 - int value
	//3 - unimportant
	//4,5,6 - boolean expression
		//4 - reg
		//5 - binary operator
		//6 - test value
// >, <, >=, <=, ==, !=

