package handlers

import (
	"server/models"
	"testing"
)

func profile(achievements, minions, mounts int, private bool) *models.LodestoneProfile {
	return &models.LodestoneProfile{
		Achievements:       make([]*models.AchievementInfo, achievements),
		AllAchievementInfo: &models.AllAchievementInfo{Private: private, TotalAchievements: uint32(achievements)},
		Minions:            make([]*models.Minion, minions),
		Mounts:             make([]*models.Mount, mounts),
	}
}

func TestKeepCollectedEntities(t *testing.T) {
	tests := []struct {
		name                     string
		fresh, stored            *models.LodestoneProfile
		achievements, min, mount int
	}{
		{"truncated scrape keeps stored", profile(100, 0, 0, false), profile(1698, 400, 300, false), 1698, 400, 300},
		{"progress is taken", profile(1700, 401, 301, false), profile(1698, 400, 300, false), 1700, 401, 301},
		{"made private drops achievements", profile(0, 401, 301, true), profile(1698, 400, 300, false), 0, 401, 301},
		{"no stored data", profile(5, 1, 1, false), nil, 5, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepCollectedEntities(tt.fresh, tt.stored)
			if got := len(tt.fresh.Achievements); got != tt.achievements {
				t.Errorf("achievements = %d, want %d", got, tt.achievements)
			}
			if got := int(tt.fresh.AllAchievementInfo.TotalAchievements); got != tt.achievements {
				t.Errorf("total achievements = %d, want %d", got, tt.achievements)
			}
			if got := len(tt.fresh.Minions); got != tt.min {
				t.Errorf("minions = %d, want %d", got, tt.min)
			}
			if got := len(tt.fresh.Mounts); got != tt.mount {
				t.Errorf("mounts = %d, want %d", got, tt.mount)
			}
		})
	}
}
