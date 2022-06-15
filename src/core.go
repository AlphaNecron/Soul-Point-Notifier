package noti

import (
	"fmt"
	"fyne.io/fyne/v2/theme"
	"noti/src/assets"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ahmetb/go-linq/v3"
)

var refreshTicker *time.Ticker

var a fyne.App

func update(data *[]WorldData, table *widget.Table) {
	success, res := fetch()
	if success {
		*data = res.transform()
		table.Refresh()
	}
}

func setWidth(table *widget.Table, headers *widget.Table, ids []int, widths []int) {
	for i, id := range ids {
		table.SetColumnWidth(id, float32(widths[i]))
		headers.SetColumnWidth(id, float32(widths[i]))
	}
}

func Start() {
	a = app.NewWithID("dev.goldor.spnoti")
	w := a.NewWindow("Soul Point Notifier")
	initPrefs()
	var data []WorldData
	w.SetIcon(assets.Logo)
	fmt.Println(RefreshInterval)
	headers := []string{"World", "Uptime", "Total players", "Time until next soulpoint"}
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Unknown")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label, r := o.(*widget.Label), data[i.Row]
			switch i.Col {
			case 0:
				label.SetText(r.WorldName)
			case 1:
				label.SetText(parseUptime(r.Uptime))
			case 2:
				if r.PlayerCount >= 40 {
					label.SetText(fmt.Sprintf("%d (Full)", r.PlayerCount))
				} else {
					label.SetText(strconv.Itoa(r.PlayerCount))
				}
			default:
				label.SetText(fmt.Sprintf("%d minutes", r.NextSoulpoint))
			}
		})
	htable := widget.NewTable(
		func() (int, int) {
			return 1, 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Unknown")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(headers[i.Col])
		})
	setWidth(table, htable, []int{1, 2, 3}, []int{125, 125, 225})
	intervalBox := newTitledNumberBox("Refresh interval: ")
	intervalBox.NumberBox.Validator = func(s string) error {
		if len(s) <= 0 {
			return nil
		}
		RefreshInterval, err := toNumber(s, RefreshInterval)
		a.Preferences().SetInt("interval", RefreshInterval)
		refreshTicker.Reset(secs(RefreshInterval))
		if err {
			intervalBox.NumberBox.SetText(strconv.Itoa(RefreshInterval))
		}
		return nil
	}
	cdBox := newTitledNumberBox("Cooldown: ")
	cdBox.NumberBox.Validator = func(s string) error {
		if len(s) <= 0 {
			return nil
		}
		NotifyCooldown, err := toNumber(s, NotifyCooldown)
		notiTicker.Reset(secs(NotifyCooldown))
		a.Preferences().SetInt("cooldown", NotifyCooldown)
		if err {
			cdBox.NumberBox.SetText(strconv.Itoa(NotifyCooldown))
		}
		return nil
	}
	// containers
	opts := container.NewBorder(nil, nil, widget.NewButtonWithIcon("Fetch", theme.ViewRefreshIcon(), func() {
		update(&data, table)
	}), nil, container.NewBorder(nil, nil, intervalBox.Container, nil, cdBox.Container))
	header := container.NewVBox(opts, htable)
	ctn := container.NewBorder(header, nil, nil, nil, table)
	w.SetContent(ctn)
	w.Resize(fyne.NewSize(575, 400))
	// refresh
	go func() {
		refreshTicker = time.NewTicker(secs(RefreshInterval))
		for range refreshTicker.C {
			if !Paused {
				update(&data, table)
				w := linq.From(data).FirstWith(func(i interface{}) bool {
					return i.(WorldData).NextSoulpoint <= NotifyCooldown
				}).(WorldData)
				notify(w)
			}
		}
	}()
	w.Show()
	start, end := initTray(w.Hide, w.Show)
	w.SetOnClosed(end)
	start()
	initNotiCooldown()
	intervalBox.NumberBox.SetText(strconv.Itoa(RefreshInterval))
	cdBox.NumberBox.SetText(strconv.Itoa(NotifyCooldown))
	update(&data, table)
	a.Run()
}
