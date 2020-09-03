## [1480. Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/)
 Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

 Return the running sum of nums.

 ## Constraints
 ```c
 1 <= nums.length <= 1000
 -10^6 <= nums[i] <= 10^6
 ```

 ## Sample input
 ### Example 1
 ```c
 Input: nums = [1,2,3,4]
 Output: [1,3,6,10]
 Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
 ```

 ### Example 2
 ```c
 Input: nums = [1,1,1,1,1]
 Output: [1,2,3,4,5]
 Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
 ```