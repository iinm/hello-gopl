package main

func f(val int) (result int) {
	defer func() {
		result = recover().(int)
	}()
	panic(val)
}
