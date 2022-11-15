package xerror

type xerr struct {
	err  error
	text string
}

func (e *xerr) Unwrap() error {
	return e.err
}
func (e *xerr) Error() string {
	text := e.text
	err := e.Unwrap()
	for {
		if err == nil {
			break
		}
		switch err.(type) {
		case *xerr:
			text += "\n" + err.(*xerr).text
			err = err.(*xerr).Unwrap()
		default:
			text += "\n" + err.Error()
			err = nil
		}

	}
	return text
}

func New(err error, text string) error {
	return &xerr{
		err:  err,
		text: text,
	}
}
