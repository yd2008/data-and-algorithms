package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze()
	steps := walk(maze, point{i: 0, j: 0}, point{i: 5, j: 4})
	for i := range steps {
		fmt.Printf("%v\n", steps[i])
	}
}

func readMaze() [][]int {
	file, err := os.Open("src/guang_du_you_xian/maze.in")
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Printf("%v\n", maze[i])
	}
	fmt.Println()
	return maze
}

type point struct {
	i, j int
}

var (
	// 上左下右四个方向
	dirs = [4]point{{i: 0, j: -1}, {i: -1, j: 0}, {i: 0, j: 1}, {i: 1, j: 0}}
)

func (p point) add(r point) point {
	return point{i: p.i + r.i, j: p.j + r.j}
}

// walk 返回走到终点需要的步数和路径图
func walk(maze [][]int, begin, end point) [][]int {
	// 维护一个代表当前格子最近的步数二维slice
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
		//fmt.Printf("%v\n", steps[i])
	}

	// 创建一个一探索到的点队列
	points := []point{begin}

	for len(points) > 0 {
		// 把第一个取出来
		cur := points[0]
		points = points[1:]

		// 找到终点
		if cur == end {
			break
		}

		// 开始探索
		for _, dir := range dirs {
			// 计算出当前点
			next := cur.add(dir)
			// 没有越界
			if next.i < 0 || next.i > len(maze)-1 {
				continue
			}
			if next.j < 0 || next.j > len(maze[next.i])-1 {
				continue
			}
			// maze中的下一个点是0 即不是墙 1
			if maze[next.i][next.j] == 1 {
				continue
			}
			// 在steps中数字不为0 说明已经走过
			if steps[next.i][next.j] != 0 {
				continue
			}
			// 不是起点
			if next == begin {
				continue
			}
			// 探索到新路
			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			// 添加进队列
			points = append(points, next)
		}
	}

	return steps
}
