# Types

## Struct

E.g.

```go
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func createBill(name string) bill {
	initBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return initBill
}

// format the bill
func (b *bill) format() string {
	// do something
}
```
## Interface

Interface basically groups types together based on their **methods**

E.g.

```go
// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// declare shapes
var shapes = []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}
```