// https://go.dev/play/p/SZq-t6YmEo_-
// svijaysuriya

package main

func genSeq(start int) func() int {
	i := start
	return func() int {
		i++
		return i - 1
	}
}
func main() {
	// concurrent vs parallelism - structing the program to be indpendent components( dealing with lot of things at once) vs doing lot of things at once

	// closures
	alphaSeq := genSeq(1)
	betaSeq := genSeq(1)

	print("address of main() = ", main, "\n")
	print("address of genSeq() = ", genSeq, "\n")
	print("address of alphaSeq() = ", alphaSeq, "\n")
	print("address of betaSeq() = ", betaSeq, "\n")

	print("alphaSeq = ", alphaSeq(), "\n")
	print("betaSeq = ", betaSeq(), "\n")

	channel := make(chan int, 2)

	go func() {
		a := alphaSeq()
		b := betaSeq()

		print("alphaSeq inside go routine = ", a, "\n")
		print("betaSeq inside go routine = ", b, "\n")
		channel <- a
		channel <- b

		close(channel)
	}()
	print("alphaSeq = ", alphaSeq(), "\n")
	print("betaSeq = ", betaSeq(), "\n")

	for {
		num, open := <-channel
		if !open {
			break
		}
		print("channel value inside go routine = ", num, "\n")
	}
}
