package rabbitmq

import (
    "os"
    "github.com/streadway/amqp"

	l "rabbitmq-blueprint/internal/logger"

	_ "github.com/joho/godotenv/autoload"
)

var (
    RabbitHost     = os.Getenv("RABBIT_HOST")
    RabbitPort     = os.Getenv("RABBIT_PORT")
    RabbitUser     = os.Getenv("RABBIT_USERNAME")
    RabbitPassword = os.Getenv("RABBIT_PASSWORD")
)

func connectRabbitMQ() (*amqp.Channel, error) {
    conn, err := amqp.Dial("amqp://" + RabbitUser + ":" + RabbitPassword + "@" + RabbitHost + ":" + RabbitPort + "/")
    if err != nil {
       l.Log.Errorf("failed to connect to RabbitMQ: %v", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        l.Log.Errorf("failed to open a channel: %v", err)

    }

    return ch, nil
}

func SubmitMessage(message string) error {
    ch, err := connectRabbitMQ()
    if err != nil {
        l.Log.Errorf("failed to connect to RabbitMQ: %v", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "publisher", // name
        true,        // durable
        false,       // delete when unused
        false,       // exclusive
        false,       // no-wait
        nil,         // arguments
    )
    if err != nil {
        l.Log.Errorf("failed to declare a queue: %v", err)
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        l.Log.Errorf("failed to publish a message: %v", err)
    }

    l.Log.Info("Publish success!")
    return nil
}