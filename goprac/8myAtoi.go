/*
idea is easy but alot of edgecases

check for empty string
check for overflow

loop through string
keep looking until we find the firstNonWhiteSpace
when we do find the first nonwhitespace check if it is a digit or if it is a sign
and remember the first position of the firstNonWhiteSpace
if it is a sign then go to the next character since its not a digit

when we have found the firstnonwhitespace
check if it is a valid digit 
if not then return what the number is right now with the correct sign
otherwise keep adding to the number 
check if overflow everytime, and if overflow return the minInt or minMax

*/

func myAtoi(str string) int {
	if IsEmptyStr(str) {
		return 0
	}
	firstNonWhiteSpace := -1
	num := 0
	overflow := CheckOverFlow()
	for i, _ := range str {
		if str[i] != ' ' && firstNonWhiteSpace == -1 {
			if IsValidDigit(str[i]) || IsSign(str[i]) {
				firstNonWhiteSpace = i
			} else {
				return 0
			}
			if IsSign(str[i]) {
				continue
			}
		}
		if firstNonWhiteSpace != -1 {
			if !IsValidDigit(str[i]) {
				if str[firstNonWhiteSpace] == '-' {
					num *= -1
				}
				return num
			}
			num = AppendDigitToNum(num, ByteToInt(str[i]))

			if overflow(num) {
				if str[firstNonWhiteSpace] == '-' {
					return MinInt
				} else {
					return MaxInt
				}
			}
		}
	}
	if firstNonWhiteSpace == -1 {
		return num
	}
	if str[firstNonWhiteSpace] == '-' {
		num *= -1
	}
	return num
}

const MaxInt = int(^uint32(0) >> 1)
const MinInt = -MaxInt - 1

func CheckOverFlow() func(int) bool {
	numOfDigits := 0
	return func(num int) bool {
		if num > 0 {
			numOfDigits++
		}
		if num < MinInt || num > MaxInt || numOfDigits > 10 {
			return true
		}
		return false
	}
}

func IsValidDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func IsSign(c byte) bool {
	if c == '-' || c == '+' {
		return true
	}
	return false
}

func ByteToInt(c byte) int {
	return int(c - '0')
}

func AppendDigitToNum(num int, digit int) int {
	if num == 0 {
		return digit
	}
	return num*10 + digit
}

func IsEmptyStr(str string) bool {
	if str == "" {
		return true
	}
	return false
}
