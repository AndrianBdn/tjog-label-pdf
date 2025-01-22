package main

import (
	"fmt"
	"math"

	"github.com/jung-kurt/gofpdf"
)

const (
	// All units are millimeters
	cellWidth     = 51.0
	cellHeight    = 44.0
	dividerHeight = 22.0

	pageMarginX = 25.0
	pageMarginY = 25.0
)

type pageConfig struct {
	width  float64
	height float64
}

var pageSizes = map[string]pageConfig{
	"A4": {
		width:  210.0,
		height: 297.0,
	},
	"Letter": {
		width:  216.0,
		height: 279.0,
	},
}

func main() {
	for pageKey, config := range pageSizes {
		pdf := gofpdf.New("P", "mm", pageKey, "")
		pdf.AddPage()

		cols := int(math.Floor((config.width - 2*pageMarginX) / cellWidth))
		rows := int(math.Floor((config.height - 2*pageMarginY) / cellHeight))

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				x := pageMarginX + float64(col)*cellWidth
				y := pageMarginY + float64(row)*cellHeight

				pdf.SetDashPattern([]float64{0.5, 0.5}, 0)
				pdf.SetDrawColor(0, 0, 0)

				pdf.Line(x, y, x+cellWidth, y)
				pdf.Line(x, y, x, y+cellHeight)
				pdf.Line(x+cellWidth, y, x+cellWidth, y+cellHeight)
				pdf.Line(x, y+cellHeight, x+cellWidth, y+cellHeight)

				gray := 200.0 / 255.0
				pdf.SetDrawColor(int(gray*255), int(gray*255), int(gray*255))
				pdf.Line(x, y+dividerHeight, x+cellWidth, y+dividerHeight)
			}
		}

		// link
		pdf.SetFont("Helvetica", "", 8)
		pdf.SetTextColor(100, 100, 100)
		linkText := "https://github.com/AndrianBdn/tjog-label-pdf"
		x := pageMarginX
		y := config.height - pageMarginY
		pdf.Text(x, y, linkText)

		filename := fmt.Sprintf("tjog_labels_%s.pdf", pageKey)
		err := pdf.OutputFileAndClose(filename)
		if err != nil {
			panic(err)
		}
	}
}
