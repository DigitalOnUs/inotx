package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Contains-0]
	_ = x[Egress-1]
	_ = x[Ingress-2]
}

const _AssociationType_name = "containsegressingress"

var _AssociationType_index = [...]uint8{0, 8, 14, 21}

func (i AssociationType) String() string {
	if i < 0 || i >= AssociationType(len(_AssociationType_index)-1) {
		return "AssociationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AssociationType_name[_AssociationType_index[i]:_AssociationType_index[i+1]]
}
