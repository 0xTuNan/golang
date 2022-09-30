package combinationSum_39

import "sort"

func CombinationSum(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

//这道题目的主要在于是否使用这个数字，如果不使用就加一跳过数字，使用则需要进行相应判断
//这个算法是看牛客网上的别人答案

var res [][]int

func CombinationSumMe(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	if candidates[0] > target {
		return [][]int{}
	}
	backtrack(candidates, target, []int{}, 0)
	return res
}

func backtrack(candidates []int, target int, temp []int, indexStart int) {
	var sum int
	for _, i := range temp {
		sum = sum + i
	}
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
	} else if sum > target {
		return
	}

	for i := indexStart; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		backtrack(candidates, target, temp, i)
		temp = temp[:len(temp)-1]
	}
}

//很简单，基本没啥区别，主要是数可复用和返回条件这里有差异
//对于数可复用，那么在处理最下面的temp的时候，他的开始(indexStart)应该就不是平常的i+1了,而是i,i+1就跳到下一个数了。
//对于返回条件这里，平常都是temp数组的长度等于一个固定值就返回,这里并不是，而是当最后的数组和大于等于target的时候才会返回，同时等于为满足条件，记录到res中，大于就忽略这个temp。
