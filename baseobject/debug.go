package baseobject

var debugC ConsoleDebug

type ConsoleDebug interface {
	Debug(opts ...interface{}) error
}

func SetConsoleDebug(obj interface{}) {

	if c, ok := obj.(ConsoleDebug); ok {
		debugC = c
	}
}

func (b BaseObject) Debug(msg string) error {
	if debugC != nil {
		return debugC.Debug(msg)
	}
	return nil
}
