package ops


func Sum(x, y int) int{
	result := x + y
	if result < 0{
		return 0
	}

	return result
}
