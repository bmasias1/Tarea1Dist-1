package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"

	"os"
	"time"

	"github.com/streadway/amqp"
)

func Abrir(port string, usuario string) {

	fmt.Println("Escuchando al " + usuario + " en el puerto " + port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Hubo un fallo al abrir el servidor: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Hubo un fallo al abrir el servidor gRPC: %s", err)
	}
}

//Funcion para arrojar errores
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Funcion encargada de leer registro 1
func readRegistroR1() [][]string {

	f, err := os.Open("../archivos/registroRetail1.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("No se puede leer el CSV", err.Error())
	}

	return rows
}

//Funcion encargada de leer registro 2
func readRegistroR2() [][]string {

	f, err := os.Open("../archivos/registroRetail2.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("No se puede leer el CSV", err.Error())
	}

	return rows
}

//Funcion encargada de leer registro del camion Normal
func readRegistroN() [][]string {

	f, err := os.Open("../archivos/registroNormal.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("No se puede leer el CSV", err.Error())
	}

	return rows
}

//Funcion encargada de enviar informacion del camion retail1 a finanzas
func senderToFinanzasR1(segundos int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	b := 0
	for {
		rows := readRegistroR1()
		for i := range rows {
			time.Sleep(time.Duration(segundos) * time.Second)
			if b != len(rows) {
				entregado := rows[i][6]
				if entregado == "0" {
					entregado = "No entregado"
				} else {
					entregado = "Entregado"
				}
				a := (rows[i][0]) + "," + entregado + "," + (rows[i][5]) + "," + (rows[i][2])

				body := a
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				log.Printf(" [x] Sent %s", body)
				failOnError(err, "Failed to publish a message")
				b++

			}
		}
		if b == len(rows) {
			println("No hay nuevas entregas del camion Retail1.")
		}

	}

}

//Funcion encargada de enviar informacion del camion retail2 a finanzas
func senderToFinanzasR2(segundos int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	b := 0
	for {
		rows := readRegistroR2()
		for i := range rows {
			time.Sleep(time.Duration(segundos) * time.Second)
			if b != len(rows) {
				entregado := rows[i][6]
				if entregado == "0" {
					entregado = "No entregado"
				} else {
					entregado = "Entregado"
				}
				a := (rows[i][0]) + "," + entregado + "," + (rows[i][5]) + "," + (rows[i][2])

				body := a
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				log.Printf(" [x] Sent %s", body)
				failOnError(err, "Failed to publish a message")
				b++

			}
		}
		if b == len(rows) {
			println("No hay nuevas entregas del camion Retail2.")
		}

	}

}

//Funcion encargada de enviar informacion del camion retail3 a finanzas
func senderToFinanzasN(segundos int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	b := 0
	for {
		rows := readRegistroN()
		for i := range rows {
			time.Sleep(time.Duration(segundos) * time.Second)
			if b != len(rows) {
				entregado := rows[i][6]
				if entregado == "0" {
					entregado = "No entregado"
				} else {
					entregado = "Entregado"
				}
				a := (rows[i][0]) + "," + entregado + "," + (rows[i][5]) + "," + (rows[i][2])

				body := a
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				log.Printf(" [x] Sent %s", body)
				failOnError(err, "Failed to publish a message")
				b++

			}
		}
		if b == len(rows) {
			println("no hay nuevas entregas del Camion Normal.")
		}

	}

}

func main() {
	fmt.Println("Abriendo servido de logística\n Bienvenido al sistema de Logistica!\n Seleccione cada cuantos segundos desea enviar informacion a finanzas: ")
	var segundos int
	fmt.Scanln(&segundos)

	go senderToFinanzasR1(segundos)
	go senderToFinanzasN(segundos)
	go senderToFinanzasR2(segundos)
	go Abrir(":50052", "Camiones")
	Abrir(":50051", "Clientes")

}
