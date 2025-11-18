/*
控制流程
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素。
https://leetcode.cn/problems/single-number/
*/
func singleNumber(nums []int) int {
	m := make(map[int]int, 5)
	fmt.Printf("初始map: %v, 长度: %d\n", m, len(m))

	for _, num := range nums {
		m[num] += 1
	}
	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}

/*
回文数
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
https://leetcode.cn/problems/palindrome-number/
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	fmt.Println(strconv.Itoa(x))
	xStr := strconv.Itoa(x)
	xStrLen := len(xStr)
	xStrHalfLen := xStrLen / 2
	for i := 0; i < xStrHalfLen; i++ {
		if xStr[i] != xStr[xStrLen-1-i] {
			return false
		}
	}
	return true
}

/*
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
链接：https://leetcode-cn.com/problems/valid-parentheses/
*/

func isValid(s string) bool {

	for {
		oldStr := s

		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")

		if s == oldStr {
			break
		}
	}
	return s == ""
}

func isValid1(s string) bool {

	stack := []uint8{}
	pairs := [3][2]uint8{
		{'(', ')'},
		{'[', ']'},
		{'{', '}'},
	}

	for index := range s {
		isFind := false
		for idx, pair := range pairs {
			fmt.Println("idx: ", idx, " s[index]: ", s[index], "pair[1]: ", pair[1])
			if s[index] == pair[1] {
				fmt.Println("len(stack)-1: ", len(stack)-1, "stack[len(stack)-1]: ",
					stack[len(stack)-1], "pair[0]: ", pair[0])
				if stack[len(stack)-1] == pair[0] {
					fmt.Println("Before stack: ", stack)
					stack = stack[:len(stack)-1]
					fmt.Println("After stack: ", stack)
					isFind = true
				}
			}
		}
		if !isFind {
			stack = append(stack, s[index])
			fmt.Println("stack: ", stack)
		}
	}
	return len(stack) == 0
}

/*
最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {

	prefixStr := strs[0]
	for i := 1; i < len(strs); i++ {
		strLen := len(strs[i])
		for idx := range prefixStr {
			if idx < strLen && prefixStr[idx] != strs[i][idx] { // 会报错，为什么这么做不行
				prefixStr = prefixStr[:idx+1]
			}
		}
		if prefixStr == "" {
			return ""
		}
	}
	return prefixStr
}

func longestCommonPrefix(strs []string) string {

	prefixStr := strs[0]
	for i := 1; i < len(strs); i++ {
		strLen := len(strs[i])
		for j := 0; j < len(prefixStr); j++ {
			if j < strLen && prefixStr[j] != strs[i][j] {
				prefixStr = prefixStr[:j]
			}
		}
		if prefixStr == "" {
			return ""
		}
	}
	return prefixStr
}

/*
基本值类型
加一
难度：简单
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
链接：https://leetcode-cn.com/problems/plus-one/
*/

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	isNeed := false
	for i := digitsLen - 1; i >= 0; i-- {
		if i == (digitsLen - 1) {
			isNeed = true
		}
		if isNeed {
			if digits[i] == 9 {
				digits[i] = 0
				isNeed = true
			} else {
				digits[i] += 1
				isNeed = false
			}
		}
	}
	if isNeed {
		digits = append([]int{1}, digits...)
	}
	return digits
}

/*
引用类型：切片
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {
	m := make(map[int]int, 100)
	i := 0
	for _, num := range nums {
		_, isExist := m[num]
		if !isExist {
			nums[i] = num
			m[num] = 1
			i++
		}
	}
	return i
}

/*
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
https://leetcode.cn/problems/merge-intervals/
*/
func merge(intervals [][]int) [][]int {
	lenI := len(intervals)
	if lenI <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	for i := 0; i < lenI; i++ {
		lenR := len(result)
		if lenR == 0 || intervals[i][0] > result[lenR-1][1] {
			result = append(result, intervals[i])
		} else {
			result[lenR-1][1] = intervals[i][1]
		}
	}
	return result
}

/*
基础
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
链接：https://leetcode-cn.com/problems/two-sum/
*/

func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for idx, num := range nums {
		if num < target {
			value, exist := m1[num]
			fmt.Println(value, exist)
			if num*2 == target {
				if exist {
					return []int{value, idx}
				}
			}

			m1[num] = idx
			m2[target-num] = idx
		}
	}
	fmt.Println(m1)
	fmt.Println(m2)
	for key, value := range m1 {
		fmt.Println(key, value)
		val, exist := m2[key]
		if exist && value != val {
			fmt.Println(key, value)
			return []int{value, val}
		}
	}
	return []int{0, 0}
}


