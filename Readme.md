#D-REQ GoLang

This small project is meant to improve my Golang.

D-Req is a client server application meant for a home lan. 
A tcp client makes a simple request to a tcp server component in the house
requesting specific downloads.
The server will receive these requests sort them by priority and notify the client once completed

Make sure the sockets are directed towards the ip address of the machine you want to use as a server, and set the Path in server.go where you want the csv file to be created

Two options to run:
1. Open two CMDs on in the server and one in the client directories
2. From the 2 CMDs execute: go run server.go / go run server.go
3. Follow client.go prompt

=========================
1. Create two executables and run from cmd
2. go executables are created with command go build *file*.go


Next Step:
    Create a dumbed down GUI for this