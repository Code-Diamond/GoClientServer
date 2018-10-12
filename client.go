package main
import (
    "net"
    "bufio"
    "os"
)
//Client side
func main() {
    for{
        //Get input from client user
        print(">>")
        reader := bufio.NewReader(os.Stdin)
        message, _ := reader.ReadString('\n')
        
        //Connect to the server 
        serverAddress := "localhost:3333" //Change ip/port information here

        tcpAddr, _ := net.ResolveTCPAddr("tcp", serverAddress)
        connection, _ := net.DialTCP("tcp", nil, tcpAddr)
        
        //Write to the server
        connection.Write([]byte(message))
        
        //get bounced byte
        reply := make([]byte, 1024)
        connection.Read(reply)
        
        //Print and close connection
        print(string(reply))
        connection.Close()        
    }
}