package day_3

func TestFunc2() {
	done := make(chan bool)
	c := make(chan string)

	go func() {
		defer close(done)
		s := <-c
		println(s)
	}()

	c<-"hi"
	<-done
}
