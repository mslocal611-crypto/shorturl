package errors

import stderrors "errors"

var (
	ErrInvalidURL = stderrors.New("параметр url должен быть валидным http адресом")
	ErrNotFound   = stderrors.New("url не найден")
	ErrInternal   = stderrors.New("внутренняя ошибка сервера")
)
