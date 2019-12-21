package myAtoi

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

func myAtoi(str string) int {
	n := 0
	init := false
	m := 1
	for _, char := range []byte(str) {
		switch {
		case char == ' ':
			if init {
				return n
			}
		case char == '+':
			if !init {
				init = true
				m = 1
			} else {
				return n
			}
		case char == '-':
			if !init {
				init = true
				m = -1
			} else {
				return n
			}
		case char >= '0' && char <= '9':
			init = true
			n = n*10 + int(char-'0')*m
			if n > INT_MAX {
				return INT_MAX
			}
			if n < INT_MIN {
				return INT_MIN
			}
		default:
			if init {
				return n
			} else {
				return 0
			}
		}
	}
	return n
}
