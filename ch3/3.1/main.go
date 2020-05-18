// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
    "os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 30                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	//Create a file
	f, err := os.Create("./data.svg")
	check(err)
	defer f.Close()
		
	//Put SVG in file
	n3, err := f.WriteString(svg)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if math.IsNaN(ay) || math.IsNaN(by) || math.IsNaN(cy) || math.IsNaN(dy) {
				continue
			}
			
			polyPoint := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)

			n3, err := f.WriteString(polyPoint)
			check(err)
			fmt.Printf("wrote %d bytes\n", n3)

			
		}
	}

	//Put closing tag in file
	n3, err = f.WriteString("</svg>")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r
	return result
}

//!-
