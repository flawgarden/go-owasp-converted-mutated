package controllers

type ImplBinaryOpInterfaceClass1 struct{}

func (i ImplBinaryOpInterfaceClass1) InterfaceCall(l string, r string) string {
    return l
}




type ImplBinaryOpInterfaceClass2 struct{}

func (i ImplBinaryOpInterfaceClass2) InterfaceCall(l string, r string) string {
    return r
}

