package main

import (
    "fmt"
    "strconv"
    "os"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64
type Kilogram float64
type Pound float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gºK", k) }
func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
func (f Feet) String() string { return fmt.Sprintf("%g feet", f) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kilos", k) }
func (p Pound) String() string { return fmt.Sprintf("%g pounds", p) }


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5/9) }
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
func MToF(m Meter) Feet { return Feet(m * 3.28) }
func FToM(f Feet) Meter { return Meter(f * 0.304) }
func KToP(k Kilogram) Pound { return Pound(k * 2.2) }
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2) }


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
        m := Meter(t)
        ft := Feet(t)
        kg := Kilogram(t)
        p := Pound(t)

        fmt.Printf("%s = %s, %s = %s\n",
            f, FToC(f), c, CToF(c))
        fmt.Printf("%s = %s, %s = %s\n",
            k, KToC(k), c, CToK(c))
        fmt.Printf("%s = %s, %s = %s\n",
            k, KToF(k), f, FToK(f))
        fmt.Printf("%s = %s, %s = %s\n",
            m, MToF(m), ft, FToM(ft))
        fmt.Printf("%s = %s, %s = %s\n",
            kg, KToP(kg), p, PToK(p))
    }
}
