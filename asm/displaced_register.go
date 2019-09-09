package asm

import "fmt"

type DisplacedRegister struct {
	*Register
	Displacement uint8
}

func (t *DisplacedRegister) Type() Type {
	return T_DisplacedRegister
}

func (t *DisplacedRegister) String() string {
	return fmt.Sprintf("0x%x(%s)", t.Displacement, t.Register.String())
}
