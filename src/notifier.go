package noti

import (
	"fmt"
	"fyne.io/fyne/v2"
	"time"
)

var notiTicker *time.Ticker

var onCooldown bool

func initNotiCooldown() {
	notiTicker = time.NewTicker(secs(NotifyCooldown))
	go func() {
		for range notiTicker.C {
			onCooldown = false
		}
	}()
}

func notify(w WorldData) {
	if onCooldown {
		return
	}
	a.SendNotification(fyne.NewNotification("Soul point available!", fmt.Sprintf("A soul point is going to be available in %d minutes in %s.", w.NextSoulpoint, w.WorldName)))
	notiTicker.Reset(secs(NotifyCooldown))
	onCooldown = true
}
