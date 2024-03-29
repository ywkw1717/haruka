package rubik

import (
	"fmt"
	"strings"
)

const (
	WHITE = iota
	ORANGE
	GREEEN
	RED
	BLUE
	YELLOW
)

type Cube struct {
	U [][]int
	L [][]int
	F [][]int
	R [][]int
	B [][]int
	D [][]int
}

func NewCube() *Cube {
	cube := new(Cube)
	cube.U = makeCubeData(WHITE)
	cube.L = makeCubeData(ORANGE)
	cube.F = makeCubeData(GREEEN)
	cube.R = makeCubeData(RED)
	cube.B = makeCubeData(BLUE)
	cube.D = makeCubeData(YELLOW)
	return cube
}

func makeCubeData(color int) [][]int {
	data := make([][]int, 3)
	for i := range data {
		data[i] = []int{color, color, color}
	}
	return data
}

func Copy(cube *Cube) *Cube {
	dup := NewCube()

	deepCopyTwoDimensionalArray(cube.U, dup.U)
	deepCopyTwoDimensionalArray(cube.L, dup.L)
	deepCopyTwoDimensionalArray(cube.F, dup.F)
	deepCopyTwoDimensionalArray(cube.R, dup.R)
	deepCopyTwoDimensionalArray(cube.B, dup.B)
	deepCopyTwoDimensionalArray(cube.D, dup.D)

	return dup
}

func deepCopyTwoDimensionalArray(from [][]int, to [][]int) {
	for i, v := range from {
		copy(to[i], v)
	}
}

func (cube *Cube) PrintState() {
	fmt.Println("       -------                    ")
	fmt.Printf("       |%d|%d|%d|                    \n", cube.U[0][0], cube.U[0][1], cube.U[0][2])
	fmt.Printf("       |%d|%d|%d|                    \n", cube.U[1][0], cube.U[1][1], cube.U[1][2])
	fmt.Printf("       |%d|%d|%d|                    \n", cube.U[2][0], cube.U[2][1], cube.U[2][2])

	fmt.Println(" ------|-----|------------  ")

	fmt.Printf(" |%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|  \n",
		cube.L[0][0], cube.L[0][1], cube.L[0][2],
		cube.F[0][0], cube.F[0][1], cube.F[0][2],
		cube.R[0][0], cube.R[0][1], cube.R[0][2],
		cube.B[0][0], cube.B[0][1], cube.B[0][2],
	)

	fmt.Printf(" |%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|  \n",
		cube.L[1][0], cube.L[1][1], cube.L[1][2],
		cube.F[1][0], cube.F[1][1], cube.F[1][2],
		cube.R[1][0], cube.R[1][1], cube.R[1][2],
		cube.B[1][0], cube.B[1][1], cube.B[1][2],
	)

	fmt.Printf(" |%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|%d|  \n",
		cube.L[2][0], cube.L[2][1], cube.L[2][2],
		cube.F[2][0], cube.F[2][1], cube.F[2][2],
		cube.R[2][0], cube.R[2][1], cube.R[2][2],
		cube.B[2][0], cube.B[2][1], cube.B[2][2],
	)

	fmt.Println(" ------|-----|------------  ")

	fmt.Printf("       |%d|%d|%d|                    \n", cube.D[0][0], cube.D[0][1], cube.D[0][2])
	fmt.Printf("       |%d|%d|%d|                    \n", cube.D[1][0], cube.D[1][1], cube.D[1][2])
	fmt.Printf("       |%d|%d|%d|                    \n", cube.D[2][0], cube.D[2][1], cube.D[2][2])

	fmt.Println("       -------                    ")
}

var rotationMap map[string](func(*Cube)) = map[string](func(*Cube)){
	"R":  (*Cube).RRotation,
	"R'": (*Cube).RPrimeRotation,
	"R2": (*Cube).R2Rotation,
	"L":  (*Cube).LRotation,
	"L'": (*Cube).LPrimeRotation,
	"L2": (*Cube).L2Rotation,
	"U":  (*Cube).URotation,
	"U'": (*Cube).UPrimeRotation,
	"U2": (*Cube).U2Rotation,
	"D":  (*Cube).DRotation,
	"D'": (*Cube).DPrimeRotation,
	"D2": (*Cube).D2Rotation,
	"F":  (*Cube).FRotation,
	"F'": (*Cube).FPrimeRotation,
	"F2": (*Cube).F2Rotation,
	"B":  (*Cube).BRotation,
	"B'": (*Cube).BPrimeRotation,
	"B2": (*Cube).B2Rotation,
}

func Rotation(cube *Cube, s string) {
	commands := strings.Split(s, " ")

	for _, v := range commands {
		rotationMap[v](cube)
	}
}

// from: dup, to: cube
func baseRotation(from [][]int, to [][]int) {
	to[0][2] = from[0][0]
	to[2][2] = from[0][2]
	to[2][0] = from[2][2]
	to[0][0] = from[2][0]

	to[1][2] = from[0][1]
	to[2][1] = from[1][2]
	to[1][0] = from[2][1]
	to[0][1] = from[1][0]
}

func basePrimeRotation(from [][]int, to [][]int) {
	to[2][0] = from[0][0]
	to[2][2] = from[2][0]
	to[0][2] = from[2][2]
	to[0][0] = from[0][2]

	to[1][0] = from[0][1]
	to[2][1] = from[1][0]
	to[1][2] = from[2][1]
	to[0][1] = from[1][2]
}

func (cube *Cube) RRotation() {
	dup := Copy(cube)

	baseRotation(dup.R, cube.R)

	cube.U[0][2] = dup.F[0][2]
	cube.U[1][2] = dup.F[1][2]
	cube.U[2][2] = dup.F[2][2]

	cube.F[0][2] = dup.D[0][2]
	cube.F[1][2] = dup.D[1][2]
	cube.F[2][2] = dup.D[2][2]

	cube.D[2][2] = dup.B[0][0]
	cube.D[1][2] = dup.B[1][0]
	cube.D[0][2] = dup.B[2][0]

	cube.B[0][0] = dup.U[2][2]
	cube.B[1][0] = dup.U[1][2]
	cube.B[2][0] = dup.U[0][2]
}

func (cube *Cube) RPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.R, cube.R)

	cube.U[0][2] = dup.B[2][0]
	cube.U[1][2] = dup.B[1][0]
	cube.U[2][2] = dup.B[0][0]

	cube.F[0][2] = dup.U[0][2]
	cube.F[1][2] = dup.U[1][2]
	cube.F[2][2] = dup.U[2][2]

	cube.D[0][2] = dup.F[0][2]
	cube.D[1][2] = dup.F[1][2]
	cube.D[2][2] = dup.F[2][2]

	cube.B[0][0] = dup.D[2][2]
	cube.B[1][0] = dup.D[1][2]
	cube.B[2][0] = dup.D[0][2]
}

func (cube *Cube) R2Rotation() {
	cube.RRotation()
	cube.RRotation()
}

func (cube *Cube) LRotation() {
	dup := Copy(cube)

	baseRotation(dup.L, cube.L)

	cube.U[0][0] = dup.B[2][2]
	cube.U[1][0] = dup.B[1][2]
	cube.U[2][0] = dup.B[0][2]

	cube.B[2][2] = dup.D[0][0]
	cube.B[1][2] = dup.D[1][0]
	cube.B[0][2] = dup.D[2][0]

	cube.D[0][0] = dup.F[0][0]
	cube.D[1][0] = dup.F[1][0]
	cube.D[2][0] = dup.F[2][0]

	cube.F[0][0] = dup.U[0][0]
	cube.F[1][0] = dup.U[1][0]
	cube.F[2][0] = dup.U[2][0]
}

func (cube *Cube) LPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.L, cube.L)

	cube.U[0][0] = dup.F[0][0]
	cube.U[1][0] = dup.F[1][0]
	cube.U[2][0] = dup.F[2][0]

	cube.F[0][0] = dup.D[0][0]
	cube.F[1][0] = dup.D[1][0]
	cube.F[2][0] = dup.D[2][0]

	cube.D[2][0] = dup.B[0][2]
	cube.D[1][0] = dup.B[1][2]
	cube.D[0][0] = dup.B[2][2]

	cube.B[0][2] = dup.U[2][0]
	cube.B[1][2] = dup.U[1][0]
	cube.B[2][2] = dup.U[0][0]
}

func (cube *Cube) L2Rotation() {
	cube.LRotation()
	cube.LRotation()
}

func (cube *Cube) URotation() {
	dup := Copy(cube)

	baseRotation(dup.U, cube.U)

	cube.F[0][0] = dup.R[0][0]
	cube.F[0][1] = dup.R[0][1]
	cube.F[0][2] = dup.R[0][2]

	cube.R[0][0] = dup.B[0][0]
	cube.R[0][1] = dup.B[0][1]
	cube.R[0][2] = dup.B[0][2]

	cube.B[0][0] = dup.L[0][0]
	cube.B[0][1] = dup.L[0][1]
	cube.B[0][2] = dup.L[0][2]

	cube.L[0][0] = dup.F[0][0]
	cube.L[0][1] = dup.F[0][1]
	cube.L[0][2] = dup.F[0][2]
}

func (cube *Cube) UPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.U, cube.U)

	cube.F[0][0] = dup.L[0][0]
	cube.F[0][1] = dup.L[0][1]
	cube.F[0][2] = dup.L[0][2]

	cube.L[0][0] = dup.B[0][0]
	cube.L[0][1] = dup.B[0][1]
	cube.L[0][2] = dup.B[0][2]

	cube.B[0][0] = dup.R[0][0]
	cube.B[0][1] = dup.R[0][1]
	cube.B[0][2] = dup.R[0][2]

	cube.R[0][0] = dup.F[0][0]
	cube.R[0][1] = dup.F[0][1]
	cube.R[0][2] = dup.F[0][2]
}

func (cube *Cube) U2Rotation() {
	cube.URotation()
	cube.URotation()
}

func (cube *Cube) DRotation() {
	dup := Copy(cube)

	baseRotation(dup.D, cube.D)

	cube.F[2][0] = dup.L[2][0]
	cube.F[2][1] = dup.L[2][1]
	cube.F[2][2] = dup.L[2][2]

	cube.L[2][0] = dup.B[2][0]
	cube.L[2][1] = dup.B[2][1]
	cube.L[2][2] = dup.B[2][2]

	cube.B[2][0] = dup.R[2][0]
	cube.B[2][1] = dup.R[2][1]
	cube.B[2][2] = dup.R[2][2]

	cube.R[2][0] = dup.F[2][0]
	cube.R[2][1] = dup.F[2][1]
	cube.R[2][2] = dup.F[2][2]
}

func (cube *Cube) DPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.D, cube.D)

	cube.F[2][0] = dup.R[2][0]
	cube.F[2][1] = dup.R[2][1]
	cube.F[2][2] = dup.R[2][2]

	cube.R[2][0] = dup.B[2][0]
	cube.R[2][1] = dup.B[2][1]
	cube.R[2][2] = dup.B[2][2]

	cube.B[2][0] = dup.L[2][0]
	cube.B[2][1] = dup.L[2][1]
	cube.B[2][2] = dup.L[2][2]

	cube.L[2][0] = dup.F[2][0]
	cube.L[2][1] = dup.F[2][1]
	cube.L[2][2] = dup.F[2][2]
}

func (cube *Cube) D2Rotation() {
	cube.DRotation()
	cube.DRotation()
}

func (cube *Cube) FRotation() {
	dup := Copy(cube)

	baseRotation(dup.F, cube.F)

	cube.R[0][0] = dup.U[2][0]
	cube.R[1][0] = dup.U[2][1]
	cube.R[2][0] = dup.U[2][2]

	cube.D[0][2] = dup.R[0][0]
	cube.D[0][1] = dup.R[1][0]
	cube.D[0][0] = dup.R[2][0]

	cube.L[2][2] = dup.D[0][2]
	cube.L[1][2] = dup.D[0][1]
	cube.L[0][2] = dup.D[0][0]

	cube.U[2][0] = dup.L[2][2]
	cube.U[2][1] = dup.L[1][2]
	cube.U[2][2] = dup.L[0][2]
}

func (cube *Cube) FPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.F, cube.F)

	cube.R[0][0] = dup.D[0][2]
	cube.R[1][0] = dup.D[0][1]
	cube.R[2][0] = dup.D[0][0]

	cube.D[0][2] = dup.L[2][2]
	cube.D[0][1] = dup.L[1][2]
	cube.D[0][0] = dup.L[0][2]

	cube.L[2][2] = dup.U[2][0]
	cube.L[1][2] = dup.U[2][1]
	cube.L[0][2] = dup.U[2][2]

	cube.U[2][0] = dup.R[0][0]
	cube.U[2][1] = dup.R[1][0]
	cube.U[2][2] = dup.R[2][0]
}

func (cube *Cube) F2Rotation() {
	cube.FRotation()
	cube.FRotation()
}

func (cube *Cube) BRotation() {
	dup := Copy(cube)

	baseRotation(dup.B, cube.B)

	cube.R[0][2] = dup.D[2][2]
	cube.R[1][2] = dup.D[2][1]
	cube.R[2][2] = dup.D[2][0]

	cube.D[2][2] = dup.L[2][0]
	cube.D[2][1] = dup.L[1][0]
	cube.D[2][0] = dup.L[0][0]

	cube.L[2][0] = dup.U[0][0]
	cube.L[1][0] = dup.U[0][1]
	cube.L[0][0] = dup.U[0][2]

	cube.U[0][0] = dup.R[0][2]
	cube.U[0][1] = dup.R[1][2]
	cube.U[0][2] = dup.R[2][2]
}

func (cube *Cube) BPrimeRotation() {
	dup := Copy(cube)

	basePrimeRotation(dup.B, cube.B)

	cube.R[0][2] = dup.U[0][0]
	cube.R[1][2] = dup.U[0][1]
	cube.R[2][2] = dup.U[0][2]

	cube.U[0][0] = dup.L[2][0]
	cube.U[0][1] = dup.L[1][0]
	cube.U[0][2] = dup.L[0][0]

	cube.L[2][0] = dup.D[2][2]
	cube.L[1][0] = dup.D[2][1]
	cube.L[0][0] = dup.D[2][0]

	cube.D[2][2] = dup.R[0][2]
	cube.D[2][1] = dup.R[1][2]
	cube.D[2][0] = dup.R[2][2]
}

func (cube *Cube) B2Rotation() {
	cube.BRotation()
	cube.BRotation()
}
