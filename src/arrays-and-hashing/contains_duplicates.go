/*
*
leetcode problem 217:

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true

Example 2:

Input: nums = [1,2,3,4]
Output: false

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:

	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9

*
*/

package arraysandhashing

// NOTE: my solution:
func containsDuplicate(nums []int) bool {
	frequencyOfNums := make(map[int]int)
	for _, item := range nums {
		frequencyOfNums[item]++
		if frequencyOfNums[item] > 1 {
			return true
		}
	}
	return false
}

/*
*
NOTE: another solution that is similar to the one above, but is slightly better. this uses a hashset where
you map what is seen already and if it already exists in the set, then you just return true otherwise it's false.
*
*/
func containsDuplicate2(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		_, exists := seen[num]
		if exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

/*
*

	NOTE: another example from another leetcode submitter thats just way smarter imo and has less boiler plate,

adding it here because i think it's really cool the way they saw the problem
*
*/
func containsDuplicate3(nums []int) bool {
	seen := make(map[int]bool)

	for _, v := range nums {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
