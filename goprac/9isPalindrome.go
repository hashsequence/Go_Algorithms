/*
idea : idea is simple since you need to iterate from left and right to meet in the middle
the way to increment is
we get the rightstarter so if the number is 9433 then rightstarter is 1000
and left starter is 1
and to extract the digit the formula is num % ( starter * 10) / 10
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	j := GetRightStarter(x)
	i := 1
	for i <= j {
		if x%(i*10)/i != x%(j*10)/j {
			return false
		}
		i = i * 10
		j = j / 10

	}
	return true
}

func GetRightStarter(x int) int {
	i := 1
	for ; x/i != 0; i = i * 10 {
	}
	return i / 10
}
