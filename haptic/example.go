// SPDX-License-Identifier: Unlicense OR MIT

package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
	"gioui.org/x/haptic"
)

var buzzer *haptic.Buzzer

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	btn := widget.Clickable{}
	buzzer = haptic.NewBuzzer(w)
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				if btn.Clicked() {
					buzzer.Buzz()
				}
				gtx := layout.NewContext(&ops, e)
				layout.Center.Layout(gtx, material.Button(th, &btn, "buzz").Layout)
				e.Frame(gtx.Ops)
			default:
				ProcessPlatformEvent(e)
			}
		case err := <-buzzer.Errors():
			if err != nil {
				log.Printf("buzzer error: %v", err)
			}
		}
	}
}
