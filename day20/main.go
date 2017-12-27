package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

type vector struct {
	x int
	y int
	z int
}

type particle struct {
	pos vector
	vel vector
	acc vector
}

func step(p *particle) {
	p.vel.x += p.acc.x
	p.vel.y += p.acc.y
	p.vel.z += p.acc.z
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func parseAccel() {
	min := 11.0
	file,_ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	index := 0
	for i := 0;scanner.Scan(); i++{
		line := scanner.Text()
		split := strings.Split(line, "a=")
		split[1] = strings.TrimRight( strings.TrimLeft(split[1],"<"),">")
		partAccel := strings.Split(split[1],",")
		x,_ := strconv.Atoi(partAccel[0])
		y,_ := strconv.Atoi(partAccel[1])
		z,_ := strconv.Atoi(partAccel[2])
		a := math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))

		if a <= min {
			fmt.Println(line)
			min = a
			index = i
		}

	}
	fmt.Println(index)
}
func distance(p particle) (int) {
	return p.pos.x + p.pos.y + p.pos.z
}
func simulateParticles() {

}

//in the long run, the particle with the lowest acceleration will stay closest to 0,0,0
func main() {
	parseAccel()
	//there are only 2 particles with the lowest acceleration (2), 
	//so just manually compute their initial velocity to find particle that stays closest to origin.
}