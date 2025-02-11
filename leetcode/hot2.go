package leetcode

func twoSum2(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, v := range nums {
		if j, ok := hashmap[target-v]; ok {
			return []int{i, j}
		}

		hashmap[v] = i
	}
	return nil
}
