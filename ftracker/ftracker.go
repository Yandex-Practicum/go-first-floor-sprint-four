package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMSec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// Константы для расчета калорий, расходуемых при плавании.
const (
	//swimmingLenStep                  = 1.38 // длина одного гребка.
	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// distance float64 — дистанция(в километрах), которую преодолел пользователь за время тренировки.
// duration float64 — длительность тренировки в часах.
func meanSpeed(distance, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance / duration
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
// speed float64 — средняя скорость при беге.
func runningSpentCalories(weight, duration, speed float64) float64 {
	return (runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// speed float64 — средняя скорость при ходьбе.
func walkingSpentCalories(duration, weight, height, speed float64) float64 {
	speedInMSec := speed * kmhInMSec // переводим скорость из км\ч в м\сек
	heightInM := height / 100        // переводим рост из сантиметров в метры
	return ((walkingCaloriesWeightMultiplier * weight) + ((math.Pow(speedInMSec, 2) / heightInM) * walkingSpeedHeightMultiplier * weight)) * duration * minInH
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// speed float64 — средняя скорость при плавании.
func swimmingSpentCalories(duration, weight, speed float64) float64 {
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func showTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	// ваш код здесь
	switch {
	case trainingType == "Бег":
		distance := distance(action)
		speed := meanSpeed(distance, duration)
		calories := runningSpentCalories(weight, duration, speed)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := distance(action)
		speed := meanSpeed(distance, duration)
		calories := walkingSpentCalories(duration, weight, height, speed)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := distance(action)
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := swimmingSpentCalories(duration, weight, speed)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}
