// tempconvパッケージは摂氏 (Celsius) と華氏 (Fahrenheit) の温度変換を行う。
package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius    = -273.15
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
)

func (c Kelvin) String() string     { return fmt.Sprintf("%gK", c) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
