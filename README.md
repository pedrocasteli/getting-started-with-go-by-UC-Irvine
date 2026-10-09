# Getting Started with Go

## Advantages of Go

### Runs faster

All software must be translated into the machine language of the processor. This translation step can roughly go in one of two ways:

1. Compilation
   The translation happens once, before the code is executed. Ex.: Go, C, C++, Java (partially).
   Theses languages have a **compiler** that generate the **executable**, which is basically machine language.
   So, when you run the code, you are running machine language, which makes its execution faster.

2. Interpratation
   The code is translated while the code is executed. Meaning that the code is translated every single time it runs. Ex.: Python, Java (partially)
   Theses languages have an **interpreter** that runs alongside the execution of the code, which makes them tipically slower than compiled languages.
   However, the interpreter can help: It can manage memory automatically and infer variable types.

### Garbage collection

Even though Go in a compiled language, it has garbage collection.
It does, however, slows down a bit the execution of the compiled code.

### Simples objects

Go can be consired a **sort of object-oriented language**.
It does not use the term "class". Instead it uses **structs**. But its implementation is a bit simpler than classes in traditional OOP languages:

- No inheritance
- No contructors
- No generics

## Concurrency

Concurrency is the management of multiple tasks at the same time. They might not be executing at the same time, but thay are all "alive" at the same time, maybe on standby.
It is a key requirements for large systems.
Go uses **Goroutines**, **Channels** and **Select** relating to concurrency.
