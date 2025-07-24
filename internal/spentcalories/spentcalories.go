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
	ErrInvalidFormat      = errors.New("ошибка формата данных")
	ErrNumberParsing      = errors.New("ошибка получания числа шагов")
	ErrTimeParsing        = errors.New("ошибка получения значения времени")
	ErrZeroSteps          = errors.New("количество шагов меньше или равно 0")
	ErrDuration           = errors.New("ошибка формата, либо продолжительность < 0")
	ErrUnknownType        = errors.New("тип переменной не соответствует")
	ErrInvalidInput       = errors.New("ошибка входных данных")
	ErrInvalidTriningType = errors.New("неизвестный тип тренировки")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	units := strings.Split(data, ",") // разделение строки на слайс строк
	if len(units) != 2 {              // проверка длины слайса
		return 0, "", 0, fmt.Errorf("%w: ожидаемый формат 'число,строка,время'", ErrInvalidFormat) //проверка соответствия формату
	}
	number, err := strconv.Atoi(units[0])
	if err != nil { //проверка преобразования строки в число (количество шагов)
		return 0, "", 0, fmt.Errorf("%w: %v", ErrNumberParsing, err)
	} else if number <= 0 { //проверка шагов на > 0
		return 0, "", 0, fmt.Errorf("%w: %v", ErrZeroSteps, err)
	}
	time, err := time.ParseDuration(units[2])
	if err != nil { //проверка формата времени
		return 0, "", 0, fmt.Errorf("%w: %v", ErrTimeParsing, err)
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
		log.Println(err) //логирование ошибки
	}
	runCal, err := RunningSpentCalories(steps, weight, height, time)
	if err != nil {
		log.Println(err) //логирование ошибки
	}
	walkCal, err := WalkingSpentCalories(steps, weight, height, time)
	if err != nil {
		log.Println(err) //логирование ошибки
	}
	switch trainingType {
	case "Бег":
		return fmt.Sprintf("Тип тренеровок: %s\nДлительность: %v ч.\n Дистанция: %.2f км.\nСкорость: %.2fкм/ч\nСожгли калорий: %.2f", trainingType, time, distance(steps, height), meanSpeed(steps, height, time), runCal), nil

	case "Шаг":
		return fmt.Sprintf("Тип тренеровок: %s\nДлительность: %v ч.\n Дистанция: %.2f км.\nСкорость: %.2fкм/ч\nСожгли калорий: %.2f", trainingType, time, distance(steps, height), meanSpeed(steps, height, time), walkCal), nil
	default:
		return "", fmt.Errorf("%w: ", ErrInvalidTriningType)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	//реализация возможных ошибок
	if steps <= 0 {
		return 0, fmt.Errorf("%w: число шагов должно быть больше 0", ErrInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: вес должен быть больше 0", ErrInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: рост должен быть больше 0", ErrInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: длительность тренеровки должна быть больше 0", ErrInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * meanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("%w: число шагов должно быть больше 0", ErrInvalidInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: вес должен быть больше 0", ErrInvalidInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: рост должен быть больше 0", ErrInvalidInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: длительность тренеровки должна быть больше 0", ErrInvalidInput)
	}
	// реализация функции ** для себя** - требует теста
	return weight * meanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil

}
