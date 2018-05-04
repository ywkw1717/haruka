package cube

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

func (cube *Cube) rRotation() {
	dup := Copy(cube)

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

func (cube *Cube) rPrimeRotation() {
	dup := Copy(cube)

	cube.U[0][2] = dup.B[2][0]
	cube.U[1][2] = dup.B[1][0]
	cube.U[2][2] = dup.B[0][0]

	cube.F[0][2] = dup.U[0][2]
	cube.F[1][2] = dup.U[1][2]
	cube.F[2][2] = dup.U[2][2]

	cube.D[0][2] = dup.F[0][0]
	cube.D[1][2] = dup.F[1][0]
	cube.D[2][2] = dup.F[2][0]

	cube.B[0][0] = dup.D[0][2]
	cube.B[1][0] = dup.D[1][2]
	cube.B[2][0] = dup.D[2][2]
}

func (cube *Cube) r2Rotation() {
	cube.rRotation()
	cube.rRotation()
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

func (cube *Cube) printState() {
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
