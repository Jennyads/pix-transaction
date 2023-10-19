package errors

type ErrDuplicated struct {
	Msg string
}

func (e *ErrDuplicated) Error() string {
	return e.Msg
}
