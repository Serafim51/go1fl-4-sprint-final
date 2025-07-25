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
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	errInvalidFormat      = errors.New("invalid data format")
	errNumberParsing      = errors.New("failed to parse step count")
	errTimeParsing        = errors.New("failed to parse time value")
	errZeroSteps          = errors.New("step count must be greater than 0")
	errDuration           = errors.New("invalid format or duration <= 0")
	errInvalidInput       = errors.New("invalid input data")
	errInvalidTriningType = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	units := strings.Split(data, ",") // разделение строки на слайс строк
	if len(units) != 3 {              // проверка длины слайса
		return 0, "", 0, fmt.Errorf("%w: expected format - 'number,string,duration'", errInvalidFormat) //проверка соответствия формату
	}
	number, err := strconv.Atoi(units[0])
	if err != nil { //проверка преобразования строки в число (количество шагов)
		return 0, "", 0, fmt.Errorf("%w: %v", errNumberParsing, err)
	} else if number <= 0 { //проверка шагов на > 0
		return 0, "", 0, fmt.Errorf("%w: %v", errZeroSteps, err)
	}
	time, err := time.ParseDuration(units[2])
	if err != nil { //проверка формата времени
		return 0, "", 0, fmt.Errorf("%w: %v", errTimeParsing, err)
	}
	if time <= 0 { //проверка формата времени
		return 0, "", 0, fmt.Errorf("%w: %v", errDuration, err)
	}
	return number, units[1], time, nil //units[1] - тип тренеровки
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {

		return 0
	}

	return distance(steps, height) / duration.Hours()
}

// требует решения, пауза
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, trainingType, time, err := parseTraining(data)
	if err != nil {
		log.Println(err) // При ошибке логируем и прерываем выполнение, чтобы не обрабатывать некорректные данные
		return "", err
	}

	switch trainingType {
	case "Бег":
		runCal, err := RunningSpentCalories(steps, weight, height, time)
		if err != nil {
			log.Println(err) // При ошибке логируем и прерываем выполнение, чтобы не обрабатывать некорректные данные
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, time.Hours(), distance(steps, height), meanSpeed(steps, height, time), runCal), nil

	case "Ходьба":
		walkCal, err := WalkingSpentCalories(steps, weight, height, time)
		if err != nil {
			log.Println(err) // При ошибке логируем и прерываем выполнение, чтобы не обрабатывать некорректные данные
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, time.Hours(), distance(steps, height), meanSpeed(steps, height, time), walkCal), nil
	default:
		return "", fmt.Errorf("%w: ", errInvalidTriningType)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//реализация возможных ошибок
	if steps <= 0 {
		return 0, fmt.Errorf("%w: step count must be greater than 0", errInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: weight must be greater than 0", errInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: height must be greater than 0", errInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: training duration must be greater than 0", errInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * meanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("%w: step count must be greater than 0", errInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: weight must be greater than 0", errInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: height must be greater than 0", errInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: training duration must be greater than 0", errInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * meanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil

}
