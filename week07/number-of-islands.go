package week07

/*
题目: 岛屿数量(并查集求解)
思路: 将所有相连的岛屿都并入一个集合, 有多少个不同的集合就有多少个岛屿的数量
*/

type DisJoinSet struct {
	fa []int
}

func (djs *DisJoinSet) find(x int) int {
	if x == djs.fa[x] {
		return x
	}
	djs.fa[x] = djs.find(djs.fa[x])
	return djs.fa[x]
}

func (djs *DisJoinSet) unionSet(x, y int) {
	x = djs.find(x)
	y = djs.find(y)
	if x != y {
		djs.fa[x] = y
	}
}

func numIslands(grid [][]byte) int {
	var m int = len(grid)
	var n int = len(grid[0])
	var djs *DisJoinSet = &DisJoinSet{
		fa: make([]int, m*n+1, m*n+1),
	}
	for i := 0; i < m*n+1; i++ {
		djs.fa[i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				djs.unionSet(i*n+j+1, 0)
			} else {
				// 将相连的岛屿并入同一个集合
				if j > 0 && grid[i][j-1] == '1' {
					djs.unionSet(i*n+j+1, i*n+j)
				}
				if i > 0 && grid[i-1][j] == '1' {
					djs.unionSet(i*n+j+1, (i-1)*n+j+1)
				}
			}
		}
	}
	var ans int = 0
	// 计算不同集合数量
	for i := 0; i <= m*n; i++ {
		if djs.fa[i] == i && djs.fa[i] != 0 {
			ans++
		}
	}
	return ans
}
