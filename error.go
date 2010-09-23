package main

import (
	"fmt"
	"io"
	"os"
	"container/vector"
)

type Error os.Error

func NewError(desc string) Error {
	return os.NewError(desc)
}

type Warnings interface {
	AddWarning(desc string)
	AddWarningFromError(err Error)
	GetWarnings() []string
}

type mWarnings struct {
	messages []string
}

func (w *mWarnings) AddWarning(desc string) {
	if w.messages == nil {
		w.messages = make([]string, 1, 10)
	}

	v := (*vector.StringVector)(&w.messages)
	v.Push(desc)
}

func (w *mWarnings) AddWarningFromError(err Error) {
	w.AddWarning(err.String())
}

func (w *mWarnings) GetWarnings() []string {
	return w.messages
}
