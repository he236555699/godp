package pattern

import "fmt"

/**
抽象工厂模式
**/
type Worker interface {
	Work(task *string)
}

type WorkCreator interface {
	Create() Worker
}

type Programmer struct {
}

func (p *Programmer) Work(task *string) {
	fmt.Println("Programmer process:", *task)
}

type ProgrammerCreator struct {
}

func (pc *ProgrammerCreator) Create() Worker {
	worker := new(Programmer)
	return worker
}

type Farmer struct {
}

func (f *Farmer) Work(task *string) {
	fmt.Println("Farmer process:", *task)
}

type FarmerCreator struct {
}

func (fc *FarmerCreator) Create() Worker {
	worker := new(Farmer)
	return worker
}
