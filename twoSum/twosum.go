package main

import "fmt"

// 暴力破解
// 複雜度評估
// 設有 n 個數時間複雜度為 O(n²)
// 此方法所需記憶體與 n 無關， 空間複雜度為 O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		fmt.Printf(" i %d \n", i)
		for j := i + 1; j < len(nums); j++ {
			fmt.Printf("j %d \n", j)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// Hash
// 複雜度評估
// 該演算法至多只需要查找 n 個數，所需要的時間複雜度為O(n)。
// 而此方法需要另外另外儲存 n 個數的對應 map ，空間複雜度為 O(n)。
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		j, exist := m[target-v]
		if exist {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{-1, -1}
}

//解法 1 與 2 各有優劣，我們一般會採用時間複雜度較低的解法

func main() {
	A1 := twoSum([]int{1, 5, 2}, 6)
	fmt.Printf("A1 %v", A1)

	A2 := twoSum2([]int{1, 5, 2}, 6)
	fmt.Printf("A2 %v", A2)

}
