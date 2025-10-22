package main

import "fmt"

func diadaSemana(numero int) string {
	//um switch
	switch numero {
	case 1:
		return "segunda"
	case 2:
		return "terça"
	case 3:
		return "Quarta"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "sabado"
	case 7:
		return "Domingo"
	default:
		return "Número invalido!"

	}

}
func main() {
	fmt.Println(diadaSemana(1))

}