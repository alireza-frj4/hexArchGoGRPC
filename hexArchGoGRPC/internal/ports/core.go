package ports

type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	division(a int32, b int32) (int32, error)
}
