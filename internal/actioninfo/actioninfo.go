package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
	// TODO: добавить методы
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// 1. Парсим строку данных
		err := dp.Parse(data)
		if err != nil {
			log.Printf("ошибка парсинга данных: %v", err)
			continue // Переходим к следующей итерации
		}

		// 2. Формируем строку с информацией
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка получения информации: %v", err)
			continue
		}

		// 3. Выводим результат
		fmt.Println(info)
	}
}

// TODO: реализовать функцию
