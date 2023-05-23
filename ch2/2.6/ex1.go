package main

import (
    "fmt"
    "strconv"
    "os"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gºK", k) }


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5/9) }
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }


func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := Fahrenheit(t)
        c := Celsius(t)
        k := Kelvin(t)
        fmt.Printf("%s = %s, %s = %s\n",
            f, FToC(f), c, CToF(c))
        fmt.Printf("%s = %s, %s = %s\n",
            k, KToC(k), c, CToK(c))
        fmt.Printf("%s = %s, %s = %s\n",
            k, KToF(k), f, FToK(f))
    }
}
