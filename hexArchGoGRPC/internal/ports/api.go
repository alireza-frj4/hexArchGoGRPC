package ports

type APIPorts interface {
	GetAddition(a int32, b int32) (int32, error)
	GetSubtraction(a int32, b int32) (int32, error)
	GetMultiplication(a int32, b int32) (int32, error)
}
