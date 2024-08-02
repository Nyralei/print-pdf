package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func CreatePDF(imagePath, pdfPath string, landscape bool) error {
	orientation := "P"
	if landscape {
		orientation = "L"
	}

	pdf := gofpdf.New(orientation, "mm", "A4", "")
	pdf.AddPage()

	imageWidth, imageHeight := getImageDimensions(landscape)

	pdf.Image(imagePath, 0, 0, imageWidth, imageHeight, false, "", 0, "")

	err := pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		return fmt.Errorf("error saving PDF: %v", err)
	}

	return nil
}

func getImageDimensions(landscape bool) (float64, float64) {
	// A4 dimensions
	if landscape {
		return 297.0, 210.0
	}
	return 210.0, 297.0
}
