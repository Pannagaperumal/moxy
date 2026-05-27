# 🪨 Moxy - Embeddable Programming Language in Go

> A lightweight interpreted programming language built in Go, designed for embeddability, runtime scripting, and extensibility.

---

# Problem

Most applications eventually require:

- Runtime scripting
- Dynamic business rules
- Modding systems
- Configuration logic
- Embedded automation
- Hot-reloadable behavior

Traditional approaches usually rely on:

- Hardcoded logic
- External scripting runtimes
- Unsafe eval systems
- Heavy embedded languages

This introduces several issues:

- Tight coupling between app logic and runtime behavior
- Difficult extensibility
- Unsafe execution environments
- Increased deployment friction
- Poor control over embedded execution

Moxy solves this by providing:

- A lightweight interpreted language
- Easy Go integration
- Runtime script execution
- Safe sandboxed evaluation
- Embeddable scripting capabilities

It enables developers to add programmable behavior into applications without requiring recompilation.

---

# Architecture

## High-Level Architecture

```text
            Source Code
                 │
                 ▼
        ┌─────────────────┐
        │      Lexer      │
        │ Tokenization    │
        └────────┬────────┘
                 │
                 ▼
        ┌─────────────────┐
        │      Parser     │
        │ Builds AST      │
        └────────┬────────┘
                 │
                 ▼
        ┌─────────────────┐
        │   Abstract      │
        │ Syntax Tree     │
        └────────┬────────┘
                 │
                 ▼
        ┌─────────────────┐
        │    Evaluator    │
        │ Runtime Engine  │
        └────────┬────────┘
                 │
                 ▼
        ┌─────────────────┐
        │ Runtime Objects │
        │ Env / Scope     │
        └─────────────────┘
```

---

# Internal Components

## Lexer

The lexer converts raw source code into tokens.

Example:

```moxy
let x = 10 + 20;
```

Becomes:

```text
LET IDENT ASSIGN INT PLUS INT SEMICOLON
```

Responsibilities:

- Token recognition
- Keyword parsing
- Operator handling
- String/int parsing
- Whitespace skipping

---

## Parser

The parser transforms tokens into an AST (Abstract Syntax Tree).

Example:

```moxy
1 + 2 * 3
```

AST:

```text
      (+)
     /   \
   (1)   (*)
         / \
       (2) (3)
```

Responsibilities:

- Pratt parsing
- Operator precedence
- Expression parsing
- Statement parsing
- Syntax validation

---

## AST Layer

The AST represents the program structure independent of execution.

Node types include:

- Let statements
- Return statements
- Function literals
- Call expressions
- If expressions
- Array literals
- Hash maps
- Loops

This separation makes future compiler/bytecode support easier.

---

## Evaluator

The evaluator walks the AST and executes code.

Responsibilities:

- Expression evaluation
- Scope resolution
- Function execution
- Runtime object creation
- Control flow execution

Supported runtime types:

- Integers
- Booleans
- Strings
- Arrays
- Maps
- Functions
- Null

---

## Environment System

Moxy uses scoped environments for variable resolution.

```text
Global Scope
   │
   ├── Function Scope
   │       │
   │       └── Nested Scope
```

This enables:

- Closures
- Lexical scoping
- Function isolation
- Runtime safety

---

# Language Features

## Core Features

- Variables
- Arithmetic operations
- Boolean expressions
- Conditionals
- Functions
- Arrays
- Hash maps
- Loops
- Built-in functions

---

## Example

```moxy
let fibonacci = fn(n) {
    if (n < 2) {
        return n;
    }

    return fibonacci(n - 1) + fibonacci(n - 2);
};

print(fibonacci(10));
```

---

# Tech Stack

## Core Runtime

- Go
- Custom AST
- Pratt Parser
- Recursive Evaluator

---

## Tooling

- Go Modules
- CLI Runtime
- REPL Support

---

# Scaling Considerations

Although Moxy is currently an interpreted language, the architecture is intentionally designed for future scalability.

---

## AST Separation

The AST layer enables future support for:

- Bytecode compilation
- JIT execution
- Static analysis
- Optimization passes

without rewriting the parser.

---

## Embeddable Runtime

Moxy can be embedded directly into Go applications:

```go
engine := moxy.New()

engine.Execute(script)
```

This enables:

- Game scripting
- Plugin systems
- Workflow engines
- Runtime automation

---

## Sandboxed Execution

The runtime is intentionally isolated from direct system access.

Benefits:

- Safer embedded execution
- Controlled APIs
- Reduced attack surface

Future capability-based permissions can extend this further.

---

# Tradeoffs

## Advantages

- Lightweight runtime
- Simple architecture
- Easy embeddability
- Beginner-friendly interpreter design
- Extensible language core

---

## Limitations

- Interpreted execution is slower than compiled runtimes
- No static typing
- No bytecode VM yet
- Limited standard library
- Garbage collection relies on Go runtime

---

# Example Workflow

## Example Program

```moxy
let nums = [1, 2, 3, 4];

let sum = fn(arr) {
    let total = 0;

    for (x in arr) {
        total = total + x;
    }

    return total;
};

print(sum(nums));
```

---

## Execution Flow

### Step 1 — Lexing

Source code becomes tokens.

---

### Step 2 — Parsing

Tokens are transformed into AST nodes.

---

### Step 3 — Evaluation

Evaluator recursively walks the AST.

---

### Step 4 — Runtime Execution

Objects/functions/scopes are created dynamically.

---

### Step 5 — Output

```text
10
```

---

# Future Roadmap

## Short Term

- Better error handling
- Improved REPL
- Module system
- File imports
- Standard library expansion

---

## Mid Term

- Bytecode compiler
- Virtual machine (VM)
- Performance optimizations
- Debugger support
- Package manager

---

## Long Term

- JIT compilation
- WASM target
- Concurrency primitives
- Native FFI bindings
- Language server support (LSP)

---

# Why Moxy Matters

Most embedded scripting systems are either:

- Too heavy
- Too unsafe
- Too difficult to customize

Moxy aims to provide:

> A lightweight, hackable, embeddable language runtime for modern Go applications.

It serves as both:

- A practical embeddable scripting engine
- A deep exploration into language/runtime engineering
- A foundation for future VM and compiler experimentation
