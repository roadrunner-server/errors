package errors

type Kind uint8

// General
// 0 - 99
const (
	Undefined Kind = iota
	TimeOut
)

// Endure Kinds of errors.
// 100 - 199
const (
	Register Kind = iota + 100
	Providers
	Logger
	ArgType
	Init
	Serve
	Unsupported
	Disabled
	Traverse
	FunctionCall
)

// RR kinds
// 200 - 299
const (
	RR Kind = iota + 200 // just sample, replace when add first RR error kind
)

func (k Kind) String() string {
	switch k {
	case Undefined:
		return "UNDEF"
	case Register:
		return "Register error"
	case Providers:
		return "Providers error"
	case Logger:
		return "Logger error"
	case Init:
		return "Init error"
	case Serve:
		return "Serve error"
	case Disabled:
		return "Vertex disabled"
	case ArgType:
		return "Wrong arg type, or return type"
	case Traverse:
		return "Traverse error"
	case FunctionCall:
		return "Function call error"
	case Unsupported:
		return "Unsupported"
	case TimeOut:
		return "Timeout"
	default:
		return "UNDEF"
	}
}
