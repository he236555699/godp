package main

import (
	"fmt"

	"github.com/he236555699/godp/pattern"
)

func testAbstractFactory() {
	var workCreator pattern.WorkCreator
	workCreator = new(pattern.ProgrammerCreator)
	taskProject := "Project"
	workCreator.Create().Work(&taskProject)

	workCreator = new(pattern.FarmerCreator)
	taskFarmland := "Farmland"
	workCreator.Create().Work(&taskFarmland)
}

func testSimpleFactory() {
	operateFactory := new(pattern.OperateFactory)
	operator := operateFactory.CreateOperate("+")
	fmt.Println(operator.Operater(1, 2))

	operator = operateFactory.CreateOperate("-")
	fmt.Println(operator.Operater(7, 2))
}

func testFactoryMehtod() {
	aof := new(pattern.AddOperationFactory)
	operation := aof.CreateOperation()
	operation.SetA(5)
	operation.SetB(3)
	fmt.Println(operation.GetResult())

	sof := new(pattern.SubOperationFactory)
	operation = sof.CreateOperation()
	operation.SetA(5)
	operation.SetA(3)
	fmt.Println(operation.GetResult())
}

func main() {
	//testAbstractFactory()
	//testSimpleFactory()
	testFactoryMehtod()
}
