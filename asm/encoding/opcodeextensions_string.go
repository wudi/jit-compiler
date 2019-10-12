// Code generated by "stringer -type=OpcodeExtensions"; DO NOT EDIT.

package encoding

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoExtensions-0]
	_ = x[ImmediateByte-1]
	_ = x[ImmediateWord-2]
	_ = x[ImmediateDouble-3]
	_ = x[Slash0-4]
	_ = x[Slash1-5]
	_ = x[Slash2-6]
	_ = x[Slash3-7]
	_ = x[Slash4-8]
	_ = x[Slash5-9]
	_ = x[Slash6-10]
	_ = x[Slash7-11]
	_ = x[SlashR-12]
	_ = x[RexW-13]
}

const _OpcodeExtensions_name = "NoExtensionsImmediateByteImmediateWordImmediateDoubleSlash0Slash1Slash2Slash3Slash4Slash5Slash6Slash7SlashRRexW"

var _OpcodeExtensions_index = [...]uint8{0, 12, 25, 38, 53, 59, 65, 71, 77, 83, 89, 95, 101, 107, 111}

func (i OpcodeExtensions) String() string {
	if i < 0 || i >= OpcodeExtensions(len(_OpcodeExtensions_index)-1) {
		return "OpcodeExtensions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpcodeExtensions_name[_OpcodeExtensions_index[i]:_OpcodeExtensions_index[i+1]]
}
