package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const password = "FFS_IS_MY_APP"

func StartServer() {
	log.Println("Server start up")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/name", func(c *gin.Context) {
		var data NameData

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		BreachID := data.BreachID

		// Запуск горутины для отправки статуса
		go sendName(BreachID, password, fmt.Sprintf("http://localhost:8000/breaches/%d/name/", BreachID))

		c.JSON(http.StatusOK, gin.H{"message": "Status update initiated"})
	})
	router.Run(":5000")

	log.Println("Server down")
}

func genRandomName(password string) Result {
	time.Sleep(10 * time.Second)
	fio := [15]string{
		"Иванов Иван Иванович",
		"Петров Петр Петрович",
		"Сидоров Сидор Сидорович",
		"Ильин Илья Ильич",
		"Янов Ян Янович",
		"Прохоров Прохор Прохорович",
		"Кузьмин Кузьма Кузьмич",
		"Абрамов Артем Абрамович",
		"Булгаков Иван Иванович",
		"Сазонов Дмитрий Сазонов",
		"Петров Иван Петров",
		"Сидорова Анна Сидорова",
		"Дубовский Иван Дубовский",
		"Иванченко Александр Иванченко",
		"Фомина Мария Фомина",
	}

	rand.Seed(time.Now().UnixNano())   // Устанавливаем seed для генератора случайных чисел
	randomIndex := rand.Intn(len(fio)) // Генерируем случайный индекс в пределах длины массива
	randomFIO := fio[randomIndex]      // Выбираем случайную запись из массива

	fmt.Println(randomFIO)

	return Result{randomFIO, password}
}

// Функция для отправки статуса в отдельной горутине
func sendName(BreachID int, password string, url string) {
	// Выполнение расчётов с randomStatus
	result := genRandomName(password)

	// Отправка PUT-запроса к основному серверу
	_, err := performPUTRequest(url, result)
	if err != nil {
		fmt.Println("Error sending Name:", err)
		return
	}

	fmt.Println("Name sent successfully for BreachID:", BreachID)
}

type Result struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type NameData struct {
	BreachID int `json:"breach_id"`
}

func performPUTRequest(url string, data Result) (*http.Response, error) {
	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Создание PUT-запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}
