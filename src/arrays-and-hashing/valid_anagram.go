package arraysandhashing

import "strings"

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false



Constraints:

    1 <= s.length, t.length <= 5 * 104
    s and t consist of lowercase English letters.



Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

// this was a good attempt and a decent idea BUT the logic fails in the case of s = "aacc" and t = "ccac"
// it's best to compare the frequency of each char from both s and t and then compare the maps to see if the values are the same.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	seen := make(map[string]bool)
	sToChar := strings.Split(s, "")
	tToChar := strings.Split(t, "")

	for _, sValues := range sToChar {
		seen[sValues] = true
	}
	for _, tValues := range tToChar {
		if !seen[tValues] {
			return false
		}
	}

	return true
}
