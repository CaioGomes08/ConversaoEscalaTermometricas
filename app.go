package main

import "fmt"

const escalaK = 373.2

func main() {
	var escalaCelsius = escalaK - 273
	fmt.Printf("O ponto de ebulição da água de Kelvin para Celsius é de: %g", escalaCelsius)
}
