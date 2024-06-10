# Linked List Library in Go

This library provides a generic implementation of a circular doubly linked list in Go. It uses Go's new generics feature (available from Go 1.18) to allow for lists of any comparable type.

## Structures

### Node

The `Node` struct represents a node in the linked list. It has three fields:

- `Value`: The value stored in the node. This can be of any type that is comparable (i.e., can be compared with `==` and `!=`).
- `Next`: A pointer to the next node in the list.
- `Prev`: A pointer to the previous node in the list.

### List

The `List` struct represents the linked list itself. It has two fields:

- `sentinel`: A sentinel node to simplify list operations. The sentinel's `Next` field points to the first node in the list, and its `Prev` field points to the last node in the list.
- `length`: The number of elements in the list.

## Functions

### New

The `New` function creates a new empty linked list.

## Installation

To use this library in your Go project, you need to fetch it using the `go get` command:

```bash
go get github.com/FrMnJ/lists
```

## Usage

To use this library, import it in your Go code:

```go
import "github.com/FrMnJ/lists"
```
