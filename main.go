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



//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
//Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
//The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.


func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}


//Input: nums = [3,2,2,3], val = 3
//Output: 2, nums = [2,2,_,_]
//Explanation: Your function should return k = 2, with the first two elements of nums being 2.
//It does not matter what you leave beyond the returned k (hence they are underscores).

// 第一輪迴圈 nums[0] == val == 3


func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
			i--
		}
	}
	return n
}

//nums = [3,2,3]
//        nums[i] != nums[i+1]
//        nums[i] = nums [i+2]
//output = 3

package main

import "fmt"

func majorityElement(nums []int) int {
	// 候选众数和计数器
	candidate, count := 0, 0

	// 遍历数组
	for _, num := range nums {
		if count == 0 {
			// 如果计数器为 0，选择当前数字为候选众数
			candidate = num
			count = 1
		} else if num == candidate {
			// 如果当前数字等于候选众数，计数器增加
			count++
		} else {
			// 如果当前数字不等于候选众数，计数器减少
			count--
		}
	}

	// 返回候选众数
	return candidate
}

//改良
// nums = [3,2,3]
// sort [2,3,3]
//output = 3

func majorityElement(nums []int) int {
	// 首先对数组进行排序
	sort.Ints(nums)

	// 因为众数一定会出现在中间位置，直接返回即可
	return nums[len(nums)/2]
}

//Input: nums = [1,1,2]
//Output: 2, nums = [1,2,_]


func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1

}


//Input: nums = [1,1,1,2,2,3]

//Output: 5, nums = [1,1,2,2,3,_]
//#### **範例 1**
//輸入：`nums = [1, 1, 1, 2, 2, 3]`
//1. 起始：`j = 2`，因為前兩個元素始終有效（`nums[0]`、`nums[1]`）。
//2. 從 `i = 2` 開始遍歷：
//    - **`nums[i] = 1`**，與**`nums[j-2] = 1`** 相同（多餘），跳過。
//    - **`nums[i] = 2`**，與**`nums[j-2] = 1`** 不同，保留到 `nums[j]`，即 `nums[2] = 2`，`j++`。
//    - **`nums[i] = 2`**，與**`nums[j-2] = 2`** 不同，保留到 `nums[j]`，`nums[3] = 2`，`j++`。
//    - **`nums[i] = 3`**，與**`nums[j-2] = 2`** 不同，保留到 `nums[j]`，`nums[4] = 3`，`j++`。

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	j := 2
	for i := 2; i < len(nums); i++ {
		// 檢查 nums[i] 是否超過允許
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

//Input: s = "III"
//Output: 3
//Explanation: III = 3.

func romanToInt(s string) int {
	romanToIntMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	n := len(s)
	result := 0
	for i := 0; i < n; i++ {
		value := romanToIntMap[string(s[i])]

		if i < n-1 && value < romanToIntMap[string(s[i+1])] {
			result -= value
		} else {
			result += value
		}
	}
	return result


}

//Input: s = "Hello World"
//Output: 5
//Explanation: The last word is "World" with length 5.

func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	if n == 0 {
		return 0
	}
	for i >= 0 && s[i] != ' ' {
		i --
	}

	length := 0
	for i >= 0 && s[i] == ' ' {
		i--
	}
	for i >= 0 {
		length++
	}
	return length


}


//better version

func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1

	// 忽略尾部空格
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// 計算最後一個單詞的長度
	length := 0
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}










