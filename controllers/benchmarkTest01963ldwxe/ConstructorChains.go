package controllers

type ConstructorChains struct {
    condition bool
    text      string
}

func NewConstructorChains() *ConstructorChains {
    return NewConstructorChainsWithParams(true, "")
}

func NewConstructorChainsWithText(text string) *ConstructorChains {
    return NewConstructorChainsWithParams(true, text)
}

func NewConstructorChainsWithParams(condition bool, text string) *ConstructorChains {
    if !condition {
        text = ""
    }
    return &ConstructorChains{condition: condition, text: text}
}

func (c *ConstructorChains) GetText(condition bool) string {
    if c.condition || condition {
        return c.text
    }
    return ""
}

