## [844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)
 Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

 Note that after backspacing an empty text, the text will continue empty.

## Constraints
 - 1 <= S.length <= 200
 - 1 <= T.length <= 200
 - S and T only contain lowercase letters and '#' characters.

 ## Sample input
 ### Example 1
 ```c
 Input: S = "ab#c", T = "ad#c"
 Output: true
 Explanation: Both S and T become "ac".
 ```

 ### Example 2
 ```c
 Input: S = "ab##", T = "c#d#"
 Output: true
 Explanation: Both S and T become "".
 ```