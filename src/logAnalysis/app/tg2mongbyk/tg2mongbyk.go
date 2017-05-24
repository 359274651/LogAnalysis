package main

import (
	"fmt"
	"log"
	//"sync"
	"github.com/streadway/amqp"
	//"github.com/Shopify/sarama"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"／
	"github.com/Shopify/sarama"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

//
//type Person struct {
//	Name  string
//	Phone string
//}
//bin/kafka-console-consumer.sh --zookeeper localhost:2181 --topic telegraf --from-beginning
//
//func main() {
//	session, err := mgo.Dial("mongodb.qiniu.io")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//
//	c := session.DB("test").C("people")
//	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
//		&Person{"Cla", "+55 53 8402 8510"})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	result := Person{}
//	err = c.Find(bson.M{"name": "Ale"}).One(&result)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Phone:", result.Phone)
//}
