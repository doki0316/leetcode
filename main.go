// input: nums = [1, 2, 3, 1]
// return true
package main

import (
	"fmt"
	"strings"
	"sort"
)


func containsDuplicate(nums []int) bool {
	InitNums := len(nums) //4

	for i := 0; i < InitNums; i++ {

		for j := i + 1; j < InitNums; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// [1, 1, 2, 3]

func containsDuplicate(nums []int) bool {
	InitNums := len(nums)
	sort.Ints(nums)
	for i := 1; i < InitNums; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

//nums [3, 2, 4] target = 6
//      0  1  2
//output [1, 2]

func twoSum(nums []int, target int) []int {
	//InitNums := len(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

//解題：
//for _, 迴圈
//split拆分字母變字符切片
//sort對切片字母重新排序
//join 變回字母

package main

import (
"fmt"
"sort"
"strings"
)

func groupAnagrams(strs []string) [][]string {
	// 用 map 保存分組
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// 將字符串轉換成字符切片並排序
		chars := strings.Split(str, "") // 將字符串轉為字符切片
		sort.Strings(chars)             // 對字符切片排序
		sortedStr := strings.Join(chars, "") // 排序後重新拼成字符串

		// 以排序後結果為鍵，將原始字符串添加到對應分組
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// 將 map 的值（分組）轉換成結果切片
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}





//有點難還是不太懂
func topKFrequent(nums []int, k int) []int {
	// Step 1: 計算每個數字的頻率
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Step 2: 使用一次迴圈找到前 k 個最大頻率的數字
	result := []int{} // 儲存結果
	for i := 0; i < k; i++ {
		// 找到當前最大頻率的數字
		maxFreq := 0
		maxNum := 0
		for num, freq := range frequencyMap {
			if freq > maxFreq {
				maxFreq = freq
				maxNum = num
			}
		}

		// 把這個數字加到結果裡
		result = append(result, maxNum)

		// 刪除已經處理過的數字，避免重複
		delete(frequencyMap, maxNum)
	}

	return result
}


// Input: nums = [1,2,3,4]
//     first num in output = 2x3x6=24
//     second num in output = 1x3x4=12
//     third num in output = 1x2x4
//      .........
//Output: [24,12,8,6]
//explain
//nums = [1, 2, 3, 4]
//        0  1  2  3
//answer = [24, 12, 8, 6]

//answer[0] = 2x3x4=24
//answer[1] = 1x3x4=12
//answer[2] = 1x2x4=8
//      [位置]
// 設位置 = i
//answer[i]




package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n) // 用來存結果，初始化為 0

	// 計算左側乘積，存入 answer
	leftProduct := 1
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		leftProduct *= nums[i]
	}

	// 計算右側乘積，累乘進 answer
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}// so fucking hard


