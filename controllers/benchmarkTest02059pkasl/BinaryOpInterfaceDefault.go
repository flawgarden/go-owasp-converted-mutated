package controllers

type BinaryOpInterface interface {
    InterfaceCall(l, r string) string
}




type BinaryOpInterfaceDefaultBase interface {
    InterfaceCall(l, r string) string
}




type BinaryOpInterfaceDefault1 interface{
    InterfaceCall(l, r string) string
}




type BinaryOpInterfaceDefault1Implementation struct{}

func (b BinaryOpInterfaceDefault1Implementation) InterfaceCall(l, r string) string {
    return l
}




type BinaryOpInterfaceDefault2 interface{
    InterfaceCall(l, r string) string
}




type BinaryOpInterfaceDefault2Implementation struct {
    BinaryOpInterfaceDefault2
}

func (b BinaryOpInterfaceDefault2Implementation) InterfaceCall(l, r string) string {
    return r
}




type BinaryOpInterfaceDefault interface {
    BinaryOpInterfaceDefault1
    BinaryOpInterfaceDefault2
}




type BinaryOpInterfaceDefaultImplementation struct{}

func (b BinaryOpInterfaceDefaultImplementation) InterfaceCall(l, r string) string {
    return ""
}

