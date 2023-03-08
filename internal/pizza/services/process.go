package services

import "fmt"

type Process struct {
}

func NewProcessImpl() *Process {
	return &Process{}
}

func (p *Process) PrintMsj(msj string) error {
	fmt.Println("Process...")
	fmt.Println(msj)
	return nil
}
