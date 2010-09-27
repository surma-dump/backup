package main

import (
	"os"
	"container/vector"
)

type Error os.Error

func NewError(desc string) Error {
	return os.NewError(desc)
}

type Warnings []string

func (w *Warnings) AddWarning(desc string) {
	if w == nil {
		t := make([]string, 1, 4)
		w = (*Warnings)(&t)
	}

	v := (*vector.StringVector)(w)
	v.Push(desc)
}

func (w *Warnings) AddWarningFromError(err Error) {
	w.AddWarning(err.String())
}
