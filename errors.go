package IMEKeyboard

type KeyboardErrorType int64

const (
	YamlReaderError KeyboardErrorType = iota
	ReadYamlError
	InitKeyboardError
	DeviceNotMatchError
	InitKeyboardPosError
	Others
)

type KeyboardError struct {
	Message   string `json:"message"`
	ErrorType KeyboardErrorType
}

func (sErr *KeyboardError) Error() string {
	return sErr.Message
}
