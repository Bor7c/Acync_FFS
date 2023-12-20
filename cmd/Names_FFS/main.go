package main

import (
	"ASYNC_FFS/internal/api"
	"log"
)

func main() {
	log.Println("App start")
	api.StartServer()
	log.Println("App stop")
}

// func genRandomNumber() int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(851)
// }

// func genRandomName() {
// 	time.Sleep(3 * time.Second)
// 	fio := [15]string{
// 		"Иванов Иван Иванович",
// 		"Петров Петр Петрович",
// 		"Сидоров Сидор Сидорович",
// 		"Ильин Илья Ильич",
// 		"Янов Ян Янович",
// 		"Прохоров Прохор Прохорович",
// 		"Кузьмин Кузьма Кузьмич",
// 		"Абрамов Артем Абрамович",
// 		"Булгаков Иван Иванович",
// 		"Сазонов Дмитрий Сазонов",
// 		"Петров Иван Петров",
// 		"Сидорова Анна Сидорова",
// 		"Дубовский Иван Дубовский",
// 		"Иванченко Александр Иванченко",
// 		"Фомина Мария Фомина",
// 	}

// 	rand.Seed(time.Now().UnixNano())   // Устанавливаем seed для генератора случайных чисел
// 	randomIndex := rand.Intn(len(fio)) // Генерируем случайный индекс в пределах длины массива
// 	randomFIO := fio[randomIndex]      // Выбираем случайную запись из массива

// 	fmt.Println(randomFIO)
// }
