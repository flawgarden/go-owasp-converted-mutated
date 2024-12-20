package controllers

type ArrayHolder struct {
    Values []string
}

func NewArrayHolder(value string) *ArrayHolder {
    return NewArrayHolderWithValues([]string{value, ""})
}

func NewArrayHolderWithValues(initialValues []string) *ArrayHolder {
    return &ArrayHolder{Values: initialValues}
}

