package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

func main() {

	connection, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer connection.Close()

	log.Println("Connected to AMQP")

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = time.Second * 1
	var connection *amqp.Connection

	for {
		dial, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Println("Failed to connect to RabbitMQ:", err)
			counts++
		} else {
			connection = dial
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
