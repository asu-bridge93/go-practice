package functions

func NakedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
