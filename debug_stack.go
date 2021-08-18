//go:build debug

package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type stack struct {
	callers []uintptr
}

func (e *Error) populateStack() {
	e.callers = callers()

	e2, ok := e.Err.(*Error)
	if !ok {
		return
	}

	i := 0

	ok = false
	for ; i < len(e.callers) && i < len(e2.callers); i++ {
		// check for similar
		if e.callers[len(e.callers)-1-i] != e2.callers[len(e2.callers)-1-i] {
			break
		}
		ok = true
	}

	if ok { //we have common PCs
		e2Head := e2.callers[:len(e2.callers)-i]
		eTail := e.callers

		e.callers = make([]uintptr, len(e2Head)+len(eTail))

		copy(e.callers, e2Head)
		copy(e.callers[len(e2Head):], eTail)

		e2.callers = nil
	}
}

// frame returns the nth frame, with the frame at top of stack being 0.
func frame(callers []uintptr, n int) runtime.Frame {
	frames := runtime.CallersFrames(callers)
	var f runtime.Frame
	for i := len(callers) - 1; i >= n; i-- {
		var ok bool
		f, ok = frames.Next()
		if !ok {
			break // Should never happen, and this is just debugging.
		}
	}
	return f
}

// printStack formats and prints the stack for this Error to the given buffer.
// It should be called from the Error's Error method.
func (e *Error) printStack(b *bytes.Buffer) {
	printCallers := callers()

	// Iterate backward through e.callers (the last in the stack is the
	// earliest call, such as main) skipping over the PCs that are shared
	// by the error stack and by this function call stack, printing the
	// names of the functions and their file names and line numbers.
	var prev string // the name of the last-seen function
	var diff bool   // do the print and error call stacks differ now?
	for i := 0; i < len(e.callers); i++ {
		thisFrame := frame(e.callers, i)
		name := thisFrame.Func.Name()

		if !diff && i < len(printCallers) {
			if name == frame(printCallers, i).Func.Name() {
				// both stacks share this PC, skip it.
				continue
			}
			// No match, don't consider printCallers again.
			diff = true
		}

		// Don't print the same function twice.
		// (Can happen when multiple error stacks have been coalesced.)
		if name == prev {
			continue
		}

		// Find the uncommon prefix between this and the previous
		// function name, separating by dots and slashes.
		trim := 0
		for {
			j := strings.IndexAny(name[trim:], "./")
			if j < 0 {
				break
			}
			if !strings.HasPrefix(prev, name[:j+trim]) {
				break
			}
			trim += j + 1 // skip over the separator
		}

		// Do the printing.
		pad(b, Separator)
		_, _ = fmt.Fprintf(b, "%v:%d: ", thisFrame.File, thisFrame.Line)
		if trim > 0 {
			b.WriteString("...")
		}
		b.WriteString(name[trim:])

		prev = name
	}
}

func callers() []uintptr {
	var stk [64]uintptr
	const skip = 4
	n := runtime.Callers(skip, stk[:])
	return stk[:n]
}
