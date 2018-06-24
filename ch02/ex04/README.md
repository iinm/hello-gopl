実行結果:

```
-> % bash test.bash
=== RUN   TestPopCountByShift
--- PASS: TestPopCountByShift (0.00s)
=== RUN   TestPopCountByLoop
--- PASS: TestPopCountByLoop (0.00s)
BenchmarkPopCount-4             2000000000               0.32 ns/op            0 B/op          0 allocs/op
BenchmarkPopCountByLoop-4       100000000               20.2 ns/op             0 B/op          0 allocs/op
BenchmarkPopCountByShift-4      20000000                80.6 ns/op             0 B/op          0 allocs/op
PASS
ok      _/hello-gopl/ch02/ex04      4.434s
```
