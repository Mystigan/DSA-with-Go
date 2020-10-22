## [852. Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)
 Let's call an array arr a mountain if the following properties hold:

 - arr.length >= 3
 - There exists some i with 0 < i < arr.length - 1 such that:
   - arr[0] < arr[1] < ... arr[i-1] < arr[i]
   - arr[i] > arr[i+1] > ... > arr[arr.length - 1]
 Given an integer array arr that is guaranteed to be a mountain, return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

 ## Constraints
 - 3 <= arr.length <= 104
 - 0 <= arr[i] <= 106
 - arr is guaranteed to be a mountain array.

 ## Sample input
 ### Example 1
 ```c
 Input: arr = [0,1,0]
 Output: 1
 ```
### Example 2
 ```c
 Input: arr = [24,69,100,99,79,78,67,36,26,19]
 Output: 2
```