package controllers

type StringFactory struct {
	val string
}

//Создает новый экземпляр StringFactory с заданным значением
func CreateStringFactoryWithValue(value string) *StringFactory {
	return &StringFactory{val: value}
}

//Создает новый экземпляр StringFactory с пустым значением
func CreateStringFactory() *StringFactory {
	return &StringFactory{val: ""}
}
