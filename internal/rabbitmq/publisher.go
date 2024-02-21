package publisher

import (
    "fmt"
    "os"
    "github.com/streadway/amqp"
	_ "github.com/joho/godotenv/autoload"
)

var (
    RabbitHost     = os.Getenv("RABBIT_HOST")
    RabbitPort     = os.Getenv("RABBIT_PORT")
    RabbitUser     = os.Getenv("RABBIT_USERNAME")
    RabbitPassword = os.Getenv("RABBIT_PASSWORD")
)

func SubmitMessage(message string) error {
    conn, err := amqp.Dial("amqp://" + RabbitUser + ":" + RabbitPassword + "@" + RabbitHost + ":" + RabbitPort + "/")
    if err != nil {
        return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        return fmt.Errorf("failed to open a channel: %v", err)
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
        return fmt.Errorf("failed to declare a queue: %v", err)
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
        return fmt.Errorf("failed to publish a message: %v", err)
    }

    fmt.Println("Publish success!")
    return nil
}
