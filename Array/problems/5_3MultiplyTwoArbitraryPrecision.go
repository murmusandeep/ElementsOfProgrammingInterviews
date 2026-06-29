package problems

// 43. Multiply Strings
// https://leetcode.com/problems/multiply-strings/

func Multiply(num1 string, num2 string) string {

	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			mul := int(num1[i]-'0') * int(num2[j]-'0')

			p1 := i + j
			p2 := i + j + 1

			sum := mul + result[p2]

			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}

	var ans []byte
	leadingZero := true

	for i := 0; i < len(result); i++ {
		digit := result[i]

		if digit == 0 && leadingZero {
			continue
		}

		leadingZero = false
		ans = append(ans, byte(digit)+'0')
	}

	if len(ans) == 0 {
		return "0"
	}

	return string(ans)
}
