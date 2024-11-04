package condigo

// Conditional represents the initial condition that is evaluated.
// If the condition is true, Then logic will execute.
type Conditional struct {
	condition bool
}

// ElseIf represents the state after a Then or ElseIf condition.
// It keeps track of whether a result has been returned and holds the last result.
type ElseIf struct {
	returned   bool        // Tracks if a condition has already been met
	lastResult interface{} // Stores the result to return if this condition is met
}

// When initializes a Conditional with a given condition.
// Returns a pointer to Conditional for further chaining.
func When(condition bool) *Conditional {
	return &Conditional{condition: condition}
}

// Then checks if the initial condition was true.
// If so, it stores the result and marks the condition as met by setting returned to true.
// Otherwise, it continues to the next condition (e.g., ElseIf).
func (c *Conditional) Then(result interface{}) *ElseIf {
	if c.condition {
		return &ElseIf{returned: true, lastResult: result}
	}
	return &ElseIf{returned: false, lastResult: nil}
}

// ElseIf checks an additional condition if previous conditions (When or earlier ElseIf)
// have not been met. If this condition is true, it stores the result and marks
// the condition as met. If not, it continues to the next condition.
func (ei *ElseIf) ElseIf(condition bool, result interface{}) *ElseIf {
	if !ei.returned && condition {
		return &ElseIf{returned: true, lastResult: result}
	}
	return &ElseIf{returned: ei.returned, lastResult: ei.lastResult}
}

// Else provides a default result if none of the previous conditions (When, Then, or ElseIf) were met.
// If no conditions were met, it returns the defaultResult.
// If any condition was met, it returns the result of the first matched condition.
func (ei *ElseIf) Else(defaultResult interface{}) interface{} {
	if !ei.returned {
		return defaultResult
	}
	return ei.lastResult
}
