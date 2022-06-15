package noti

type (
	WorldList struct {
		Worlds map[string]World `json:"servers"`
	}
	World struct {
		Uptime  int64    `json:"firstSeen"`
		Players []string `json:"players"`
	}
	WorldData struct {
		WorldName     string
		Uptime        int64
		PlayerCount   int
		NextSoulpoint int
	}
)
