package services

import (
	"log"
)

type Decorator struct {
	process IProcess
}

func NewDecoratorImp(process IProcess) *Decorator {
	return &Decorator{
		process: process,
	}
}

func (d *Decorator) PrintMsj(msj string) error {
	log.Println("Hey, I'm a log!")
	return d.process.PrintMsj(msj)
}
