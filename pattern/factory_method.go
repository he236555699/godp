package pattern

/*
模式特点：定义一个用于创建对象的接口，让子类决定实例化哪一个类。这使得一个类的实例化延迟到其子类。
程序实例：计算器。
*/
type Operation struct {
	a float64
	b float64
}

type OperationI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

func (o *Operation) SetA(value float64) {
	o.a = value
}

func (o *Operation) SetB(value float64) {
	o.b = value
}

type AddOperation struct {
	Operation
}

func (ao *AddOperation) GetResult() float64 {
	return ao.a + ao.b
}

type SubOperation struct {
	Operation
}

func (so *SubOperation) GetResult() float64 {
	return so.a - so.b
}

type OperationFactory interface {
	CreateOperation() Operation
}

type AddOperationFactory struct {
}

func (af *AddOperationFactory) CreateOperation() OperationI {
	return &AddOperation{}
}

type SubOperationFactory struct {
}

func (sf *SubOperationFactory) CreateOperation() OperationI {
	return &SubOperation{}
}
