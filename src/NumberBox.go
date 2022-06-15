package noti

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

type (
	NumberBox struct {
		widget.Entry
	}
	TitledNumberBox struct {
		Container *fyne.Container
		NumberBox *NumberBox
	}
)

func newTitledNumberBox(title string) *TitledNumberBox {
	box := newNumberBox()
	label := canvas.NewText(title, color.White)
	label.TextStyle = fyne.TextStyle{Bold: true}
	return &TitledNumberBox{Container: container.NewBorder(nil, nil, label, nil, box), NumberBox: box}
}

func newNumberBox() *NumberBox {
	entry := &NumberBox{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *NumberBox) TypedRune(r rune) {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		e.Entry.TypedRune(r)
	}
}

func (e *NumberBox) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseInt(content, 0, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}
