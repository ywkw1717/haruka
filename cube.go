package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

/*
##########################################
#          Cube State                    #
#                                        #
#          ----------                    #
#          |U1|U2|U3|                    #
#          |U4|U5|U6|                    #
#          |U7|U8|U9|                    #
# ---------|--------|------------------  #
# |L1|L2|L3|F1|F2|F3|R1|R2|R3|B1|B2|B3|  #
# |L4|L5|L6|F4|F5|F6|R4|R5|R6|B4|B5|B6|  #
# |L7|L8|L9|F7|F8|F9|R7|R8|R9|B7|B8|B9|  #
# ---------|--------|------------------  #
#          |D1|D2|D3|                    #
#          |D4|D5|D6|                    #
#          |D7|D8|D9|                    #
#          ----------                    #
#                                        #
##########################################
*/
const (
	WHITE = iota
	ORANGE
	GREEEN
	RED
	BLUE
	YELLOW
)

type U *mat.Dense
type L *mat.Dense
type F *mat.Dense
type R *mat.Dense
type B *mat.Dense
type D *mat.Dense

type Cube struct {
	u U
	l L
	f F
	r R
	b B
	d D
}

func NewCube() *Cube {
	cube := new(Cube)
	cube.u = mat.NewDense(3, 3, makeCubeData(WHITE))
	cube.l = mat.NewDense(3, 3, makeCubeData(ORANGE))
	cube.f = mat.NewDense(3, 3, makeCubeData(GREEEN))
	cube.r = mat.NewDense(3, 3, makeCubeData(RED))
	cube.b = mat.NewDense(3, 3, makeCubeData(BLUE))
	cube.d = mat.NewDense(3, 3, makeCubeData(YELLOW))
	return cube
}

func makeCubeData(color int) []float64 {
	data := make([]float64, 9)
	for i := range data {
		data[i] = float64(color)
	}
	return data
}

func main() {
	cube := NewCube()
	fmt.Println(cube.u)
	fmt.Println(cube.l)
	fmt.Println(cube.f)
	fmt.Println(cube.r)
	fmt.Println(cube.b)
	fmt.Println(cube.d)
}
