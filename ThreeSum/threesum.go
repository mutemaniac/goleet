package ThreeSum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] > 0 {
			break
		}
		lo := i + 1
		hi := numsLen - 1
		sum := 0 - nums[i]
		for {
			if lo >= hi {
				break
			}
			if nums[lo]+nums[hi] == sum {
				result = append(result, []int{nums[i], nums[lo], nums[hi]})
				for {
					if nums[lo] == nums[lo+1] && (lo+1) < hi {
						lo++
					} else {
						break
					}
				}
				for {
					if nums[hi] == nums[hi-1] && (hi-1) > lo {
						hi--
					} else {
						break
					}
				}
				lo++
				hi--
			} else if nums[lo]+nums[hi] > sum {
				hi--
			} else {
				lo++
			}
		}
		for (i + 1) < numsLen {
			if nums[i] == nums[i+1] {
				i++
			} else {
				break
			}
		}
	}
	return result
}

func getFirstPositive(nums []int) int {
	for i, val := range nums {
		if val >= 0 {
			return i
		}
	}
	return -1
}

func threeSum1(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	firstPositive := getFirstPositive(nums)

	numsLen := len(nums)
	for i := 0; i <= firstPositive; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// for {
		// 	if i < (numsLen-2) && nums[i]+nums[i+1]+nums[numsLen-1] > 0 {
		// 		numsLen--
		// 	} else {
		// 		break
		// 	}
		// }

		for j := i + 1; j < numsLen; j++ {
			twoSum := nums[i] + nums[j]
			if twoSum > 0 {
				break
			}
			if twoSum+nums[numsLen-1] < 0 {
				continue
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			firstK := firstPositive
			if firstK < j+1 {
				firstK = j + 1
			}
			for k := firstK; k < numsLen; k++ {
				if (nums[i] + nums[j] + nums[k]) == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					break
				}
				if (nums[i] + nums[j] + nums[k]) > 0 {
					break
				}
				if k > firstK && nums[k] == nums[k-1] {
					continue
				}
			}
		}
	}
	return result
}
