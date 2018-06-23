package tempconv

// KToC 絶対温度を摂氏へ変換する
func KToC(k Kelvin) Celsius { return Celsius(float64(AbsoluteZeroC) + float64(k)) }

// KToF 絶対温度を華氏へ変換する
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(float64(k)*9/5 + float64(AbsoluteZeroF)) }

// CToF 摂氏の温度を華氏へ変換する
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK 摂氏の温度を絶対温度に変換する
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToC 華氏の温度を摂氏へ変換する
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK 華氏の温度を絶対温度に変換する
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - AbsoluteZeroF) * 5 / 9) }
