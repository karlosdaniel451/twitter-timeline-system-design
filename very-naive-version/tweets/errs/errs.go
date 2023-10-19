package errs

type ErrorNotFound struct {
	Msg string
}

func (err ErrorNotFound) Error() string {
	return err.Msg
}

type ErrorUserAlreadyFollow struct {
	Msg string
}

func (err ErrorUserAlreadyFollow) Error() string {
	return err.Msg
}

type ErrorUserDoesNotFollow struct {
	Msg string
}

func (err ErrorUserDoesNotFollow) Error() string {
	return err.Msg
}
