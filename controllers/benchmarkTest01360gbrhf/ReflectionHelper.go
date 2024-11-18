package controllers

type ReflectionHelper struct {
    value string
}

func NewReflectionHelper(value string) *ReflectionHelper {
    return &ReflectionHelper{value: value}
}

func (rh *ReflectionHelper) GetValue() string {
    return rh.value
}

