package ff2

type T2 struct {
	name string
	age int
}

func (t T2) H() string {
	return "T2.H"
}
func (t T2) F() string {
	return "T2.F"
}
