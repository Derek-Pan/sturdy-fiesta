package main

// 和为k的子数组

// 主体思路: 由于题目要求的是连续子数组的和为K, 于是可以使用前缀和, 预先计算数组的前缀和
// 细节: 因为该题目的元素不是都 >=0,无法保证前缀和数组的单调性, 所以在进行前缀和数组每个数出现的频次时,不能够预先全部统计出来,只能在最后遍历前缀和数组的时候统计j<i的项,这样能保证不会有重复

// 总结: 与优美子数组相比,该题的最大的区别就是,前缀和数组不在具有单调性, 无法保证 s[j] = s[i] - k (i > j) 统计出来的j的个数中每一个的j<i,所以在解题的时候出现了,某些输入的结果比正确的结果要大;
// 优美子数组具有单调性,所以保证了上面等式可以成立. 该题的话就只能在计算出前缀和数组之后,一边遍历数组,一边进行出现数字频次的统计

func subarraySum(nums []int, k int) int {
	length := len(nums)

	//统计前缀和数组中不同的数出现的次数
	cntMap := map[int]int{0: 1}

	answer := 0
	sum := make([]int, length+1, length+1)

	// 计算原始数组的前缀和
	for i := 0; i < length; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	for i := 1; i <= length; i++ {
		// 统计对于i, 有多少个j可以满足s[j] = sum[i] - k
		answer += cntMap[sum[i]-k]
		//将当前索引的前缀和加入到统计项当中
		cntMap[sum[i]]++
	}
	return answer
}

func main() {

}
