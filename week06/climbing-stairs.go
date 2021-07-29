package week06


/*
题目: 爬楼梯

定义状态:
	i: 当前在第几阶楼梯
	目标: 有多少中方案可以爬到楼顶

状态转移方程
			当 i == 1时: F(i) = 1
	F(i) =  当 i == 2时: F(i) = 2
			当 i > 2时: F(i) =F(i-1) + F(i-2)
*/


func climbStairs(n int) int {
	// dp数组
	var f []int = make([]int, n+1, n+1)
	f[0] = 0

	// 遍历状态
	for i := 1; i <= n; i++ {
		// 状态转移方程的实现
		if i == 1 {
			f[i] = 1
		} else if i == 2 {
			f[i] = 2
		} else {
			f[i] = f[i-1] + f[i-2]
		}
	}
	return f[n]
}
