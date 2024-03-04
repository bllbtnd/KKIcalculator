package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type Subject struct {
	Grade  int
	Weight int
}

func main() {
	myApp := app.New()
	iconResource, _ := fyne.LoadResourceFromPath("icon.png")
	myApp.SetIcon(iconResource)

	w := myApp.NewWindow("Grades")
	getGrades(myApp, w, displayResult)
	w.ShowAndRun()
}

func displayResult(myApp fyne.App, w fyne.Window, sbjList []Subject) {

	fmt.Printf("here")
	resultWindow := myApp.NewWindow("Results")

	kkiLabel := widget.NewLabel("The KKI is: " + strconv.FormatFloat(calcKKI(sbjList), 'f', 2, 64))
	avgLabel := widget.NewLabel("The weighted average is: " + strconv.FormatFloat(calcAvg(sbjList), 'f', 2, 64))

	content := container.NewVBox(
		kkiLabel,
		avgLabel,
	)

	resultWindow.SetContent(content)
	resultWindow.Show()
}

func getGrades(myApp fyne.App, w fyne.Window, callback func(fyne.App, fyne.Window, []Subject)) {
	var sbjList []Subject
	grade := widget.NewEntry()
	weight := widget.NewEntry()
	grdLabel := widget.NewLabel("Grade:")
	wegLabel := widget.NewLabel("Credit:")
	button := widget.NewButton("Add", func() {
		if grade.Text != "" && weight.Text != "" && sTOiInput(grade.Text) > 0 && sTOiInput(grade.Text) <= 5 && sTOiInput(weight.Text) > 0 {
			sbjList = append(sbjList, Subject{Grade: sTOiInput(grade.Text), Weight: sTOiInput(weight.Text)})
		}
		grade.SetText("")
		weight.SetText("")
	})
	buttond := widget.NewButton("Done", func() {
		if grade.Text != "" && weight.Text != "" && sTOiInput(grade.Text) > 0 && sTOiInput(grade.Text) <= 5 && sTOiInput(weight.Text) > 0 {
			sbjList = append(sbjList, Subject{Grade: sTOiInput(grade.Text), Weight: sTOiInput(weight.Text)})
		} else if grade.Text == "" && weight.Text == "" {
			w.Close()
			callback(myApp, w, sbjList)
		}
		grade.SetText("")
		weight.SetText("")
	})

	content := container.NewGridWithColumns(2,
		grdLabel, grade,
		wegLabel, weight,
		button, buttond,
	)
	w.SetContent(content)
	w.Resize(fyne.Size{200, 50})
}

func sTOiInput(ss string) int {
	num, err := strconv.Atoi(ss)
	if err != nil {
		fmt.Println("Your input is not an integer.")
		return 0
	}
	return num
}

func calcKKI(values []Subject) float64 {
	vallKerd := 0
	teljKredit := 0
	var teljSubj []Subject
	for i := 0; i < len(values); i++ {
		if values[i].Grade != 1 {
			teljKredit += values[i].Weight
			teljSubj = append(teljSubj, Subject{Grade: values[i].Grade, Weight: values[i].Weight})
		}
		vallKerd += values[i].Weight
	}
	astr := 0.0
	for i := 0; i < len(teljSubj); i++ {
		astr += float64(teljSubj[i].Grade * teljSubj[i].Weight)
	}
	astr = float64(astr) / 30
	return astr * (float64(teljKredit) / float64(vallKerd))
}

func calcAvg(values []Subject) float64 {
	var avg float64
	cre := 0.0
	all := 0.0
	for i := 0; i < len(values); i++ {
		if values[i].Grade != 1 {
			cre += float64(values[i].Weight)
			all += float64(values[i].Grade * values[i].Weight)
		}
	}
	avg = all / cre
	return avg
}
