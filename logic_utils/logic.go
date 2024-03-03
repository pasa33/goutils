package logic_utils

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func Pointer[T any](e T) *T {
	return &e
}
