package commconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64
type Inches float64
type Centimeter float64

const (
	inchcoeff  Inches = 2.54
	poundcoeff Pound  = 2.20462
)

func (lb Pound) String() string {
	return fmt.Sprintf("%glb", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func (in Inches) String() string {
	return fmt.Sprintf("%g\"", in)
}

func (cm Centimeter) String() string {
	return fmt.Sprintf("%gcm", cm)
}
