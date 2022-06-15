package noti

import (
	"math"
	"sort"
	"strconv"
	"time"
)

func (w World) calcNextSoulpoint() int {
	return int(20 - math.Mod(math.Floor(math.Mod(float64((time.Now().UnixMilli()-w.Uptime)/1000), 3600)/60), 20))
}

func (w WorldList) transform() []WorldData {
	worlds := make([]WorldData, 0, len(w.Worlds))
	for k, e := range w.Worlds {
		worlds = append(worlds, WorldData{k, e.Uptime, len(e.Players), e.calcNextSoulpoint()})
	}
	sort.SliceStable(worlds, func(i, j int) bool {
		return worlds[i].NextSoulpoint < worlds[j].NextSoulpoint
	})
	return worlds
}

func parseUptime(t int64) string {
	return time.Since(time.UnixMilli(t)).Round(time.Second).String()
}

func secs(n int) time.Duration {
	return time.Duration(n) * time.Second
}

func toNumber(s string, fb int) (int, bool) {
	if res, err := strconv.Atoi(s); err == nil {
		if res <= 0 {
			return fb, true
		}
		return res, false
	}
	return fb, true
}
