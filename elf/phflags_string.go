// Code generated by "stringer -type=PHFlags"; DO NOT EDIT.

package elf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PF_X-1]
	_ = x[PF_W-2]
	_ = x[PF_WX-3]
	_ = x[PF_R-4]
	_ = x[PF_RX-5]
	_ = x[PF_RW-6]
	_ = x[PF_RWX-7]
	_ = x[PF_MASKPROC-4026531840]
}

const (
	_PHFlags_name_0 = "PF_XPF_WPF_WXPF_RPF_RXPF_RWPF_RWX"
	_PHFlags_name_1 = "PF_MASKPROC"
)

var (
	_PHFlags_index_0 = [...]uint8{0, 4, 8, 13, 17, 22, 27, 33}
)

func (i PHFlags) String() string {
	switch {
	case 1 <= i && i <= 7:
		i -= 1
		return _PHFlags_name_0[_PHFlags_index_0[i]:_PHFlags_index_0[i+1]]
	case i == 4026531840:
		return _PHFlags_name_1
	default:
		return "PHFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
