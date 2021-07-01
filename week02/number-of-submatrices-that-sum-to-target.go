package week02

// 元素和为目标值的子矩阵数量
/*
主体思路: 题目是要求子矩阵的和,首先
1. 计算原数组的前缀和
2. 遍历两个点的所有可能性
(暴力求解法)
*/

//MatrixPreSum 矩阵前缀和
type MatrixPreSum struct {
	preSum [][]int //前缀和数组
}

// 获取某个坐标的值
func (recviver *MatrixPreSum) getSum(r, c int) int {
	if r >= 0 && c >= 0 {
		return recviver.preSum[r][c]
	}
	return 0
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	row := len(matrix)
	col := len(matrix[0])
	mps := &MatrixPreSum{
		preSum: make([][]int, row, row),
	}
	ans := 0
	// 预计算二维数组的前缀和
	for x := 0; x < row; x++ {
		mps.preSum[x] = make([]int, col, col)
		for y := 0; y < col; y++ {
			mps.preSum[x][y] = mps.getSum(x-1, y) + mps.getSum(x, y-1) - mps.getSum(x-1, y-1) + matrix[x][y]
		}
	}
	// 对左上边界点进行遍历
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans += mps.subMatrixSum(i, j, target)
		}
	}
	return ans
}

func (recviver *MatrixPreSum) subMatrixSum(row, col, target int) (ans int) {
	// 遍历右端点
	for i := row; i < len(recviver.preSum); i++ {
		for j := col; j < len(recviver.preSum[0]); j++ {
			subSum := recviver.getSum(i, j) - recviver.getSum(i, col-1) - recviver.getSum(row-1, j) + recviver.getSum(row-1, col-1)
			if subSum == target {
				ans++
			}
		}
	}
	return ans
}

// 利用map进行解题思路暂时还没有完全理解,尴尬了!!!
// 将二维数组求子矩阵和的问题 => 一维数组求子数组和的问题

func numSubmatrixSumTargetPlan2(matrix [][]int, target int) int {
	row := len(matrix)
	col := len(matrix[0])
	ans := 0
	mps := &MatrixPreSum{
		preSum: make([][]int, row, row),
	}
	// 预计算二维数组的前缀和
	for x := 0; x < row; x++ {
		mps.preSum[x] = make([]int, col, col)
		for y := 0; y < col; y++ {
			mps.preSum[x][y] = mps.getSum(x-1, y) + mps.getSum(x, y-1) - mps.getSum(x-1, y-1) + matrix[x][y]
		}
	}
	// 固定上端点
	for top := 0; top < row; top++ {
		//固定下端点
		for bottom := top; bottom < row; bottom++ {
			// 消除不同上下端点的影响
			cntMap := make(map[int]int)
			cntMap[0] = 1
			// 横向遍历所有列
			// 横向的上下端点固定后, 子矩阵的前缀和就转化为了一维数组中连续子数组和为K的问题
			for cloumn := 0; cloumn < col; cloumn++ {
				//计算某列的前缀和
				cur := mps.getSum(bottom, cloumn) - mps.getSum(top-1, cloumn)
				ans += cntMap[cur-target]
				cntMap[cur]++
			}
		}
	}
	return ans
}
