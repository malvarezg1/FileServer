// Client
package main 

import(
	"fmt"
	"net"
)

func main(){
	 
	go client()

	var userInput string 
	fmt.Scanln(&userInput)

}

func client(){
	c, err := net.Dial("tcp", ":7070")
	if err != nil{
		fmt.Println(err)
		return
	}

	msg := "HolaMundo desde el cliente"
	fmt.Println(msg)
	c.Write([]byte(msg))
	c.Close()
}