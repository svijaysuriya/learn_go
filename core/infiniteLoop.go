package main

func main() {
	var i int
	go func() {
		for i = 0; ; i += 2 {
			print(i, "\t")

		}
	}()
	for i = 1; ; i += 2 {
		print(i, "\t")
	}
}
