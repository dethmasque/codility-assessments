## Properly Nested Strings — Problem Description

A string `S` consisting of `N` characters is considered to be **properly nested** if any of the following conditions is true:

- `S` is empty;
- `S` has the form ``(U)``, ``[U]``, or ``{U}`` where `U` is a properly nested string;
- `S` has the form `VW` where `V` and `W` are properly nested strings.

### Examples:
- The string ``{[()()]}`` is properly nested.
- The string ``([)()]`` is **not** properly nested.

---

## Task

Write a function:

```go
func Solution(S string) int
```

that, given a string `S` consisting of `N` characters, returns:

- `1` if `S` is properly nested,
- `0` otherwise.

---

## Examples

- Given `S = "{[()()]}"`, the function should return `1`.
- Given `S = "([)()]"`, the function should return `0`.

---

## Constraints

- `N` is an integer within the range `[0..200,000]`;
- String `S` contains only the following characters:  
  `'('`, `')'`, `'{'`, `'}'`, `'['`, `']'`

