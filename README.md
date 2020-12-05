# getdoc ![CI](https://github.com/gotd/getdoc/workflows/CI/badge.svg)

Documentation extract utilities for Telegram schema using [goquery](https://github.com/PuerkitoBio/goquery).

## Example
Latest schema is embedded to package, so you can just use it:
```go
doc, err := getdoc.Load(121)
if err != nil {
    panic(err)
}
fmt.Printf("Layer %d, constructors: %d\n", doc.Index.Layer, len(doc.Constructors))
// Output:
// Layer 121, constructors: 851
```

## Reference

Please use [official documentation](https://core.telegram.org/schema), it is for humans,
this package is not.
