// Package pascal contains a function to calulate pascal's triangle
//
//
// In Pascal's Triangle each number is computed by adding the numbers to
// the right and left of the current position in the previous row.
//
// ```plain
//     1
//    1 1
//   1 2 1
//  1 3 3 1
// 1 4 6 4 1
// # ... etc
// ```
// See http://mathworld.wolfram.com/PascalsTriangle.html for a more
// in depth overview.
package pascal

var testVersion = 1

// Triangle computes Pascal's Triangle with the given number of rows.
func Triangle(rows int) [][]int {
	t := make([][]int, rows)
	for i := 0; i < rows; i++ {
		t[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			t[i][j] = calc(t, i, j)
		}
	}
	return t
}

func calc(t [][]int, row, col int) int {
	switch col {
	case 0, row:
		return 1
	case 1, row - 1:
		return row
	default:
		return t[row-1][col-1] + t[row-1][col]
	}

}
