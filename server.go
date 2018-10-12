package main
import (
  "fmt"
  "net"
  "strings"
)
//Change networking properties here
const (
  HOST = "localhost"
  PORT = "3333"
  TYPE = "tcp"
)
//Spawn a tcp server
func main() {

  fmt.Println("Started server @ " + HOST + ":" + PORT)
  listen, _ := net.Listen(TYPE, HOST + ":" + PORT)
  
  for {
      // Listen for an incoming connection. 
      connection, _ := listen.Accept()

      // Make a buffer to hold incoming data.
      buffer := make([]byte, 1024)

      //Read the Ip address
      userIP := connection.RemoteAddr().String()
      userIP = convertAddress(userIP)

      //Read the incoming connection into the buffer.
      requestLength, _ := connection.Read(buffer)

      //Pint out the message to a string
      message := string(buffer[:requestLength])
      fmt.Print("<" + userIP + ">: " + message)

      //Bounce message back
      message = "<" + userIP + ">: " + message
      bounceByteArray := []byte(message)
      connection.Write(bounceByteArray)

      // Close the connection when you're done with it.
      connection.Close()
  }
}
//Converts address to a string with dictated port
func convertAddress(raw string)(string){
  var data [5]string
  stringSlice := strings.Split(raw, ":")
  for i := range stringSlice {
    data[i]=stringSlice[i]
  }
  userIP := data[0]+"."+":3333"
  return userIP
}