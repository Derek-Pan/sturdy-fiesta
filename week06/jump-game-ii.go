package week06

import "math"

/*
题目: 跳跃游戏ii

思路: 要想到达f[i], j有机会从f[i-1], f[i-2]... f[0]的位置跳跃到, 遍历i-1的点中跳的次数最少的;

状态:
	i: 当前所在位置


状态转移方程:
	f[i] = min(f[i], f[j]+1) 0 <= j < i

目标: 到最后一个下标到最小跳跃数


*/

func jump(nums []int) int {
	var length int = len(nums)
	var f []int = make([]int, length, length)
	for i := 0; i < length; i++ {
		f[i] = 1 << 32
	}
	f[0] = 0

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = int(math.Min(float64(f[i]), float64(f[j]+1)))
			}
		}
	}
	return f[length-1]
}
