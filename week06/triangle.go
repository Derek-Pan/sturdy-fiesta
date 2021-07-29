package week06

import "math"

// 三角形的最小路径之和

/*
定义状态:
triangle: 每个路径的数值大小
i: 当前在第几层
j: 选取i层的哪一个数
目标: 路径之和最小

状态转移方程:
				当 j = 0 时: F[i-1][j] + triangle[i][j]
	F(i, j) = 	当 j = triangle[i].size - 1 时: F[i-1][j-1] + triangle[i][j]
				当 0 < j < triangle[i].size - 1 时: math.Min(F[i-1][j-1] + triangle[i][j], F[i-1][j] + triangle[i][j])

目标:
	ans = min(F(n, j)) 0 <= j < triangle[n].size
*/

func minimumTotal(triangle [][]int) int {
	var size int = len(triangle)
	var f [][]int = make([][]int, size, size)
	for idx := range f {
		f[idx] = make([]int, len(triangle[size-1]), len(triangle[size-1]))
		for i := range f[idx] {
			f[idx][i] = 1 << 32
		}
	}
	// 初始化边界
	f[0][0] = triangle[0][0]
	// 遍历条件
	for i := 1; i < size; i++ {
		// 遍历所有决策
		for j := 0; j < len(triangle[i]); j++ {
			// 状态转移方程的实现
			if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = int(math.Min(float64(f[i-1][j])+float64(triangle[i][j]), float64(f[i-1][j-1])+float64(triangle[i][j])))
			}
		}
	}
	ans := 1 << 32
	for _, val := range f[size-1] {
		ans = int(math.Min(float64(ans), float64(val)))
	}
	return ans
}
