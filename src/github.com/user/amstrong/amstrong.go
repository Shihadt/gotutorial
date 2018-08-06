package main

func isAmstrong(n int) bool {
	sum := 0
	temp := n
	for temp != 0 {
		rem := temp % 10
		sum += rem * rem * rem
		temp /= 10
	}
	if sum == n {
		return true
	}
	return false
}
