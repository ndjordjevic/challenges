package ctci_RobotInGrid

import "fmt"

type Point struct {
	row int
	col int
}

func main() {
	r := getPath([][]bool{{true, true, true}, {true, false, true}, {true, false, true}})
	fmt.Println(r)
}
func getPath(maze [][]bool) []Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}

	path := make([]Point, 0)
	if getPathRec(maze, len(maze)-1, len(maze[0])-1, &path) {
		return path
	}

	return nil
}

func getPathRec(maze [][]bool, row int, col int, path *[]Point) bool {
	if col < 0 || row < 0 || !maze[row][col] {
		return false
	}

	isAtOrigin := row == 0 && col == 0

	if isAtOrigin || getPathRec(maze, row, col-1, path) || getPathRec(maze, row-1, col, path) {
		p := Point{
			row: row,
			col: col,
		}
		*path = append(*path, p)
		return true
	}

	return false
}
