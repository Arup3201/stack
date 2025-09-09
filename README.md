# Stack

Stack data structure using Linked List (thread-safe)

```go
type Stack struct {
	lock *sync.Mutex
	top  *Element
	Size int
}
```

`Element` is a linked list node with type `any` so that it can take any type of value

```go
type Element struct {
	Data any
	next *Element
}
```