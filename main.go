package main

func main() {

	// data provider

	var test []string
	test = make([]string, 1)
	if nil == test {
		println("test is nil")
	}
	if len(test) > 0 {
		println("111111111111")
	}
	for idx := range test {
		println("2222222222")
		println(test[idx])
	}
	println("333333333333")
}
