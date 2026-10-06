// Package luhn validate the transactions
package luhn

// Valid validates the transaction
func Valid(id string) bool {
	sum := 0
	digits := 0
	mustDouble := false

	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue
		}

		if id[i] < '0' || id[i] > '9' {
			return false
		}

		number := int(id[i] - '0')

		if mustDouble {
			number *= 2

			if number > 9 {
				number -= 9
			}
		}

		sum += number
		digits++
		mustDouble = !mustDouble
	}

	return digits > 1 && sum%10 == 0
}
