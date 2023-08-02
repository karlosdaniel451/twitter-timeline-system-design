package repositoryerrors


type ErrorNotFound struct {
	Msg string
}

func (err *ErrorNotFound) Error() string {
	return err.Msg
}
