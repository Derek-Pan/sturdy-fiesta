package week06

import "math"

/*
题目: 跳跃游戏

定义状态:
	i: 当前所在位置

目标: 能否到达最后一个下标 == 每个下标可以到达到最远的位置

状态转移方程:
			当 F[i-1] >= i; max(F[i-1], i+nums[i])
	F[i] =  当F[i-1] < i; F[i-1]
 */
 
func canJump(nums []int) bool {
	var length int = len(nums)
	var f []int = make([]int, length, length)
	f[0] = nums[0]
	for i := 1; i < length; i++ {
		if f[i-1] >= i {
			f[i] = int(math.Max(float64(f[i-1]), float64(i+nums[i])))
		}
	}
	return f[length-1] > 0 || length == 1
}
