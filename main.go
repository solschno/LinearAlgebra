package LinearAlgebra

import "fmt"

type twoDMatrix struct {
	rows int
	columns int
	mat [][]float64
}

func NewMatrix (rows, columns int) twoDMatrix {
	var a twoDMatrix
	a.rows = rows
	a.columns = columns
	rtemp := make([][]float64, rows)
	for i := 0; i < rows; i++{
		rtemp[i] = make([]float64,columns)
	}
	a.mat = rtemp
	return a
}

func transpose(x [][]float64) [][]float64{
	a := make([][]float64,len(x[0]))
	for i := 0; i < len(a); i++{
		a[i] = make([]float64,len(x))
	}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[0]); j++{
			a[i][j] = x[j][i]
		}
	}
	return a
}

func addMat (x [][]float64, y [][]float64) [][]float64{
	a := make([][]float64,len(x))
	for i := 0; i < len(a); i++{
		a[i] = make([]float64,len(x[0]))
	}
	for i := 0; i < len(x); i++{
		for j := 0; j < len(x[0]); j++{
			a[i][j] = x[i][j] +y[i][j]
		}
	}
	return a
}

func scalMult (x [][] float64, y float64) [][]float64{
	a := make([][]float64,len(x))
	for i := 0; i < len(a); i++{
		a[i] = make([]float64,len(x[0]))
	}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(a[0]); j++{
			a[i][j] = x[i][j] * y
		}
	}
	return a
}

func dotProd (x [][]float64, y [][]float64)float64{
	return 0
	}

func main() {
	a := make([][]float64,2)
	a[0] = make([]float64,3)
	a[1] = make([]float64,3)
	a[0][0] = 1
	a[0][1] = 2
	a[0][2] = 3
	a[1][0] = 4
	a[1][1] = 5
	a[1][2] = 6
	c := dotProd (a,a)
	b := transpose(a)
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(addMat(a,a))
	fmt.Println(scalMult(a,4))
}