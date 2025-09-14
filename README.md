# Stack

Stack data structure using slice (thread-safe)

```go
type stack struct {
	items  []itemType
	rwLock sync.RWMutex
}
```

`itemType` is a `any` type so that it can take any type of value

```go
type itemType any
```