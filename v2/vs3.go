package ff2

import "github.com/itsunixiknowthis/ff1"

type T2 struct {
	ff1.T1
	name string
	age int
}

func (t T2) H() string {
	return "T2.G"
}

func (t T2) F() string {
	return "T2.F"
}
