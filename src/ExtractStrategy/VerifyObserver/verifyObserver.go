package main

import (
	"fmt"
	"errors"
)

type Verifier struct {
	id string
}

func (v *Verifier) getID() string {
	return c.id
}

func (v *Verifier) verify()