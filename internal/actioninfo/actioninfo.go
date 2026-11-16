package actioninfo

import (
	"fmt"
	"log"
)

// DataParser позволяет одинаково обрабатывать тренировки и прогулки.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Ошибка в одной записи не мешает обработать следующие записи.
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(info)
	}
}
