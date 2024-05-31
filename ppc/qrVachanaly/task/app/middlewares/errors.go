package middlewares

type CustomError struct {
	message string // Сообщение об ошибке
}

// Реализация метода Error() для типа MyError
func (e *CustomError) Error() string {
	return e.message
}
