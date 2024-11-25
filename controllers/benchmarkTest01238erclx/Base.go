package controllers
type Base struct {
    Value string
}

type Derived struct {
    Base
}

func (b Base) Concat(s1 string, s2 string) string {
	return s1 + s2
}

type Derived2 struct {
    Base
    Value string
}