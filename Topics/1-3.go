package Topics

import "fmt"

type cprefixSorting struct {
	// 烙饼个数
	cakeCnt int
	// 最多交换次数。最大值为 cakeCnt * 2
	maxSwap int
	// 烙饼信息数量
	cakeArray []int
	// 交换结果数组
	swapArray []int
	// 当前翻转烙饼信息数组
	reverseCaKeArray []int
	// 当前翻转烙饼交换结果数组
	reverseCakeArraySwap []int
	// 当前搜索次数信息
	search int
}

// init 初始化信息
func (cp *cprefixSorting) init(cakeArr []int, cakeCount int) {
	cp.cakeCnt = cakeCount
	// 初始化烙饼数组
	cp.cakeArray = make([]int, cakeCount)
	copy(cp.cakeArray, cakeArr)
	// 设置最多交换次数信息
	cp.maxSwap = cp.UpperBound(cakeCount)
	// 初始交换结果数组
	cp.swapArray = make([]int, cp.maxSwap+1)
	// 初始化中间交换结果信息
	cp.reverseCaKeArray = make([]int, cakeCount)
	copy(cp.reverseCaKeArray, cakeArr)
	cp.reverseCakeArraySwap = make([]int, cp.maxSwap+1)
}

// UpperBound 寻找当前翻转的上界
func (cp *cprefixSorting) UpperBound(cakeCnt int) int {
	return cakeCnt * 2
}

// LowerBound 找到当前翻转的下界
func (cp *cprefixSorting) LowerBound(cakeArray []int, cakeCnt int) (ret int) {
	t := 0
	// 根据当前数组的排序信息情况来判断最少需要交换多少次
	for i := 1; i < cakeCnt; i++ {
		//判断位置相邻的两个烙饼，是否为尺寸排序上相邻的
		t = cakeArray[i] - cakeArray[i-1]
		if t != 1 && t != -1 {
			ret += 1
		}
	}
	return
}

func (cp *cprefixSorting) Search(step int) {
	i, estimate := 0, 0
	cp.search += 1
	// 估算这次搜索所需要的最小交换数
	estimate = cp.LowerBound(cp.reverseCaKeArray, cp.cakeCnt)
	if step+estimate > cp.maxSwap {
		return
	}
	// 如果已经排好序，即翻转完成，输出结果
	if cp.IsSorted(cp.reverseCaKeArray, cp.cakeCnt) {
		if step < cp.maxSwap {
			cp.maxSwap = step
			for i = 0; i < cp.maxSwap; i++ {
				cp.swapArray[i] = cp.reverseCakeArraySwap[i]
			}
		}
		return
	}
	// 递归进行翻转
	for i = 1; i < cp.cakeCnt; i++ {
		cp.Revert(0, i)
		cp.reverseCakeArraySwap[step] = i
		cp.Search(step + 1)
		cp.Revert(0, i)
	}
}

func (cp *cprefixSorting) Revert(begin, end int) {
	i, j, t := begin, end, 0
	for i < j {
		t = cp.reverseCaKeArray[i]
		cp.reverseCaKeArray[i] = cp.reverseCaKeArray[j]
		cp.reverseCaKeArray[j] = t
		i += 1
		j -= 1
	}
}

func (cp *cprefixSorting) IsSorted(cakeArray []int, cakeCnt int) bool {
	for i := 1; i < cakeCnt; i++ {
		if cakeArray[i-1] > cakeArray[i] {
			return false
		}
	}
	return true
}

func (cp *cprefixSorting) Run(cakeArray []int, cakeCnt int) {
	cp.init(cakeArray, cakeCnt)
	cp.search = 0
	cp.Search(0)
}

func (cp *cprefixSorting) String() string {
	output := ""
	for i := 0; i < cp.maxSwap; i++ {
		output += fmt.Sprintf("%d ", cp.swapArray[i])
	}
	output += fmt.Sprintf("\n | Search Times| : %d\n", cp.search)
	output += fmt.Sprintf("Total Swap Times=%d\n", cp.maxSwap)
	return output
}

func Solution() {
	cp := cprefixSorting{}
	cp.Run([]int{3, 2, 1, 6, 5, 4, 9, 8, 7, 0}, 10)
	fmt.Printf(cp.String())
}
