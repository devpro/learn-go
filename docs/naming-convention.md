# Naming Convention in Go

## Variables

* Use `camelCase`
* Acronyms should be all capitals, as in `ServeHTTP`
* Single letter represents index: `i`, `j`, `k`
* Short but descriptive names: `cust` not `customer`
* Repeat letters to represent collection, slice, or array and use single letter in loop:

```go
var tt []*Thing
for i, t := range tt {
  // ...
}
```

* Avoid repeating package name:

```go
log.Info()    // good
log.LogInfo() // bad
```

* Don’t name like getters or setters:

```go
custSvc.cust()    // good
custSvc.getCust() // bad
```

* Add `er to Interface:

```go
type Stringer interfaces {
  String() string
}
```

See [Go best practices (Ashley McNamara + Brian Ketelsen)](https://www.youtube.com/watch?v=MzTcsI6tn-0)
