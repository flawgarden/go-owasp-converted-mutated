package controllers

type StaticFieldHolder struct {
	value string
}

var DEFAULT_VALUE = ""

//Конструктор для StaticFieldHolder с использованием DEFAULT_VALUE
func NewStaticFieldHolder() *StaticFieldHolder {
	return &StaticFieldHolder{
		value: DEFAULT_VALUE,
	}
}