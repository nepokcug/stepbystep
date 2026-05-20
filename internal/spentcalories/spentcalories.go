package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделяем строку на слайс
	parts := strings.Split(data, ",")
	//Проверяем чтобы длина слайса была именно три
	if len(parts) != 3 {
		return 0, "", 0, errors.New("invalid format: expected 'steps,activity,duration'")
	}
	// Преобразуем шаги в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format: %w", err)
	}
	// Проверяем чтобы количество шагов было больше нуля
	if steps <= 0 {
		return 0, "", 0, errors.New("steps must be greater than 0")
	}
	// Преобразуем время в time
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format: %w", err)
	}
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := stepLengthCoefficient * height
	distanceMeters := stepLength * float64(steps)
	distanceKm := distanceMeters / float64(mInKm)
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dis := distance(steps, height)
	hours := duration.Hours()
	speed := dis / hours
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Получаем значения из строки данных
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err) // Выводим ошибку в лог
		return "", err
	}
	// Находим дистанцию и скорость
	distanceKm := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	// Определяем тип тренировки и формируем строку
	switch activity {
	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f",
			duration.Hours(), distanceKm, speed, calories), nil

	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f",
			duration.Hours(), distanceKm, speed, calories), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка входных параметров на корректность
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}
	// Находим среднюю скорость
	meanSpeedValue := meanSpeed(steps, height, duration)
	// Находим время в минутах
	durationInMinutes := duration.Minutes()
	// Находим калории
	calories := (weight * meanSpeedValue * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка входных параметров на корректность
	if steps <= 0 {
		return 0, errors.New("steps must be greater than 0")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}
	// Находим среднюю скорость
	meanSpeedValue := meanSpeed(steps, height, duration)
	// Находим время в минутах
	durationInMinutes := duration.Minutes()
	// Находим калории
	calories := (weight * meanSpeedValue * durationInMinutes) / minInH
	reducedCalories := calories * walkingCaloriesCoefficient
	return reducedCalories, nil
}
