# Form Benchmark

PC Specification:

AMD FX-9580 4.7Ghz, 32GB RAM 1066MHz

Go Version:

`go version go1.6.2 windows/amd64`

## Test 1: Simple

`$ $ GOGC=100 go test -benchmem -bench=Simple$`

```
BenchmarkSchemaSimple-8          1000000              1633 ns/op             152 B/op          6 allocs/op
BenchmarkFormamSimple-8          2000000               777 ns/op              40 B/op          3 allocs/op
BenchmarkCjToolkitFormSimple-8   1000000              1772 ns/op             272 B/op          3 allocs/op
```

`$ GOGC=800 go test -benchmem -bench=Simple$`

```
BenchmarkSchemaSimple-8          1000000              1016 ns/op             152 B/op          6 allocs/op
BenchmarkFormamSimple-8          2000000               666 ns/op              40 B/op          3 allocs/op
BenchmarkCjToolkitFormSimple-8   2000000               851 ns/op             272 B/op          3 allocs/op
```

## Test 2: Simple with unused data

`$ GOGC=100 go test -benchmem -bench=SimpleWithUnusedData$`

```
BenchmarkSchemaSimpleWithUnusedData-8            100000             11152 ns/op            1112 B/op         43 allocs/op
BenchmarkFormamSimpleWithUnusedData-8           1000000              1696 ns/op             146 B/op          6 allocs/op
BenchmarkCjToolkitFormSimpleWithUnusedData-8    1000000              1880 ns/op             272 B/op          3 allocs/op
```

`$ GOGC=800 go test -benchmem -bench=SimpleWithUnusedData$`

```
BenchmarkSchemaSimpleWithUnusedData-8            200000              7030 ns/op            1112 B/op         43 allocs/op
BenchmarkFormamSimpleWithUnusedData-8           1000000              1093 ns/op             145 B/op          6 allocs/op
BenchmarkCjToolkitFormSimpleWithUnusedData-8    2000000               854 ns/op             272 B/op          3 allocs/op
```

## Test 3: Complex

`$ GOGC=100 go test -benchmem -bench=Complex$`

```
BenchmarkSchemaComplex-8         1000000              4238 ns/op             360 B/op         18 allocs/op
BenchmarkFormamComplex-8         1000000              2590 ns/op             128 B/op         12 allocs/op
BenchmarkCjToolkitFormComplex-8  1000000              3161 ns/op             272 B/op          3 allocs/op
```

`$ GOGC=800 go test -benchmem -bench=Complex$`

```
BenchmarkSchemaComplex-8         1000000              2948 ns/op             360 B/op         18 allocs/op
BenchmarkFormamComplex-8         1000000              2165 ns/op             128 B/op         12 allocs/op
BenchmarkCjToolkitFormComplex-8  1000000              2075 ns/op             272 B/op          3 allocs/op
```

## Test 4: Complex with unused data

`$ GOGC=100 go test -benchmem -bench=ComplexWithUnusedData$`

```
BenchmarkSchemaComplexWithUnusedData-8           100000             14338 ns/op            1320 B/op         55 allocs/op
BenchmarkFormamComplexWithUnusedData-8          1000000              2086 ns/op             177 B/op         10 allocs/op
BenchmarkCjToolkitFormComplexWithUnusedData-8   1000000              3160 ns/op             272 B/op          3 allocs/op
```

`$ GOGC=800 go test -benchmem -bench=ComplexWithUnusedData$`

```
BenchmarkSchemaComplexWithUnusedData-8           200000              9550 ns/op            1320 B/op         55 allocs/op
BenchmarkFormamComplexWithUnusedData-8          1000000              1910 ns/op             192 B/op         11 allocs/op
BenchmarkCjToolkitFormComplexWithUnusedData-8   1000000              2159 ns/op             272 B/op          3 allocs/op
```

## Conclusion

As the benchmark above shows, [CJToolkit Form 3.0](https://github.com/cjtoolkit/form) performs reasonably well against [Gorilla Schema](http://www.gorillatoolkit.org/pkg/schema) and [Formam](https://github.com/monoculum/formam)

In most cases it's outperformed Schema and in some cases it can outperform Formam, especially if you tune up the Garbage Collection to `800` (default is `100`) than CJToolkit Form get faster most likely because it's consistently `3 allocs/op`.

Introducing unused data to `url.Values` (or even `HTTP POST`) will cause a negative impact to Schema, thus memory allocation goes up dramatically, therefore performance goes down dramatically and this is quite scary, imo that a very serious design flaw on Gorilla part, I mean any user can tamper with HTTP POST and cause havoc.  But only has little to no impact with CJToolkit Form and Formam.

CJToolkit Form has got a bigger codebase than the other two, but the codes are highly maintainable, testable and make no use of `reflect`, while the other two uses `reflect`.  I firmly believe that a robust, easy-to-test and modular designed codebase is much better than having a quick and dirty codebase with the abuse of `reflect`.

I choosed not to use `reflect` because I just don't like the idea of using that for something as simple as `map[string][]string`.