// Server
package main 

import(
	"fmt"
	"net"
)



func main(){

	// Server initlization 
	go server()

	// Await for user inputs
	var userInput string 
	fmt.Scanln(&userInput)

}

func server(){

	// Open port
	s, err := net.Listen("tcp", ":7070")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for cannections
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn){
	b:= make( []byte, 100)
	bs, err := c.Read(b)
	if err != nil{
		fmt.Println(err)
		return
	} else{
		fmt.Println("Mensaje", string(b[:bs]))
		fmt.Println("Bystes", bs)
	}
}