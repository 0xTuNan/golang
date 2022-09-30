package combinationSum2_40

import (
	"sort"
)

var res [][]int

func CombinationSum2(candidates []int, target int) [][]int {
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
		//当前树层中有使用过相同元素
		if i > 0 && candidates[i] == candidates[i-1] && i > indexStart {
			continue
		}
		temp = append(temp, candidates[i])
		backtrack(candidates, target, temp, i+1)
		temp = temp[:len(temp)-1]
	}

}

//暂时我也不知道去重的i>indexStart是怎么把树层中的相同元素去除而不去除树枝中的，先记住吧
