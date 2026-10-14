# Getting Started with Go

## Advantages of Go

### Runs faster

All software must be translated into the machine language of the processor. This translation step can roughly go in one of two ways:

1. Compilation
   The translation happens once, before the code is executed. Ex.: Go, C, C++, Java (partially).
   Theses languages have a **compiler** that generate the **executable**, which is basically machine language.
   So, when you run the code, you are running machine language, which makes its execution faster.

2. Interpretation
   The code is translated while the code is executed. Meaning that the code is translated every single time it runs. Ex.: Python, Java (partially).
   Theses languages have an **interpreter** that runs alongside the execution of the code, which makes them tipically slower than compiled languages.
   However, the interpreter can help: It can manage memory automatically and infer variable types.

### Garbage collection

Even though Go in a compiled language, it has garbage collection.
It does, however, slows down a bit the execution of the compiled code.

### Simpler objects

Go can be consired a **sort of object-oriented language**.
It does not use the term "class". Instead it uses **structs**. But its implementation is a bit simpler than classes in traditional OOP languages:

- No inheritance;
- No contructors;
- No generics.

## Concurrency

Concurrency is the management of multiple tasks at the same time. They might not be executing at the same time, but thay are all "alive" at the same time, maybe on standby.
It is a key requirements for large systems.
Go uses **Goroutines**, **Channels** and **Select** relating to concurrency.

## Pointers

A pointer is an address to data in memory. There are two main operators that are associated with pointers:

- & (ampersand): Returns the address of a variable/function;
- \*: Does the opposite of the ampersand, returning the data at an address (dereferencing).

It's important to understand the **declaration** of variables of type **pointer**, and its **subsequent use** throughout the code. There is more than one way to declare a pointer, as there is more than one way to declare any type of variable:

```
x := 10
### 1st way
var ptr *int = &x
### 2nd way
ptr := &x
```

To access the value inside that particular address, we use "\*":

```
fmt.Println(*ptr)
```

We can only use both operators in **pointer variables**. Other types only accept "&" to get the address.

## Basic data types

### Arrays

Fixed-length series of elements of a chosen type. If not specified, elements are initialized to zero value.

```
var array [5]int
```

We can also declare them as **array literals**, with pre-defined values:

```
var x = [5]int{1, 2, 3, 4, 5}
y := [...]int{1, 9, 37}
```

### Slices

It's a "window" on an **underlying array**. They have a variable size, up to the whole array.

```
var arr = [...]int{1, 2, 3, 4, 5, 6}
s1 := arr[0:2]
s2 := arr[1:3]
```

Slices can overlap.
`len()` function returns the length, while `cap()` function returns the capacity.
Just like with arrays, we can declare slices as **slice literals**:

```
sli := []int{1, 2, 3}
```
