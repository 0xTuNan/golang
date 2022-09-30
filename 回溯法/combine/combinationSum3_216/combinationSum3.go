package combinationSum3_216

var res [][]int

func CombinationSum3(k int, n int) [][]int {
	var sum int
	res = [][]int{}
	for i := 1; i <= k; i++ {
		sum = i + sum
	}
	if sum > n {
		return [][]int{}
	}
	backtrack(k, n, 1, []int{})
	return res
}

func backtrack(k, n, indexstart int, temp []int) {
	var sum int
	if len(temp) == k {
		for _, i := range temp {
			sum = sum + i
		}
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
	}

	for i := indexstart; i < 10; i++ {
		temp = append(temp, i)
		backtrack(k, n, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}

//同77 先进行小减枝 然后进行回溯法
//1.记录temp和res 当满足条件的时候就可以进行判断结果满足不满足我们需求的和为n，为n就加入最终集合res，然后return
//2.进行for循环，组合3个数字出来，取一个放一个到temp里面，然后执行递归函数，最后记得弹出temp的最后一个值
