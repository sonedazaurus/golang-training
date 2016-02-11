package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.5) }

// KToC converts a Fahrenheit temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.5) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.5)*9/5 + 32) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + 273.5) }
