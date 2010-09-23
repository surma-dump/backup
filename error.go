package main

import (
	"os"
	"container/vector"
)

type Error struct {
	message string
}
// Making it comply with os.Error
func (e *Error) String() string {
	return e.message
}

type Warnings struct {
	Messages []string
}

func (e *Error) NewError(desc string) {
	if e == nil {
		e = new(Error)
	}
	e.message = desc
}

func (e *Error) NewErrorFromError(err os.Error) {
	if e == nil {
		e = new(Error)
	}
	e.message = err.String()
}

func (w *Warnings) NewWarning(desc string) {
	if w == nil {
		w.Messages = make([]string, 1, 3)
	}

	v := (*vector.StringVector)(&w.Messages)
	v.Push(desc)
}
