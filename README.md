# D-REQ GoLang

This small project is meant to improve my Golang.

D-Req is a client server application meant for a home lan, which uses the Walk GUI [walk GUI] (https://github.com/lxn/walk). 
A tcp client makes a simple request to a tcp server component in the house
requesting specific downloads.
The server will receive these requests sort them by priority in a csv file and notify the client once completed.

Make sure the sockets are directed towards the ip address of the machine you want to use as a server, and set the Path in server.go where you want the csv file to be created

Two options to run:
1. Open two CMDs, one in the server and one in the client directories
2. From the 2 CMDs execute: go run server.go / go run server.go
3. Follow client.go prompt

=========================

1. Create two executables and run from cmd
2. go executables are created with command go build *file*.go
3. Then just file.exe on cmd (server doesn't have a gui)

Next Possible Steps:
1. ~~Create a dumbed down GUI for this~~
2. Create Gui for Server
3. Make server send reply when download is completed
