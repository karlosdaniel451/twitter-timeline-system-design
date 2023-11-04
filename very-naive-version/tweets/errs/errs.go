package errs

type ErrorNotFound struct {
	Msg string
}

func (err ErrorNotFound) Error() string {
	return err.Msg
}

type ErrorAlreadyExists struct {
	Msg                      string
	FieldsWithRepeatedValues []string
}

func (err ErrorAlreadyExists) Error() string {
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
