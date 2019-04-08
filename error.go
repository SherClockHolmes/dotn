package dotn

import (
	"errors"
	"fmt"
)

var (
	ErrIndexOutOfRange = errors.New("Index is out of range")

	ErrInterfaceTerminated     = errors.New("Interface terminated before end of dot notation string")
	ErrInterfaceNotPointer     = errors.New("Interface must be a pointer")
	ErrInterfaceNotAddressable = errors.New("Interface must be addressable")
)

// KeyNotFoundError is used to inspect an incorrect key
type KeyNotFoundError struct {
	Key string
}

func (k *KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key not foound: %s", k.Key)
}
