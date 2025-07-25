package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

var (
	errInvalidFormat = errors.New("data format error")
	errNumberParsing = errors.New("step count retrieval error")
	errTimeParsing   = errors.New("time value retrieval error")
	errZeroSteps     = errors.New("step count <= 0")
	errTimeZero      = errors.New("time value <= 0")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	units := strings.Split(data, ",") // разделение строки data на слайс строк
	if len(units) != 2 {              // проверка длины слайса
		return 0, 0, fmt.Errorf("%w: expected format - 'number,duration'", errInvalidFormat)
	}
	number, err := strconv.Atoi(units[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %v", errNumberParsing, err)
	}
	if number <= 0 { //Проверка количества шагов
		return 0, 0, fmt.Errorf("%w: step count must be greater than 0", errZeroSteps)
	}
	time, err := time.ParseDuration(units[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %v", errTimeParsing, err)
	}
	if time <= 0 {
		return 0, 0, fmt.Errorf("%w: %v", errTimeZero, err)
	}
	return number, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkTIme, err := parsePackage(data)
	if err != nil { //определение ошибки
		log.Println(err)
		return ""
	}
	if steps <= 0 { //если не верное количество ошибок - вернуть пустую строку
		return ""
	}
	distant := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTIme)
	if err != nil { //определение ошибки
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distant, calories)

}
