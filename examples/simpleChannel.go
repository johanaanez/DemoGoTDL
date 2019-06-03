package main

import (
	"fmt"
)

//Public struct
type Client struct {
	cuil  string
	saldo float64
}

func hasAvailableBalance(client chan Client) {
	clientData := <-client
	if clientData.saldo > 0 {
		fmt.Printf("Cliente: %s tiene saldo \n", clientData.cuil)
	} else {
		fmt.Printf("Client: %s no tiene saldo \n", clientData.cuil)
	}
}

func main() {
	fmt.Println("Start")
	client := Client{"27123456784", 100}

	//Creacion de el channel, como puntero
	clientCh := make(chan Client)

	//deadlock
	//clientCh <- client

	//Se le pasa el channel vacío, no tiene data para procesar
	go hasAvailableBalance(clientCh)

	fmt.Println("Routine did not receive channel value yet")

	//La rutina recibe el dato que necesita para procesar
	clientCh <- client

	fmt.Println("End")
}
