package condigo

import (
	"testing"
)

func TestWhenThen(t *testing.T) {
	// Case: When the condition is true, it should return the Then result
	result := When(true).Then("Success").Else("Failure")
	if result != "Success" {
		t.Errorf("Expected 'Success', got %v", result)
	}

	// Case: When the condition is false, it should return the Else result
	result = When(false).Then("Success").Else("Failure")
	if result != "Failure" {
		t.Errorf("Expected 'Failure', got %v", result)
	}
}

func TestWhenThenElseIf(t *testing.T) {
	// Case: When the first condition is false but ElseIf condition is true
	result := When(false).Then("First").ElseIf(true, "Second").Else("Default")
	if result != "Second" {
		t.Errorf("Expected 'Second', got %v", result)
	}

	// Case: When both When and ElseIf conditions are false, should return Else
	result = When(false).Then("First").ElseIf(false, "Second").Else("Default")
	if result != "Default" {
		t.Errorf("Expected 'Default', got %v", result)
	}
}

func TestWhenMultipleElseIf(t *testing.T) {
	// Case: Multiple ElseIf with the second condition true
	result := When(false).Then("First").
		ElseIf(false, "Second").
		ElseIf(true, "Third").
		Else("Default")
	if result != "Third" {
		t.Errorf("Expected 'Third', got %v", result)
	}

	// Case: All conditions are false, should return Else result
	result = When(false).Then("First").
		ElseIf(false, "Second").
		ElseIf(false, "Third").
		Else("Default")
	if result != "Default" {
		t.Errorf("Expected 'Default', got %v", result)
	}
}

func TestWithFunction(t *testing.T) {
	res, _ := When(true).
		Then(func() bool { return true }).
		Else(func() bool { return false }).(func() bool)

	if true != res() {
		t.Errorf("Expected 'true', got %v", res())
	}
}
