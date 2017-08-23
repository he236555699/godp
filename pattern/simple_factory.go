package pattern

import "fmt"

type Operater interface {
	Operater(int, int) int
}

type AddOperate struct {
}

func (ao *AddOperate) Operater(a int, b int) int {
	return a + b
}

type SubOperate struct {
}

func (so *SubOperate) Operater(a int, b int) int {
	return a - b
}

type OperateFactory struct {
}

func (of *OperateFactory) CreateOperate(operateName string) (operator Operater) {
	switch operateName {
	case "+":
		operator = &AddOperate{}
	case "-":
		operator = &SubOperate{}
	default:
		fmt.Println("invalidation Operate")
		operator = nil
	}

	return
}
