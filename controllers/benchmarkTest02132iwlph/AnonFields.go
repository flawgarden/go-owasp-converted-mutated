package controllers
type Anon struct {
    Value1  string
}

type AnonFieldHolder struct {
    Value2 string
    Anon
}
