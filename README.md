# 🪨 Moxy

**The Scripting Language Go Developers Already Know.**

Moxy is a high-performance, sandboxed scripting language designed specifically for embedding in Go applications. It combines the simplicity of Go with the flexibility of a dynamic scripting engine.

---

## 🚀 Why Moxy?

If you are a Go developer, you've likely faced the "scripting dilemma":
*   **Lua** is fast but has `1-based` indexing and non-Go syntax.
*   **Starlark** is safe but restrictive (no recursion) and Python-based.
*   **Embedded Go** is powerful but complex to sandbox and heavy.

**Moxy** bridges this gap by providing a **Go-native VM** that uses a syntax you already know.

### Core Value Proposition
- **Go-Like Syntax**: `func`, `var`, `:=`, and `0-based` indexing.
- **Fast Bytecode VM**: Compiles to bytecode for high performance without JIT overhead.
- **Pure Go**: Zero-dependency embedding. No CGO.
- **Safely Sandboxed**: Controlled execution environment for plugins and rules engines.

---

## 🛠 Usage

### In your Go project
```go
import "github.com/yourusername/moxy/package/vm"
import "github.com/yourusername/moxy/package/compiler"

func main() {
    p := parser.New(lexer.New("x := 10; return x + 5;"))
    prog := p.ParseProgram()
    
    comp := compiler.New()
    bytecode := comp.Compile(prog)
    
    machine := vm.New(bytecode)
    machine.Run()
    
    result := machine.LastPoppedStackElem()
    fmt.Println(result) // 15
}
```

### CLI
```bash
./moxy examples/demo.pb
```

---

## 📋 Language Specification (v1.0)

| Feature | Syntax |
|---------|--------|
| **Variables** | `var x = 1` or `x := 1` |
| **Functions** | `func add(a, b) { return a + b }` |
| **Loops** | `while condition { ... }` (Go-style `for` coming soon) |
| **Conditions**| `if x > 10 { ... } else { ... }` |
| **Data Types**| `int`, `string`, `bool`, `array`, `map` |

---

## 🌟 Practical Examples

### 1. Business Rule Engine
```go
// discount_rules.pb
func calculate_discount(order) {
    if order.total > 500 {
        return order.total * 0.1 // 10% off
    }
    return 0
}
```

### 2. Plugin System
```go
// filter.pb
func process(event) {
    if event.type == "metric" && event.value < 0 {
        return null // drop invalid metrics
    }
    return event
}
```

---

## ⚖️ Comparison

| | Moxy | Lua | Starlark |
|---|---|---|---|
| **Syntax** | **Go** | Pascal/C | Python |
| **Indexing** | **0-based** | 1-based | 0-based |
| **Implementation** | **Pure Go** | C (GopherLua is Go) | Go/Java |
| **Performance** | **High** | Extreme (C) | Moderate |

---

## Architecture Diagram

```mermaid
flowchart TD

subgraph group_entrypoints["Entry points"]
  node_cli["CLI<br/>command<br/>[main.go]"]
  node_repl(("REPL<br/>interactive shell<br/>[repl.go]"))
  node_host_api["Host API<br/>embedding surface<br/>[moxy.go]"]
  node_public_pkg["moxy pkg<br/>public package"]
end

subgraph group_frontend["Language frontend"]
  node_token["Tokens<br/>[token.go]"]
  node_lexer["Lexer<br/>scanner<br/>[lexer.go]"]
  node_parser["Parser<br/>[parser.go]"]
  node_ast["AST<br/>syntax tree<br/>[ast.go]"]
end

subgraph group_runtime["Runtime core"]
  node_compiler["Compiler<br/>[compiler.go]"]
  node_symbols["Symbols<br/>name resolution<br/>[symbol_table.go]"]
  node_code["Bytecode<br/>instruction set<br/>[code.go]"]
  node_vm["VM<br/>[vm.go]"]
  node_frame["Frames<br/>[frame.go]"]
  node_stack["Stack<br/>[stack.go]"]
  node_evaluator["Evaluator<br/>tree-walk interpreter<br/>[evaluator.go]"]
  node_types[("Runtime Types<br/>value system<br/>[object.go]")]
end

subgraph group_embedding["Embedding boundary"]
  node_plugin_host["Plugin Host<br/>integration boundary"]
end

subgraph group_examples["Examples"]
  node_std_examples["Std Examples<br/>scripts"]
  node_plugin_example["Plugin Example<br/>host demo<br/>[host.go]"]
end

subgraph group_docs["Docs"]
  node_docs_arch["Design Docs<br/>architecture docs<br/>[VM_ARCHITECTURE.md]"]
end

node_cli -->|"runs"| node_lexer
node_repl -->|"feeds"| node_lexer
node_host_api -->|"embeds"| node_compiler
node_host_api -->|"embeds"| node_evaluator
node_lexer -->|"produces"| node_token
node_lexer -->|"streams"| node_parser
node_token -->|"consumed by"| node_parser
node_parser -->|"builds"| node_ast
node_ast -->|"compiled by"| node_compiler
node_ast -->|"executed by"| node_evaluator
node_compiler -->|"resolves"| node_symbols
node_compiler -->|"emits"| node_code
node_code -->|"loads"| node_vm
node_symbols -->|"supports"| node_vm
node_vm -->|"uses"| node_frame
node_vm -->|"uses"| node_stack
node_vm -->|"creates"| node_types
node_evaluator -->|"creates"| node_types
node_plugin_host -->|"exposes"| node_types
node_plugin_example -->|"demonstrates"| node_plugin_host
node_std_examples -->|"exercises"| node_compiler
node_std_examples -->|"exercises"| node_evaluator
node_docs_arch -->|"describes"| node_vm

click node_cli "https://github.com/pannagaperumal/moxy/blob/master/cmd/moxy/main.go"
click node_repl "https://github.com/pannagaperumal/moxy/blob/master/internal/repl/repl.go"
click node_host_api "https://github.com/pannagaperumal/moxy/blob/master/moxy.go"
click node_public_pkg "https://github.com/pannagaperumal/moxy/tree/master/moxy"
click node_token "https://github.com/pannagaperumal/moxy/blob/master/internal/token/token.go"
click node_lexer "https://github.com/pannagaperumal/moxy/blob/master/internal/lexer/lexer.go"
click node_parser "https://github.com/pannagaperumal/moxy/blob/master/internal/parser/parser.go"
click node_ast "https://github.com/pannagaperumal/moxy/blob/master/ast/ast.go"
click node_compiler "https://github.com/pannagaperumal/moxy/blob/master/internal/compiler/compiler.go"
click node_symbols "https://github.com/pannagaperumal/moxy/blob/master/internal/compiler/symbol_table.go"
click node_code "https://github.com/pannagaperumal/moxy/blob/master/internal/code/code.go"
click node_vm "https://github.com/pannagaperumal/moxy/blob/master/internal/vm/vm.go"
click node_frame "https://github.com/pannagaperumal/moxy/blob/master/internal/vm/frame.go"
click node_stack "https://github.com/pannagaperumal/moxy/blob/master/internal/vm/stack.go"
click node_evaluator "https://github.com/pannagaperumal/moxy/blob/master/internal/evaluator/evaluator.go"
click node_types "https://github.com/pannagaperumal/moxy/blob/master/types/object.go"
click node_plugin_host "https://github.com/pannagaperumal/moxy/tree/master/plugin_host"
click node_std_examples "https://github.com/pannagaperumal/moxy/tree/master/examples/standard_examples"
click node_plugin_example "https://github.com/pannagaperumal/moxy/blob/master/examples/plugin_example/host.go"
click node_docs_arch "https://github.com/pannagaperumal/moxy/blob/master/docs/VM_ARCHITECTURE.md"

classDef toneNeutral fill:#f8fafc,stroke:#334155,stroke-width:1.5px,color:#0f172a
classDef toneBlue fill:#dbeafe,stroke:#2563eb,stroke-width:1.5px,color:#172554
classDef toneAmber fill:#fef3c7,stroke:#d97706,stroke-width:1.5px,color:#78350f
classDef toneMint fill:#dcfce7,stroke:#16a34a,stroke-width:1.5px,color:#14532d
classDef toneRose fill:#ffe4e6,stroke:#e11d48,stroke-width:1.5px,color:#881337
classDef toneIndigo fill:#e0e7ff,stroke:#4f46e5,stroke-width:1.5px,color:#312e81
classDef toneTeal fill:#ccfbf1,stroke:#0f766e,stroke-width:1.5px,color:#134e4a
class node_cli,node_repl,node_host_api,node_public_pkg toneBlue
class node_token,node_lexer,node_parser,node_ast toneAmber
class node_compiler,node_symbols,node_code,node_vm,node_frame,node_stack,node_evaluator,node_types toneMint
class node_plugin_host toneRose
class node_std_examples,node_plugin_example toneIndigo
class node_docs_arch toneTeal
```
---
## 🛤 Roadmap

1.  **Phase 1 (Current)**: VM and Bytecode foundations.
2.  **Phase 2**: Standardize syntax (`func` and `:=` enforcement).
3.  **Phase 3**: Standard library (JSON, Math, Time).
4.  **Phase 4**: Concurrency-lite (channels and fibers).

---

## 🤝 Contributing

We are in early development! Feel free to open issues or PRs. Read our [CONTRIBUTING.md](CONTRIBUTING.md) to get started.
