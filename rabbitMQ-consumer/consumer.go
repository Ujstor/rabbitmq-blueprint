package main


import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
    RabbitHost     = os.Getenv("RABBIT_HOST")
    RabbitPort     = os.Getenv("RABBIT_PORT")
    RabbitUser     = os.Getenv("RABBIT_USERNAME")
    RabbitPassword = os.Getenv("RABBIT_PASSWORD")
)

func main() {
   
    consume()
}

func consume() {
    conn, err := amqp.Dial("amqp://" + RabbitUser + ":" + RabbitPassword + "@" + RabbitHost + ":" + RabbitPort + "/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"publisher", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	fmt.Println("Channel and Queue established")

	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			
			d.Ack(false)
		}
	  }()
	  
	  fmt.Println("Running...")
	  <-forever
}