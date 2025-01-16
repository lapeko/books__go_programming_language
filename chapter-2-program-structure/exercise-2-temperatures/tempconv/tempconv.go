package tempconv

import "fmt"

type (
	Celsius    float32
	Fahrenheit float32
	Kalvin     float32
)

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kalvin) String() string     { return fmt.Sprintf("%g°K", k) }
