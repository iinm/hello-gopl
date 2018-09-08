ch07, p.209: 20.0のデフォルト値は°を含んでいないのにヘルプメッセージが°を含んでいる理由は？

```
type celsiusFlag struct{ tempconv.Celsius }
```

1. ヘルプメッセージの作成にはcelsiusFlag.String()が使われる。 https://golang.org/src/flag/flag.go?s=27602:27663#L783
2. `func (*celsiusFlag) String() string` が実装されていないので、埋め込んでいるCelsiusのString()が呼ばれる。
