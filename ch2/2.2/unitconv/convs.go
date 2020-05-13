package unitconv

//IToF converts inches to feet
func IToF (i Inch) Foot {return Foot(i / 12)}

//IToY converts inches to feet
func IToY (i Inch) Yard {return Yard(i / 12 / 3)}