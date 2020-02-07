package main

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func equal2(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if xv != y[k] {
			//y["A"] = xv(0)
			return false
		}
	}
	return true
}

func main() {
	//已经存在的0，和不存在而返回零值的0区分
	println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
	println(equal2(map[string]int{"A": 0}, map[string]int{"B": 42}))
	m := map[string]int{"A": 0}
	println(m["A"], m["B"])
}
