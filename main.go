package main

import "fmt"

func main() {
	var mes, dia int
	fmt.Print("Introduce tu dia de nacimiento ")
	fmt.Scan(&dia)
	fmt.Print("Introduce tu mes de nacimiento ")
	fmt.Scan(&mes)
	if (dia >= 22 && mes == 12) || (dia <= 20 && mes == 1) {
		fmt.Print("capricornio")
	} else if (dia >= 21 && mes == 2) || (dia <= 19 && mes == 2) {
		fmt.Print("acuario")
	} else if (dia >= 20 && mes == 2) || (dia <= 20 && mes == 3) {
		fmt.Print("piscis")
	} else if (dia >= 21 && mes == 3) || (dia <= 20 && mes == 4) {
		fmt.Print("aries")
	} else if (dia >= 21 && mes == 4) || (dia <= 21 && mes == 5) {
		fmt.Print("tauro")
	} else if (dia >= 22 && mes == 5) || (dia <= 21 && mes == 6) {
		fmt.Print("geminis")
	} else if (dia >= 22 && mes == 6) || (dia <= 23 && mes == 7) {
		fmt.Print("cancer")
	} else if (dia >= 24 && mes == 7) || (dia <= 23 && mes == 8) {
		fmt.Print("leo")
	} else if (dia >= 24 && mes == 8) || (dia <= 23 && mes == 9) {
		fmt.Print("virgo")
	} else if (dia >= 24 && mes == 9) || (dia <= 23 && mes == 10) {
		fmt.Print("libra")
	} else if (dia >= 24 && mes == 10) || (dia <= 22 && mes == 11) {
		fmt.Print("escorpio")
	} else if (dia >= 23 && mes == 11) || (dia <= 21 && mes == 12) {
		fmt.Print("sagitario")
	}
}
