package common

func Max[V int | float32](vals ...V) V {
	var max V
	if len(vals) == 1 {
		return vals[0]
	}
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min[V int | float32](vals ...V) V {
	var min V
	if len(vals) == 1 {
		return vals[0]
	}
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
