package tempconv

import "fmt"

type Kelvin float64

func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
