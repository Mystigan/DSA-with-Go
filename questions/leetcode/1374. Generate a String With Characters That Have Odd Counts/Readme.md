## [1374. Generate a String With Characters That Have Odd Counts](https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/)
 Given an integer n, return a string with n characters such that each character in such string occurs an odd number of times.

 The returned string must contain only lowercase English letters. If there are multiples valid strings, return any of them.

## Constraints
 - 1 <= n <= 500

 ## Sample input
 ### Example 1
 ```c
 Input: n = 4
 Output: "pppz"
 Explanation: "pppz" is a valid string since the character 'p' occurs three times and the character 'z' occurs once. Note that there are many other valid strings such as "ohhh" and "love".
 ```
 ### Example 2
 ```c
 Input: n = 2
 Output: "xy"
 Explanation: "xy" is a valid string since the characters 'x' and 'y' occur once. Note that there are many other valid strings such as "ag" and "ur".
 ```