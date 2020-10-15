## [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)
 You are given an integer array nums sorted in ascending order, and an integer target.

 Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

 If target is found in the array return its index, otherwise, return -1.

## Constraints
 - 1 <= nums.length <= 5000
 - -10^4 <= nums[i] <= 10^4
 - All values of nums are unique.
 - nums is guranteed to be rotated at some pivot.
 - -10^4 <= target <= 10^4

 ## Sample input
 ### Example 1
 ```c
 Input: nums = [4,5,6,7,0,1,2], target = 0
 Output: 4
 ```
### Example 2
 ```c
 Input: nums = [4,5,6,7,0,1,2], target = 3
 Output: -1
```