package main

import (
	"decoradorLogs/internal/pizza/services"
	"fmt"
	"log"
)

func main() {

	var msj string

	fmt.Println("Project: Decorator")
	fmt.Println("Digite un mensaje:")
	_, err := fmt.Scanln(&msj)
	if err != nil {
		log.Println("Error")
		return
	}

	service := services.NewDecoratorImp(services.NewProcessImpl())
	err = service.PrintMsj(msj)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
