package week05

// 1011. 在 D 天内送达包裹的能力

// 解题: 题目要求计算出在D天内将所有货物运送到对岸,货船到最低运载量
// 1. 首先二分查找一个数为货物的运载量,计算以该数为货船的运载量,能否在D天将所有货物运送到对岸

func shipWithinDays(weights []int, days int) int {

	var left int = 0        //定义二分左端点
	var right int = 1 << 32 // 定义二分右端点

	for left < right {
		var mid int = (left + right) / 2
		// 判断mid是否合法
		if TransAmount(weights, days, mid) {
			// 合法的话右端点设置 mid
			right = mid
		} else {
			// 不合法的左端点 mid+1
			left = mid + 1
		}
	}
	return right
}

// TransAmount 判定函数
func TransAmount(weights []int, days, weight int) bool {

	var wg int = 0
	//遍历所有包裹
	for _, val := range weights {
		// 若一个包裹的重量大于货船的运载量,表示weightb不合法
		if val > weight {
			return false
		}
		if wg+val <= weight {
			wg += val
		} else {
			days--
			wg = val
		}
	}
	// 只有剩余天数大于0,weight才合法
	return days > 0
}
