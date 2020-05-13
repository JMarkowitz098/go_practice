package unitconv

import "fmt"

type Inch float64
type Foot float64
type Yard float64

func (i Inch) String() string { return fmt.Sprintf("%gin.", i) }

func (f Foot) String() string { return fmt.Sprintf("%gft.", f) }

func (y Yard) String() string { return fmt.Sprintf("%gyd.", y) }
