package LinearAlgebra

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

func Transpose(x [][]float64) [][]float64{
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

func AddMat (x [][]float64, y [][]float64) [][]float64{
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

func ScalMult (x [][] float64, y float64) [][]float64{
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

func DotProd (x [][]float64, y [][]float64)float64{
	return 0
	}
