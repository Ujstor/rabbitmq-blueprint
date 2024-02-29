#!/bin/sh

cat > consumer.go <<EOF
package main

import (
    "log"
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
        true,        // durable
        false,       // delete when unused
        false,       // exclusive
        false,       // no-wait
        nil,         // arguments
    )

    if err != nil {
        log.Fatalf("%s: %s", "Failed to declare a queue", err)
    }

    log.Println("Channel and Queue established")

    defer conn.Close()
    defer ch.Close()

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
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

    log.Println("Running...")
    <-forever
}
EOF

go mod init consumer 
go mod tidy