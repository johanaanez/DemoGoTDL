package main

import (
	"fmt"
)

//Public struct
type Client struct {
	cuil     string
	balance float64
}

func hasAvailableBalance(client chan Client){
	clientData := <-client
	if clientData.balance >= 0 {
		fmt.Printf("Client: %s has positive balance \n", clientData.cuil)
	}else{
		fmt.Printf("Client: %s has negative balance \n", clientData.cuil)
	}

}

func main() {
	fmt.Println("Start")
	client := Client{"27939477654", 100}

	//Creacion de el channel, como puntero
	clientCh := make(chan Client)

	//deadlock
	//clientCh <- client

	//Se le pasa el channel vacío, no tiene data para procesar
	go hasAvailableBalance(clientCh)

	fmt.Println("Routine does not receive channel value yet")

	//La rutina empieza recién acá cuando recibe el dato que necesita para procesar
	clientCh <- client

	fmt.Println("End")
}
