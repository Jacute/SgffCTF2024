package qrGenerator

type ErrorLength struct {
	Message string // Сообщение об ошибке
}

// Реализация метода Error() для типа MyError
func (e *ErrorLength) Error() string {
	return e.Message
}
