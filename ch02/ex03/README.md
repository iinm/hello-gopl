実行結果:

```
-> % bash test.bash
=== RUN   TestPopCountByLoop
--- PASS: TestPopCountByLoop (0.00s)
BenchmarkPopCount-4             2000000000               0.32 ns/op            0 B/op          0 allocs/op
BenchmarkPopCountByLoop-4       100000000               20.4 ns/op             0 B/op          0 allocs/op
PASS
ok      _/hello-gopl/ch02/ex03      2.743s
```
