package main

import (
	"fmt"
	"errors"
)

type verifier struct {
	id string
}

func (c *verifier) update(itemName string) {
	fmt.Printf("Sending email to verifier %s for item %s\n", c.id, itemName)
}

func (c *verifier) getID() string {
	return c.id
}

func (c *verifier) verify()