package ports

type DbPort interface {
	CloseDbConnecton()
	AddToHistory(answer int32, aperation string) error
}
