package daysteps

import (
	"errors"
	"fmt"
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
	ErrInvalidFormat = errors.New("ошибка формата данных")
	ErrNumberParsing = errors.New("ошибка получания числа шагов")
	ErrTimeParsing   = errors.New("ошибка получения значения времени")
	ErrZeroSteps     = errors.New("количество шагов <= 0")
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	units := strings.Split(data, ",") // разделение строки data на слайс строк
	if len(units) != 2 {              // проверка длины слайса
		return 0, 0, fmt.Errorf("%w: ожидаемый формат 'число,время'", ErrInvalidFormat)
	}
	number, err := strconv.Atoi(units[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %v", ErrNumberParsing, err)
	}
	if number <= 0 { //Проверка количества шагов
		return 0, 0, fmt.Errorf("%w: количество шагов должно быть больше 0", ErrZeroSteps)
	}
	time, err := time.ParseDuration(units[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %v", ErrTimeParsing, err)
	}
	return number, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkTIme, err := parsePackage(data)
	if err != nil { //определение ошибки
		fmt.Println(err)
		return ""
	}
	if steps <= 0 { //если не верное количество ошибок - вернуть пустую строку
		return ""
	}
	distant := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTIme)
	if err != nil { //определение ошибки
		fmt.Println(err)
		return ""
	}
	//определить количество калорий после реализации соответствующей функции -  в процессе

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %f км.\nВы сожгли %f ккал.", steps, distant, calories) //переменная calories будет объявлена позднее

}
