/*
557. Reverse Words in a String III

Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "God Ding"
Output: "doG gniD"

https://leetcode.com/problems/reverse-words-in-a-string-iii
*/
package main

func reverseWords(s string) string {
	ans := make([]byte, len(s))
	tmp := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			for k, v := range tmp {
				ans[i+k+1] = v
			}
			ans[i] = ' '
			tmp = nil
			continue
		}

		tmp = append(tmp, s[i])
		if i == 0 {
			for k, v := range tmp {
				ans[k] = v
			}
		}
	}
	return string(ans)
}
