package main

import "fmt"

var lcss_ = struct {
	matrix     [][]element
	str1, str2 string
}{
	matrix: make([][]element, 0),
}

const (
	stay = iota
	up
	left
	diagonal
)

type element struct {
	magnitude int
	direction int
}

func initialiseLcss(str1, str2 string) {
	m, n := len(str1)+1, len(str2)+1
	for _ = range m {
		lcss_.matrix = append(lcss_.matrix, make([]element, n))
	}
	lcss_.str1, lcss_.str2 = str1, str2
}

func LCS(x, y string) {
	m, n := len(x), len(y)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				lcss_.matrix[i][j] = element{magnitude: 1 + lcss_.matrix[i-1][j-1].magnitude, direction: diagonal}
			} else if lcss_.matrix[i][j-1].magnitude > lcss_.matrix[i-1][j].magnitude {
				lcss_.matrix[i][j] = lcss_.matrix[i][j-1]
			} else {
				lcss_.matrix[i][j] = lcss_.matrix[i-1][j]
			}
		}
	}
}

func traceLcs() {
	lastRowIndex := len(lcss_.matrix) - 1
	lastColIndex := len(lcss_.matrix[0]) - 1
	walker := walker{bs: make([]byte, 0), row: lastRowIndex, col: lastColIndex}
	walker.tracePath()
}

type walker struct {
	bs       []byte
	row, col int
}

func (w *walker) collect() {
	tmp := make([]byte, len((*w).bs)+1)
	tmp[0] = lcss_.str1[w.row]
	copy(tmp[1:], (*w).bs)
	(*w).bs = tmp
}

func (w *walker) tracePath() {
	for {
		if to := lcss_.matrix[w.row][w.col].direction; to != stay {
			w.walk(to)
		}
		break
	}
}

func (w *walker) walk(direction int) {
	switch direction {
	case up:
		w.row--
	case left:
		w.col--
	case diagonal:

		w.row--
		w.col--
	default:
		//stay put
	}
}

func main() {
	x, y := "GBDE", "ACDD"
	initialiseLcss(x, y)
	LCS(x, y)
	print(x, y)
}

func print(x, y string) {
	px := "e" + x
	py := " e" + y
	cols := len(lcss_.matrix[0])
	for j := range cols + 1 {
		fmt.Print(string(py[j]) + " ")
	}
	fmt.Println()
	for i, row := range lcss_.matrix {
		fmt.Print(string(px[i]) + " ")
		for _, elem := range row {
			fmt.Print(elem.magnitude)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
