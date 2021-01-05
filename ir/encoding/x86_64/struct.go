package x86_64

import (
	"fmt"

	"github.com/bspaans/jit-compiler/asm/x86_64"
	"github.com/bspaans/jit-compiler/asm/x86_64/encoding"
	"github.com/bspaans/jit-compiler/ir/expr"
	. "github.com/bspaans/jit-compiler/ir/shared"
	"github.com/bspaans/jit-compiler/lib"
)

func encode_IR_Struct(i *expr.IR_Struct, ctx *IR_Context, target encoding.Operand) ([]lib.Instruction, error) {
	// Calculate the displacement between RIP (the instruction pointer,
	// pointing to the *next* instruction) and the address of our byte array,
	// and load the resulting address into target using a LEA instruction.
	ownLength := uint(7)
	diff := uint(ctx.InstructionPointer+ownLength) - uint(i.Address)
	result := []lib.Instruction{x86_64.LEA(&encoding.RIPRelative{encoding.Int32(int32(-diff))}, target)}
	ctx.AddInstructions(result)
	return result, nil
}
func encode_IR_Struct_for_DataSection(b *expr.IR_Struct, ctx *IR_Context, readonly, rw []uint8) ([]uint8, []uint8, error) {
	b.Address = -1
	for _, v := range b.Values {
		bytes := []uint8{}
		if ir, ok := v.(*expr.IR_Uint64); ok {
			bytes = encoding.Uint64(ir.Value).Encode()
		} else if ir, ok := v.(*expr.IR_Int64); ok {
			bytes = encoding.Uint64(ir.Value).Encode()
		} else if ir, ok := v.(*expr.IR_Float64); ok {
			bytes = encoding.Float64(ir.Value).Encode()
		} else {
			return nil, nil, fmt.Errorf("Unsupported struct type %s", v.Type())
		}
		if b.Address == -1 {
			b.Address = len(rw) + 2
		}
		rw = append(rw, bytes...)
	}
	return readonly, rw, nil
}
