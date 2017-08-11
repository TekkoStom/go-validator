# GO Validator

## Intro

Since we are trying to simplify struct validation as much as possible, we have started
writing our own validation library.

At the moment the library is quite limited, but as we progress with our product,
we are going to also update the library.

## Usage

```
// our struct that we are going to validate
type Test struct {
    Name string
    Title string
}

func main() {
    validator := Validator{Entity: test}
    validator.Required("Name").
        Length("Name", 1, 10)
    
    // returns a map[string][]string that contains
    // all errors for each validated field
    log.Println(validator.ValidationErrors)  
}

```