package main

import (
	"fmt"
	"pebble/package/evaluator"
	"pebble/package/lexer"
	"pebble/package/object"
	"pebble/package/parser"
)

func main() {
	// 1. The code you want to run (e.g., from a config file or DB)
	code := `
		func calculate_tax(amount) {
			return amount * tax_rate
		}
		
		result := calculate_tax(original_price)
	`

	// 2. Prepare the environment and inject host-side variables/functions
	env := object.NewEnvironment()
	evaluator.RegisterBuiltins(env)

	// Inject a variable from the host application
	env.Set("tax_rate", &object.Float{Value: 0.15})
	env.Set("original_price", &object.Integer{Value: 100})

	// 3. Initialize Lexer, Parser, and AST
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
		return
	}

	// 4. Evaluate the code
	evaluated := evaluator.Eval(program, env)

	// 5. Extract results from the environment
	if res, ok := env.Get("result"); ok {
		fmt.Printf("The result is: %s (Type: %s)\n", res.Inspect(), res.Type())

		// If you need the raw Go value:
		if floatRes, ok := res.(*object.Float); ok {
			goValue := floatRes.Value
			fmt.Printf("Extracted Go float: %f\n", goValue)
		}
	}

	if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
		fmt.Printf("Runtime error: %s\n", evaluated.Inspect())
	}
}
