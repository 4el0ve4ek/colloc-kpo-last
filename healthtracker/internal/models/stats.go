package models

type Stats struct {
	LostCalories   int `json:"lost_calories"`
	GainedCalories int `json:"gained_calories"`
	ActivityTime   int `json:"activity_time_seconds"`
	SleepTime      int `json:"sleep_time_seconds"`
}
