package ui

import (
	"fmt"
	"github.com/Nyralei/print-pdf/internal/config"
	"github.com/Nyralei/print-pdf/internal/pdf"
	"github.com/andlabs/ui"
	"github.com/skratchdot/open-golang/open"
)

func SetupUI(cfg *config.Config, loc *Localization) {
	mainWindow := ui.NewWindow("PDF Creator", 400, 200, false)
	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	box := ui.NewVerticalBox()
	box.SetPadded(true)

	imageBox := ui.NewHorizontalBox()
	imageBox.SetPadded(true)
	imageLabel := ui.NewLabel(loc.Translate("image_path"))
	imageEntry := ui.NewEntry()
	imageEntry.SetText(cfg.ImagePath)
	imageButton := ui.NewButton(loc.Translate("choose_image"))
	imageButton.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainWindow)
		if filename != "" {
			cfg.ImagePath = filename
			imageEntry.SetText(filename)
		}
	})
	imageBox.Append(imageLabel, false)
	imageBox.Append(imageEntry, true)
	imageBox.Append(imageButton, false)
	box.Append(imageBox, false)

	pdfBox := ui.NewHorizontalBox()
	pdfBox.SetPadded(true)
	pdfLabel := ui.NewLabel(loc.Translate("pdf_path"))
	pdfEntry := ui.NewEntry()
	pdfEntry.SetText(cfg.PDFPath)
	pdfButton := ui.NewButton(loc.Translate("choose_pdf"))
	pdfButton.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainWindow)
		if filename != "" {
			cfg.PDFPath = filename
			pdfEntry.SetText(filename)
		}
	})
	pdfBox.Append(pdfLabel, false)
	pdfBox.Append(pdfEntry, true)
	pdfBox.Append(pdfButton, false)
	box.Append(pdfBox, false)

	optionsBox := ui.NewHorizontalBox()
	optionsBox.SetPadded(true)
	openCheckbox := ui.NewCheckbox(loc.Translate("open_in_browser"))
	openCheckbox.SetChecked(cfg.OpenInBrowser)
	openCheckbox.OnToggled(func(*ui.Checkbox) {
		cfg.OpenInBrowser = openCheckbox.Checked()
	})
	landscapeCheckbox := ui.NewCheckbox(loc.Translate("landscape_orientation"))
	landscapeCheckbox.SetChecked(cfg.Landscape)
	landscapeCheckbox.OnToggled(func(*ui.Checkbox) {
		cfg.Landscape = landscapeCheckbox.Checked()
	})
	optionsBox.Append(openCheckbox, false)
	optionsBox.Append(landscapeCheckbox, false)
	box.Append(optionsBox, false)

	createButton := ui.NewButton(loc.Translate("create_pdf"))
	createButton.OnClicked(func(*ui.Button) {
		if cfg.ImagePath == "" || cfg.PDFPath == "" {
			ui.MsgBoxError(mainWindow, loc.Translate("error"), loc.Translate("select_image_and_pdf"))
			return
		}
		err := pdf.CreatePDF(cfg.ImagePath, cfg.PDFPath, cfg.Landscape)
		if err != nil {
			ui.MsgBoxError(mainWindow, loc.Translate("error"), fmt.Sprintf(loc.Translate("error_creating_pdf"), err))
			return
		}
		if cfg.OpenInBrowser {
			err = open.Start(cfg.PDFPath)
			if err != nil {
				ui.MsgBoxError(mainWindow, loc.Translate("error"), fmt.Sprintf(loc.Translate("error_opening_pdf"), err))
			}
		}
		ui.MsgBox(mainWindow, loc.Translate("success"), loc.Translate("pdf_created"))
	})

	box.Append(createButton, false)

	mainWindow.SetChild(box)
	mainWindow.Show()
}
