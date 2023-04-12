package main

func main() {
	forever := make(chan bool)

	//Resolver o deadloack, o canal precisa ser preechido em outra thread.
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	//Deadloack!!
	//Fica eternamente esperado o canal
	<-forever
}
