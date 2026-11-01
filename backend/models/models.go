package models

import (
	"server/models/gcrank"
	"server/models/gender"
	"time"
)

type DiscordAuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

type User struct {
	DiscordUser *DiscordUser          `json:"discordUser"`
	Characters  map[string]*Character `json:"characters"`
	Settings    *Settings             `json:"settings"`
}

type Settings struct {
	AuthorizationCode   string `json:"authorizationCode"`
	CharacterClaimCode  string `json:"characterClaimCode"`
	LatestNewsSeen      int    `json:"latestNewsSeen"`
	LatestCountdownSeen int    `json:"latestCountdownSeen"`
	SpoilersOption      int    `json:"spoilersOption"`
	PatchNumsOption     int    `json:"patchNumsOption"`
	ThemeOption         int    `json:"themeOption"`
}

type Character struct {
	Updated       int64             `json:"updated"`
	LodestoneData *LodestoneProfile `json:"lodestoneData"`
	ChecklistData *ChecklistData    `json:"checklistData"`
	EncounterIDs  []*int            `json:"encounterIDs"`
}

type LodestoneProfile struct {
	Achievements       []*AchievementInfo
	AllAchievementInfo *AllAchievementInfo
	Character          *CharacterData
	Minions            []*Minion
	Mounts             []*Mount
}

type CharacterData struct {
	ActiveClassJob    *ClassJob
	Avatar            string
	Bio               string
	ClassJobs         []*ClassJob
	ClassJobBozjan    *ClassJobBozja
	ClassJobElemental *ClassJobEureka
	DC                string
	FreeCompanyID     string
	FreeCompanyName   string
	//GearSet           *GearSet
	Gender           gender.Gender
	GrandCompanyInfo *GrandCompanyInfo
	GuardianDeity    *IconedNamedEntity
	ID               uint32
	Name             string
	Nameday          string
	ParseDate        time.Time
	Portrait         string
	PvPTeamID        string
	Race             *GenderedEntity
	Title            *Title
	Town             *IconedNamedEntity
	Tribe            *GenderedEntity
	World            string
}

type ClassJob struct {
	ClassID       uint8
	ExpLevel      uint32
	ExpLevelMax   uint32
	ExpLevelTogo  uint32
	IsSpecialized bool
	JobID         uint8
	Level         uint8
	Name          string
	UnlockedState *ClassJobUnlockedState
}

// Lodestone entities are embedded by value: godestone omits the fields of entities it cannot resolve (e.g. content
// newer than its database), which would leave an embedded pointer nil and make the Firestore encoder panic
type AchievementInfo struct {
	IDdEntity // NamedEntity
	Date      time.Time
}

type AllAchievementInfo struct {
	Private                bool
	TotalAchievements      uint32
	TotalAchievementPoints uint32
}

type ClassJobBozja struct {
	Level  uint8
	Mettle uint32
	//mettleRaw *colly.HTMLElement // TODO: https://github.com/xivapi/godestone/issues/17
	Name string
}

type ClassJobEureka struct {
	ExpLevel     uint32
	ExpLevelMax  uint32
	ExpLevelTogo uint32
	Level        uint8
	Name         string
}

type ClassJobUnlockedState struct {
	ID   uint8
	Name string
}

type GrandCompanyInfo struct {
	GrandCompany *NamedEntity
	RankID       gcrank.GCRank
}

type Title struct {
	TitleInternal
}

type Minion struct {
	IDdEntity // IconedNamedEntity
}

type Mount struct {
	IDdEntity // IconedNamedEntity
}

type TitleInternal struct {
	GenderedEntity
	Prefix bool
}

type ChecklistData struct {
	WeeklyChecklist []*ChecklistItem `json:"weeklyChecklist"`
	DailyChecklist  []*ChecklistItem `json:"dailyChecklist"`
	AdhocChecklist  []*ChecklistItem `json:"adhocChecklist"`
	NextWeeklyReset int              `json:"nextWeeklyReset"`
	NextDailyReset  int              `json:"nextDailyReset"`
}

type ChecklistItem struct {
	Name    string `json:"name"`
	Custom  bool   `json:"custom"`
	Checked bool   `json:"checked"`
	Hidden  bool   `json:"hidden"`
}
