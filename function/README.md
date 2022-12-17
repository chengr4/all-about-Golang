# Funciton

## Pass-by-Value

In general, you should use pass-by-value **when you want to ensure that the original value of a variable is not modified**, or when you are working with small data structures that are not expected to change.

> Remember Golang will copy the reference of `map`, `slice`, `func`

```golang

```

## Receiver Functions

If we want to update `tip` in struct `bill`, we should use **receiver function with pointer**, otherwise, `mybill.tip` will not be updated

```golang
type bill struct {
  name string
  items map[string]float64
  tip float64
}

// receiver function with pointer
func (b *bill) updateTip(tip float64) {
  // it will automatically dereference for us
  b.tip = tip
}

funct main() {
  mybill := {
    name:  name,
    items: map[string]float64{},
    tip:   0,
  }
  mybill.updateTip(50)
}

```


```golang
type bill struct {
  name string
  items map[string]float64
  tip float64
}

// receiver function
func (b bill) format() string {
}

funct main() {
  mybill := {
    name:  name,
    items: map[string]float64{},
    tip:   0,
  }
  mybill.format()
}
```