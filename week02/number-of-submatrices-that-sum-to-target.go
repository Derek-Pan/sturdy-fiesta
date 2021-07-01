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
    ans int
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
    xLength := len(matrix)
    yLength := len(matrix[0])
    ps := &NumberMatrix{
        preSum: make([][]int, xLength, xLength),
        ans: 0,
    }
    // 计算二维数组的前缀和
    for x := 0; x < xLength; x++ {
        ps.preSum[x] = make([]int, yLength, yLength)
        for y:=0;y <yLength; y++ {
            ps.preSum[x][y] = ps.getSum(x-1, y) + ps.getSum(x, y-1) - ps.getSum(x-1, y-1) + matrix[x][y]
        }
    }
    // 以x1, y1点进行递归
    ps.subMatrixSum(0, 0, target)
    return ps.ans
}

func (this *NumberMatrix) getSum(x, y int) int {
    if x >= 0 && y >= 0 {
        return this.preSum[x][y]
    }
    return 0
}

func (this *NumberMatrix) subMatrixSum(x1, y1, target int) {
    fmt.Println(x1, y1)

    // x1, y1超过数组长度后结束
    if x1 >= len(this.preSum) && y1 >= len(this.preSum[0]) {
        return
    }
    // 遍历x2, y2所有的可能性
    for i := x1; i < len(this.preSum); i++ {
        for j := y1; j < len(this.preSum[0]); j++ {
            subSum := this.getSum(i, j) - this.getSum(i, y1-1) - this.getSum(x1-1, j) + this.getSum(x1-1,y1-1)
            if subSum == target {
                this.ans++
            }
        }
    }

    // 计算下一个符合条件的x1,y1点
    if y1 >= len(this.preSum[0]) {
        this.subMatrixSum(x1+1, 0, target)
    } else {
        this.subMatrixSum(x1, y1+1, target)
    }
}
