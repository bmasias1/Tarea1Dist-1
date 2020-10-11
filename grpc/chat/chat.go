package chat

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"encoding/csv"
	"io/ioutil"
	"strings"

	"golang.org/x/net/context"
)

var colaRetail []paquete
var colaPrioritario []paquete
var colaNormal []paquete

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func EscribirCSV(aEscribir string) string {

	nroOrden, err := ioutil.ReadFile("../archivos/indexAct.data")
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile("../archivos/registro.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return "0"
	}
	w := csv.NewWriter(f)
	nroOrdenStr := string(nroOrden[:])

	Hora := time.Now()
	HoraStr := Hora.Format("2006-01-02 15:04:05")

	aEscribir = HoraStr + "+" + aEscribir + "+" + nroOrdenStr

	nueva := strings.Split(aEscribir, "+")
	retorno := "El código de seguimiento de " + nueva[1] + " es " + nroOrdenStr
	w.Write(nueva)

	var aux paquete //anadiendo a las colas en memoria
	ValorInt, _ := strconv.Atoi(nueva[4])
	aux = paquete{nueva[1], nueva[7], nueva[2], ValorInt, 0, "En bodega"}
	if nueva[2] == "retail" {
		colaRetail = append(colaRetail, aux)
	} else if nueva[2] == "prioritario" {
		colaPrioritario = append(colaPrioritario, aux)
	} else {
		colaNormal = append(colaNormal, aux)
	}

	nroOrdenInt, err := strconv.Atoi(nroOrdenStr)
	bs := []byte(strconv.Itoa(nroOrdenInt + 1))

	error := ioutil.WriteFile("../archivos/indexAct.data", bs, 0777)
	if error != nil {
		fmt.Println(err)
	}

	w.Flush()
	return retorno
}

type paquete struct {
	IDPaquete   string
	seguimiento string
	tipo        string
	valor       int
	intentos    int
	estado      string
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	mensaje := strings.Split(message.Body, "_")
	log.Printf("Mensaje recibido desde el cliente: %s", mensaje[0])
	if mensaje[1] == "orden" {
		mensaje := EscribirCSV(mensaje[0])
		return &Message{Body: mensaje}, nil
	} else { //Si se mandó un codigo
		if colaRetail == nil && colaNormal == nil && colaPrioritario == nil {
			fmt.Println("Codigo no encontrado")
			return &Message{Body: "No hay ordenes ingresadas"}, nil
		} else {
			for _, orden := range colaRetail {
				if orden.seguimiento == mensaje[0] {
					return &Message{Body: orden.estado}, nil
				}
			}
			for _, orden := range colaNormal {
				if orden.seguimiento == mensaje[0] {
					return &Message{Body: orden.estado}, nil
				}
			}
			for _, orden := range colaPrioritario {
				if orden.seguimiento == mensaje[0] {
					return &Message{Body: orden.estado}, nil
				}
			}
			return &Message{Body: "No se encontró el código"}, nil
		}
	}
}

func (s *Server) SayHelloAgain(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Recibido el mensaje del Camion: %s", in.Otro)
	return &Message{Body: "ola de nuevo!"}, nil
}
