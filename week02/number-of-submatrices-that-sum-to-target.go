package week02

// 元素和为目标值的子矩阵数量
/*
主体思路: 题目是要求子矩阵的和,首先
1. 计算原数组的前缀和
2. 固定 较小的点(x1, y1) 遍历(x2, y2)求出所有的可能性
3. 当(x1, y1)超过数组的边界时,结束
*/

type NumberMatrix struct {
	preSum [][]int
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	xLength := len(matrix)
	yLength := len(matrix[0])
	ps := &NumberMatrix{
		preSum: make([][]int, xLength, xLength),
	}
	ans := 0
	// 计算二维数组的前缀和
	for x := 0; x < xLength; x++ {
		ps.preSum[x] = make([]int, yLength, yLength)
		for y := 0; y < yLength; y++ {
			ps.preSum[x][y] = ps.getSum(x-1, y) + ps.getSum(x, y-1) - ps.getSum(x-1, y-1) + matrix[x][y]
		}
	}
	// 对左上边界点进行遍历
	for i := 0; i < xLength; i++ {
		for j := 0; j < yLength; j++ {
			ans += ps.subMatrixSum(i, j, target)
		}
	}
	return ans
}

func (this *NumberMatrix) getSum(x, y int) int {
	if x >= 0 && y >= 0 {
		return this.preSum[x][y]
	}
	return 0
}

func (this *NumberMatrix) subMatrixSum(x1, y1, target int) (ans int) {
	// 遍历右边界所有的可能性
	for i := x1; i < len(this.preSum); i++ {
		for j := y1; j < len(this.preSum[0]); j++ {
			subSum := this.getSum(i, j) - this.getSum(i, y1-1) - this.getSum(x1-1, j) + this.getSum(x1-1, y1-1)
			if subSum == target {
				ans++
			}
		}
	}
	return ans
}












//利用map进行解题思路暂时还没有理解,尴尬了!!!
// type NumberMatrix struct {
//     preSum [][]int
// }

// func numSubmatrixSumTarget(matrix [][]int, target int) int {
//     lineLength := len(matrix)
//     columeLength := len(matrix[0])
//     ans := 0
//     ps := &NumberMatrix{
//         preSum: make([][]int, lineLength, lineLength),
//     }
//     // 计算二维数组的前缀和
//     for x := 0; x < lineLength; x++ {
//         ps.preSum[x] = make([]int, columeLength, columeLength)
//         for y:=0;y <columeLength; y++ {
//             ps.preSum[x][y] = ps.getSum(x-1, y) + ps.getSum(x, y-1) - ps.getSum(x-1, y-1) + matrix[x][y]
//         }
//     }

//     for top := 0; top < lineLength; top++ {
//         for bottom := top; bottom < lineLength; bottom++ {
//             cntMap := make(map[int]int)
//             cntMap[0] = 1
//             for cloumn := 0; cloumn < columeLength; cloumn++ {
//                 cur := ps.getSum(bottom, cloumn) - ps.getSum(top-1, cloumn)   
//                 ans += cntMap[cur - target]
//                 cntMap[cur]++
//             }
//         }
//     }
//     return ans
// }

// func (this *NumberMatrix) getSum(x, y int) int {
//     if x >= 0 && y >= 0 {
//         return this.preSum[x][y]
//     }
//     return 0
// }
