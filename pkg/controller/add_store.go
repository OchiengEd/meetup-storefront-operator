package controller

import (
	"meetup/storefront-operator/pkg/controller/store"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, store.Add)
}
