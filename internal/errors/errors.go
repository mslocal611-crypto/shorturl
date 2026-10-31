package errors

import stderrors "errors"

var (
	ErrInvalidURL = stderrors.New("тело запроса должно быть валидным URL-адресом")
	ErrNotFound   = stderrors.New("url не найден")
	ErrInternal   = stderrors.New("внутренняя ошибка сервера")
)
