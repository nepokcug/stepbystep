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

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделяем строку на слайс
	parts := strings.Split(data, ",")
	//Проверяем чтобы длина слайса была именно два
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid format: expected 'steps,duration'")
	}
	// Преобразуем шаги в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format: %w", err)
	}
	// Проверяем чтобы количество шагов было больше нуля
	if steps <= 0 {
		return 0, 0, errors.New("steps must be greater than 0")
	}
	// Преобразуем время в time
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be greater than 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// Получаем данные с помощью parsePackage
	steps, duration, err := parsePackage(data)
	// Если с parsePackage вернулась ошибка выводим ее и отправляем пустую строку
	if err != nil {
		log.Println("Ошибка:", err)
		return ""
	}
	// Проверяем чтобы количесвто шагов было больше нуля (з.ы. не понимаю зачем, это уже было в parsePackage)
	if steps <= 0 {
		log.Println("Ошибка, количество шагов должно быть больше 0")
		return ""
	}
	// Вычисляем км
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	// walkingSpentCalories еще возращает ошибку, выведем ее в случае возникновения
	if err != nil {
		fmt.Println("Ошибка в расчете каллорий при ходьбе:", err)
		return ""
	}
	// Формируем сроку
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)

	return result
}
