# Types

## Struct

E.g.

```go
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// format the bill
func (b *bill) format() string {
	// do something
}
```
