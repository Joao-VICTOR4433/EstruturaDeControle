package main

import "fmt"

func main() {
	var idade int8
	fmt.Scan(&idade)
	if idade>18{
		fmt.Println("Maior de idade!")
	} else{
		fmt.Println("Menor de idade!")

	}
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	var numero int64
	fmt.Scan(&numero)
	if numero%2==0{
		fmt.Println("Número par")
	}else{
		fmt.Println("Número impar")
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	var numeroposiounegativo int64
	fmt.Println("Digite seu numero:")
	fmt.Scan(&numeroposiounegativo)
	if numeroposiounegativo>0{
		fmt.Println("O numero é positivo!")
	}
	if numeroposiounegativo==0{
		fmt.Println("O seu número é zero")
	}else{
		fmt.Println("Seu número é negativo!")
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Digite primeiro o N1 e em suguida o N2")
	var numero1 int64
	var numero2 int64
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	if numero1>numero2{
		fmt.Println("O numero 1 é maior que o numero 2")
	}
	if numero2>numero1{
		fmt.Println("O numero 2 é maior que o numero 1")
	}else{
		fmt.Println("Os  doi são iguais!")
	}

}