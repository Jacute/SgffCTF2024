package utils

type JwtError struct {
	message string // Сообщение об ошибке
}

// Реализация метода Error() для типа MyError
func (e *JwtError) Error() string {
	return e.message
}
