package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n, pick int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		val := guess(mid, pick)
		if val == -1 {
			right = mid - 1
		} else if val == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
