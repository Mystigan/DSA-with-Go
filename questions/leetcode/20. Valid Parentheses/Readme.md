## [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
 Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

 An input string is valid if:

 - Open brackets must be closed by the same type of brackets.
 - Open brackets must be closed in the correct order.

## Constraints
 - 1 <= s.length <= 104
 - s consists of parentheses only '()[]{}'.

 ## Sample input
 ### Example 1
 ```c
 Input: s = "()"
 Output: true
 ```
 ### Example 2
 ```c
 Input: s = "{[]}"
 Output: true
 ```
 ### Example 3
 ```c
 Input: s = "([)]"
 Output: false
 ```