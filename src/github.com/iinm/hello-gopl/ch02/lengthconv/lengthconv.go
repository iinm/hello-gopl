package lengthconv

import "fmt"

type Foot float64
type Metre float64

func (ft Foot) String() string { return fmt.Sprintf("%gft", ft) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }

// MToFT メートルをフィートに変換する
func MToFT(m Metre) Foot { return Foot(m * 3937 / 1200) }

// FTToM メートルをフィートに変換する
func FTToM(ft Foot) Metre { return Metre(ft * 1200 / 3937) }
