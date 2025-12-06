# `lookup.go`

A type for looking up the presence of a value. Either a slice or a set depending on the
collection size.

For very small collections (around 10 or less items), it can be faster to find a
value by iterating over a slice rather than using a map. The `Lookup` type abstracts
slices and maps so that you have one interface for both.

## When to use this

- The size of the collection is known and not subject to much change.
- The order of the elements does not matter.
- You are primarily using the collection to look for the presence of a value.

## When to *not* use this

- You will be frequently adding or removing elements from the collection after
  populating its initial values.
- You need to preserve the order of the elements in the collection.

## Example

```go
const sizeThreshold int = 10

// This will be backed by a slice because the number of items is below the threshold.
colorDisplayChoices := lookup.New[string](sizeThreshold, "auto", "on", "off")
colorDisplayChoice := getUserInput()
if !colorDisplayChoices.Has(colorDisplayChoice) {
	fmt.Printf("%q is not a valid choice!\n", colorDisplayChoice)
}

const minYear int = 1900
const maxYear int = 2025
const size int = maxYear - minYear + 1
// This will be backed by a map because the size is above the threshold.
yearChoices := lookup.WithSize[int](sizeThreshold, size)
for i := minYear; i <= maxYear; i++ {
	yearChoices.Add(i)
}
yearChoice := getUserIntInput()
if !yearChoices.Has(yearChoice) {
	fmt.Printf("%d is not a valid year!\n", yearChoice)
}
```
