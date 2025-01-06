package withoutmemoization

var plane_ = struct {
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

func initialisePlane(str1, str2 string) {
	m, n := len(str1)+1, len(str2)+1
	for _ = range m {
		plane_.matrix = append(plane_.matrix, make([]element, n))
	}
	plane_.str1, plane_.str2 = str1, str2
}

func LCS(x, y string) string {
	initialisePlane(x, y)
	m, n := len(x), len(y)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				plane_.matrix[i][j] = element{magnitude: 1 + plane_.matrix[i-1][j-1].magnitude, direction: diagonal}
			} else if plane_.matrix[i][j-1].magnitude <= plane_.matrix[i-1][j].magnitude {
				plane_.matrix[i][j] = element{magnitude: plane_.matrix[i-1][j].magnitude, direction: up}
			} else {
				plane_.matrix[i][j] = element{magnitude: plane_.matrix[i][j-1].magnitude, direction: left}
			}
		}
	}
	return Lcs()
}

func Lcs() string {
	lastRowIndex := len(plane_.matrix) - 1
	lastColIndex := len(plane_.matrix[0]) - 1
	walker := walker{bs: make([]byte, 0), row: lastRowIndex, col: lastColIndex}
	return walker.tracePath()
}

type walker struct {
	bs       []byte
	row, col int
}

func (w *walker) collect() {
	if len(w.bs) == 0 {
		w.bs = append(w.bs, plane_.str1[w.row-1])
		return
	}
	tmp := make([]byte, len(w.bs)+1)
	tmp[0] = plane_.str1[w.row-1]
	copy(tmp[1:], w.bs)
	w.bs = tmp
}

func (w *walker) tracePath() string {
	for {
		to := plane_.matrix[w.row][w.col].direction
		if to == stay {
			break
		}
		w.walk(to)
	}
	return string(w.bs)
}

func (w *walker) walk(direction int) {
	switch direction {
	case up:
		w.row--
	case left:
		w.col--
	case diagonal:
		w.collect()
		w.row--
		w.col--
	default:
		//stay put
	}
}
