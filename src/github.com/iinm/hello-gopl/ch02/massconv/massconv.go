package massconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }

// KgToLB キログラムをポンドに変換する
func KgToLB(kg Kilogram) Pound { return Pound(kg * 2.20462) }

// LBToKg ポンドをキログラムに変換する
func LBToKg(lb Pound) Kilogram { return Kilogram(lb * 0.453592) }
