package service

import (
	"time"

	"github.com/pkg/errors"

	"awesomeHealth/healthtracker/internal/models"
)

func NewHealthService(
	activityRepository activityRepository,
	nutritionRepository nutritionRepository,
	sleepRepository sleepRepository,
) *healthService {
	return &healthService{
		activityRepository:  activityRepository,
		nutritionRepository: nutritionRepository,
		sleepRepository:     sleepRepository,
	}
}

type healthService struct {
	activityRepository  activityRepository
	nutritionRepository nutritionRepository
	sleepRepository     sleepRepository
}

func (s *healthService) AddActivity(activity string, durationInSeconds int, calories int) error {
	if durationInSeconds < 0 {
		return errors.New("duration is negative")
	}
	if calories < 0 {
		return errors.New("calories is negative")
	}
	duration := time.Second * time.Duration(durationInSeconds)
	return s.activityRepository.AddActivityInfo(activity, duration, calories)
}

func (s *healthService) AddNutrition(dish string, size int, calories int) error {
	if size < 0 {
		return errors.New("size is negative")
	}
	if calories < 0 {
		return errors.New("calories is negative")
	}
	return s.nutritionRepository.AddNutritionInfo(dish, size, calories)
}

func (s *healthService) AddSleep(durationInSeconds int) error {
	if durationInSeconds < 0 {
		return errors.New("duration is negative")
	}
	duration := time.Second * time.Duration(durationInSeconds)
	return s.sleepRepository.AddSleepInfo(duration)
}

func (s *healthService) GetStats() (models.Stats, error) {
	activityTime, err := s.activityRepository.GetSumDuration()
	if err != nil {
		return models.Stats{}, errors.Wrap(err, "get activity time")
	}
	sleepTime, err := s.sleepRepository.GetSumTime()
	if err != nil {
		return models.Stats{}, errors.Wrap(err, "get sleep time")
	}
	lostCalories, err := s.activityRepository.GetSumCalories()
	if err != nil {
		return models.Stats{}, errors.Wrap(err, "get lost calories")
	}
	gainedCalories, err := s.nutritionRepository.GetSumCalories()
	if err != nil {
		return models.Stats{}, errors.Wrap(err, "get gained calories")
	}
	return models.Stats{
		LostCalories:   lostCalories,
		GainedCalories: gainedCalories,
		ActivityTime:   int(activityTime.Seconds()),
		SleepTime:      int(sleepTime.Seconds()),
	}, nil
}
