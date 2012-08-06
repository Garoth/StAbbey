package util

type Error string

func NewError(message string) error {
    var e Error = Error(message)
    return e
}

func (this Error) Error() string {
    return string(this)
}
