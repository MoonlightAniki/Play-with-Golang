package main

import "fmt"

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
func lengthOfLongestSubstring(s string) int {
	res := 0
	freq := make(map[rune]int)
	start := 0
	chars := []rune(s)
	for i := 0; i < len(chars); {
		if freq[chars[i]] == 0 {
			freq[chars[i]]++
			res = max(res, i-start+1)
			i++
		} else {
			freq[chars[start]]--
			start++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabccc"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pppwwwkkkpwk"))
	fmt.Println(lengthOfLongestSubstring("一二三三二一"))
}

// Runtime: 20 ms, faster than 33.99% of Go online submissions for Longest Substring Without Repeating Characters.
