package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/VictorManoelCostaDeBarros/go-intensivo/internal/infra/database"
	"github.com/VictorManoelCostaDeBarros/go-intensivo/internal/usecase"
	"github.com/VictorManoelCostaDeBarros/go-intensivo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

// Função
func some(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera tudo rodar e dps executa o close
	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput

		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)

		fmt.Println("Mensagem processada e salva no banco: ", output)
	}
}
