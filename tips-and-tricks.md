# Tips & Tricks

## Arrays, Slices, Iterations

```go
// Create an array of 4 integers
arr := [4]int{1, 2, 3, 4}

// Create a slice of integers with length of 9 elements (and capacity of 9 elements)
sl := make([]int, 9)

// Create a slice of integers (with pre-defined values)
sl := []int{1, 2, 3, 4}

// Iterate over the slice from left (start) to right (end)...
for i := 0; i < len(sl); i++ {
    v := sl[i]
    fmt.Printf("%d ", v)
}
// Output: 1 2 3 4

// Iterate over the slice right (end) to left (start)... note how we initialize "i"
for i := len(sl) - 1; i >= 0; i-- {
    v := sl[i]
    fmt.Printf("%d ", v)
}
// Output: 4 3 2 1 

// Iterate over the slice from left (start) to right (end) using "range" keyword
for i, v := range sl {

}

// Access element at index "i"
v := arr[i]
v := sl[i]

// Append element at the end of a slice 
sl = append(sl, 5)
```

## Map (Hash Table)

```go
// Create a hash map of key of type "string", and value of type "int"
freq := make(map[string]int)

// Create a hash map with predefined keys and values
freq := map[string]int{
    "Ada": 22,
    "Mike": 69,
    "Kate": 45,
}

// Insert element at key "Todd"
freq["Todd"] = 66

// Check if element exists using "ok" operator
// If exists "ok" will evaluate to "true"
v, ok := freq["Todd"]
if ok {
    // do something...
}
```

## Bit Manipulation

We can initialize variables using multiple notations: decimal, binary, hexadecimal or octal.

```go
// Initialize variable using binary notation (prefix "0b_")
n := 0b_1001

// Initialize variable using hexadecimal notation (prefix "0x_")
n := 0x_f00f

// Initialize variable using octal notation (prefix "0o_")
n := 0o_777
```

Go has multiple build-in bit operators: AND, OR, XOR, NOT, AND NOT

[Reference](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/)

```go
0011  & 0101     0001        AND
0011  | 0101     0111        OR
0011  ^ 0101     0110        XOR
      ^ 0101     1010        NOT (same as 1111 ^ 0101)
0011 &^ 0101     0010        AND NOT
```

We can shift bits left and right:

```go
00110101 << 2    11010100    Left shift
00110101 << 100  00000000    No upper limit on shift count 
00110101 >> 2    00001101    Right shift 
```

We can manipulate individual bits:

```go
// Initialize variable of type uint
n := uint(0)

// Clears the bit at position "pos" in number "n"
n &^= (1 << pos)

// Sets the bit at position "pos" in number "n"
n |= (1 << pos)

// Check if the bit at position "pos" is set in number "n"
val := n & (1 << pos)
```

Above snippets were taken from this answer at [Stack Overflow](https://stackoverflow.com/a/23192263/1449403).

## Conversion between a digit and a character

```go
num1 := '9'
val := byte(num1) - '0'
fmt.Printf("%d", val)
// Output: 9

num2 = byte(val) + '0'
fmt.Printf("%q", num2)
// Output: '9'
```

More complicated conversions are made with "strconv" package (see below).

## Snippets

### Count sum of all digits of a number:

```go
func sumOfDigits(num int) int {
    sum := 0
    for num>0 {
        sum += num % 10
        num /= 10
    }
    return sum
}
```

### Reverse slice of bytes:

```go
func reverseSlice(s []byte, n int) []byte {
    for i, j := 0, n; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}
```

## Package "math/bits"

[Package Documentation](https://pkg.go.dev/math/bits)

Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.

Functions in this package may be implemented directly by the compiler, for better performance. For those functions the code in this package will not be used. Which functions are implemented by the compiler depends on the architecture and the Go release.

```go
i := 0xf00

// OnesCount returns the number of one bits ("population count") in "i".
v := bits.OnesCount(uint(i))
```

## Package "sort"

[Package Documentation](https://pkg.go.dev/sort)

Package sort provides primitives for sorting slices and user-defined collections.

```go
sl := []int{3, 5, 6}

// Ints sorts a slice of ints in increasing order.
sort.Ints(sl)

// Slice sorts the slice "sl" given the provided less function. It panics if "sl" is not a slice.
sort.Slice(sl, func(i, j int) bool {
    return sl[i] < sl[j]
})

```

## Package "strings"

[Package Documentation](https://pkg.go.dev/strings)

Package strings implements simple functions to manipulate UTF-8 encoded strings.

```go
// Contains reports whether substr is within s.
ok := strings.Contains("seafood", "foo")

// Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
count := strings.Count("cheese", "e")

// Fields splits the string s around each instance of one or more consecutive white space characters(...)
words := strings.Fields(str)
```

## Package "strconv"

[Package Documentation](https://pkg.go.dev/strconv)

Package strconv implements conversions to and from string representations of basic data types.

```go
// Ascii to integer conversion
i, err := strconv.Atoi("-42")

// Integer to Ascii conversion
s := strconv.Itoa(-42)
```

## Package "time"

[Package Documentation](https://pkg.go.dev/time)

Package time provides functionality for measuring and displaying time.

```go
// Parse parses a formatted string and returns the time value it represents.
t, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
```

### ʕ◔ϖ◔ʔ
