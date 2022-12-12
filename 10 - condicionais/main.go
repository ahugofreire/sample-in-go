package main

func main() {
	dia := 5
	switch dia {
	case 0:
		println("Domingo")
	case 1:
		println("segunda")
	case 2, 3:
		println("terca ou quarta")
	default:
		println("Error")
	}
}
