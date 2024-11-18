package controllers

type BaseBinaryOpClass interface{
    VirtualCall(l string, r string) string
}




type DerivedBinaryOpClass1 struct {
    BaseBinaryOpClass
}

func (d DerivedBinaryOpClass1) VirtualCall(l string, r string) string {
    return l
}




type DerivedBinaryOpClass2 struct {
    BaseBinaryOpClass
}

func (d DerivedBinaryOpClass2) VirtualCall(l string, r string) string {
    return r
}




type DerivedBinaryOpClassDefault struct {
    BaseBinaryOpClass
}

func (d DerivedBinaryOpClassDefault) VirtualCall(l string, r string) string {
    return ""
}

