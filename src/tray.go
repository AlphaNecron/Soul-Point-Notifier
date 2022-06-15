package noti

import (
	"fyne.io/systray"
	"noti/src/assets"
)

func initTray(hide, show func()) (start, end func()) {
	return systray.RunWithExternalLoop(
		func() {
			systray.SetIcon(assets.Logo.Content())
			systray.SetTooltip("Soul Point Notifier")
			h := systray.AddMenuItem("Hide", "Hide the app window.")
			go func() {
				for !h.Disabled() {
					<-h.ClickedCh
					if Hidden {
						show()
						Hidden = false
						h.SetTitle("Hide")
						h.SetTooltip("Hide the app window.")
					} else {
						hide()
						Hidden = true
						h.SetTitle("Show")
						h.SetTooltip("Show the app window.")
					}
				}
			}()
			item := systray.AddMenuItemCheckbox("Stop notifying", "Stop the application from fetching and notifying whenever there is a soul point.", Paused)
			go func() {
				for !item.Disabled() {
					<-item.ClickedCh
					if Paused {
						Paused = false
						item.Uncheck()
					} else {
						Paused = true
						item.Check()
					}
				}
			}()
			systray.AddSeparator()
			q := systray.AddMenuItem("Quit", "Quit the whole app")
			go func() {
				<-q.ClickedCh
				systray.Quit()
			}()
		}, nil)
}
