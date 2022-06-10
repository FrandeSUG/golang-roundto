package main

func RoundTo(point int32, to int32) int32 {
	var posNeg int32 = 1

	if point < 0 {
		// Turn point to positive temporarily
		posNeg = -1
		point = point * posNeg
	}

	if point < to {
		return point * posNeg
	} else {
		if (point%to) >= (to/2) && posNeg == 1 || (point%to) < (to/2) && posNeg == -1 {
			return ((point / to) + (1 * posNeg)) * posNeg
		} else {
			return (point / to) * posNeg
		}
	}
}
