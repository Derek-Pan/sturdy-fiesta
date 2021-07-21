package week05

// 875. 爱吃香蕉的珂珂

// 解题: 题目要求koko以最慢的速度吃香蕉,能够在H小时内将所有香蕉吃完
// 1. 首先二分查找一个数为吃香蕉的速度,判定能否在规定时间内吃完

func minEatingSpeed(piles []int, h int) int {
	var left int = 1
	var right int = 1 << 32

	for left < right {
		var mid int = (left + right) / 2
		if CanEatFinsh(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

// CanEatFinsh 判定以speed的速度吃香蕉,是否能够完成
func CanEatFinsh(piles []int, h, speed int) bool {
	for _, val := range piles {
		// 计算吃该堆香蕉需要多少时间, 向下取整
		var time int = val / speed
		// 若是还有剩余不足speed根香蕉,则以一小时为准
		if val%speed != 0 {
			time = time + 1
		}
		// 剩余时间小于0, 表示以speed速度不可完成
		if h-time < 0 {
			return false
		}
		h -= time
	}
	// 剩余时间大于等于0, 表示以speed速度可完成
	return h >= 0
}
