package main

func main() {
	s := []int{7, 1, 2, 3, 6, 1, 6}
	println(maxProfit(s))
}

func maxProfit(s []int) int {
	if len(s) <= 1 {
		return 0
	}

	minValue := s[0]
	var maxValue = 0

	for i := 0; i < len(s); i++ {
		if s[i] < minValue {
			minValue = s[i]
		}

		if s[i]-minValue > maxValue {
			maxValue = s[i] - minValue
		}
	}
	return maxValue
}
