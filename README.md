# Condigo
Simplify your decision-making logic with Condigo, a lightweight Go library providing a fluent interface for conditional statements. Easily define complex logic flows in a readable, linear fashion.

## How to Use Condigo

### Scenario 1: Simple Then & Else

Get a string object based on a single condition.

```go
result := Condigo.When(true).
    Then("Success").
    Else("Failure")
fmt.Println(result) // Output: Success
```

### Scenario 2: Using Then, ElseIf, & Else

Handle multiple conditions with `Then`, `ElseIf`, and `Else`.

```go
result := Condigo.When(false).
    Then("First").
    ElseIf(true, "Second").
    Else("Default")
fmt.Println(result) // Output: Second
```

### Scenario 3: Executing Functions with Then/ElseIf/Else

Execute functions based on conditions and retrieve their return values.

```go
resultFunc, _ := Condigo.When(true).
    Then(func() bool {
        fmt.Println("Inside Then")
        return true
    }).
    Else(func() bool {
        fmt.Println("Inside Else")
        return false
    }).(func() bool)
fmt.Println(resultFunc()) // Output: Inside Then, then true
```

## Example Output

Running the `main.go` file in `examples` folder provided in this repository will output:

```
Success
Second
Inside Then
true
```

## Contributing & Support


- **Issues & Features**: Report bugs or suggest new features by opening an issue on this repository.
- **Pull Requests**: Contributions are welcome. Please submit a pull request for review.
- **Contact**: For direct support, you can try contacting the maintainer on [Linkedin](https://www.linkedin.com/in/spdeepak)