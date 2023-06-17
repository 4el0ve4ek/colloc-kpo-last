package service

import (
	"time"
)

type activityRepository interface {
	AddActivityInfo(activity string, duration time.Duration, calories int) error
	GetSumDuration() (time.Duration, error)
	GetSumCalories() (int, error)
}

type nutritionRepository interface {
	AddNutritionInfo(dish string, size int, calories int) error
	GetSumCalories() (int, error)
}

type sleepRepository interface {
	AddSleepInfo(duration time.Duration) error
	GetSumTime() (time.Duration, error)
}
