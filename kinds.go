package errors

import "fmt"

type Kind uint16

// General
// 0 - 99
const (
	Undefined Kind = iota
	TimeOut
	Network
)

// general errors
func generalSwitch(k Kind) string {
	switch k {
	case Undefined:
		return "Undefined"
	case TimeOut:
		return "Timeout"
	case Network:
		return "Network"
	default:
		return fmt.Sprintf("The error number is: %d, type is UNDEF", k)
	}
}

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

// error kinds related to the endure
func endureSwitch(k Kind) string {
	switch k {
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
	default:
		return fmt.Sprintf("The error number is: %d, type is UNDEF", k)
	}
}

// RR core kinds
// 200 - 299
const (
	WatcherStopped Kind = iota + 200
	SoftJob
	WorkerAllocate
	NoFreeWorkers
)

// error kinds related to the rr core
func rrSwitch(k Kind) string {
	switch k {
	case WatcherStopped:
		return "Workers watcher stopped"
	case SoftJob:
		return "SoftJobError"
	case NoFreeWorkers:
		return "NoFreeWorkers"
	case WorkerAllocate:
		return "WorkerAllocate"
	default:
		return fmt.Sprintf("The error number is: %d, type is UNDEF", k)
	}
}

// RR plugins kinds
// 300 - 399
const (
	// kv plugin
	EmptyKey Kind = iota + 300
	EmptyItem
	NoKeys
	NoSuchBucket
	BucketShouldBeSet
	NoConfig

	// Reload plugin
	SkipFile
	NoWalkerConfig
)

// error kinds related to the rr plugins
func rrPluginsSw(k Kind) string {
	switch k {
	case SkipFile:
		return "SkipFile"
	case NoWalkerConfig:
		return "NoWalkerConfig"
	case EmptyKey:
		return "key can't be empty string"
	case EmptyItem:
		return "empty Item"
	case NoKeys:
		return "should provide at least 1 key"
	case NoSuchBucket:
		return "no such bucket"
	case BucketShouldBeSet:
		return "bucket should be set"
	case NoConfig:
		return "no config provided"
	default:
		return fmt.Sprintf("The error number is: %d, type is UNDEF", k)
	}
}

func (k Kind) String() string {
	switch {
	case k < 100: // 0-99 general
		return generalSwitch(k)
	case k >= 100 && k < 200: // 100-199, endure
		return endureSwitch(k)
	case k >= 200 && k < 300: // 200-299, rr
		return rrSwitch(k)
	case k >= 300 && k < 400: // 300-399, plugins
		return rrPluginsSw(k)
	default:
		return fmt.Sprintf("The error number is: %d, type is UNDEF", k)
	}
}
