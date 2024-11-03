package Condigo

type Conditional struct {
	condition bool
}

type ElseIf struct {
	returned   bool
	lastResult interface{}
}

func When(condition bool) *Conditional {
	return &Conditional{condition: condition}
}

func (c *Conditional) Then(result interface{}) *ElseIf {
	if c.condition {
		return &ElseIf{returned: true, lastResult: result}
	}
	return &ElseIf{returned: false, lastResult: nil}
}

func (ei *ElseIf) ElseIf(condition bool, result interface{}) *ElseIf {
	if !ei.returned && condition {
		return &ElseIf{returned: true, lastResult: result}
	}
	return &ElseIf{returned: ei.returned, lastResult: ei.lastResult}
}

func (ei *ElseIf) Else(defaultResult interface{}) interface{} {
	if !ei.returned {
		return defaultResult
	}
	return ei.lastResult
}
