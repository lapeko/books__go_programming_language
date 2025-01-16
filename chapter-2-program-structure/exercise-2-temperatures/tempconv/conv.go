package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToC(k Kalvin) Celsius     { return Celsius(float32(AbsoluteZeroC) + float32(k)) }
