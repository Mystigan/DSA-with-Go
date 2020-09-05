## [1351. Count Negative Numbers in a Sorted Matrix](https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/)
 Given a m * n matrix grid which is sorted in non-increasing order both row-wise and column-wise. 

 Return the number of negative numbers in grid.

 ## Constraints
 ```c
 m == grid.length
 n == grid[i].length
 1 <= m, n <= 100
 -100 <= grid[i][j] <= 100
 ```

 ## Sample input
 ### Example 1
 ```c
 Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
 Output: 8
 Explanation: There are 8 negatives number in the matrix.
 ```

 ### Example 2
 ```c
 Input: grid = [[3,2],[1,0]]
 Output: 0
 ```