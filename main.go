package main

import (
	"fmt"

	"github.com/he236555699/interfaceTest/pattern"
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

func main() {
	//testAbstractFactory()
	testSimpleFactory()
}
