package noti

var Paused bool

var RefreshInterval int

var NotifyCooldown int

var Hidden bool

func initPrefs() {
	RefreshInterval = a.Preferences().IntWithFallback("interval", 5)
	NotifyCooldown = a.Preferences().IntWithFallback("cooldown", 15)
}
