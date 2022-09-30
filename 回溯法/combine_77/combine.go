package combine_77

var res [][]int

func Combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}
func backtrack(n, k, indexStart int, temp []int) {
	if len(temp) == k {
		tmp := make([]int, k)
		copy(tmp, temp)
		res = append(res, tmp)
		return
	}
	if len(temp)+n-indexStart+1 < k {
		return
	}

	for i := indexStart; i <= n; i++ {
		temp = append(temp, i)
		backtrack(n, k, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}

//定义全局变量res [][]int，存放最终结果，定义temp[]存放临时结果
//1.确定递归的终止条件，然后把temp的结果放到res里面去，最终返回
//2.处理节点，即吧数据加入到temp里面，做下一层的循环。
//3.弹出之前的数据，即把刚刚加入temp的数据清除了，数组就去除最后一位[:lem(temp)-1],字符串也是一样的去除最后一位[:len(str)-1]
