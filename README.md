# CS4032: Lab1

## Lab brief [here](https://www.scss.tcd.ie/Stephen.Barrett/lectures/cs4032/lab2.html)

## Description:

Golang server that accepts socket connections from clients and handles the following commands:
  1. HELO:
      a. Command: `"HELO <TEXT>"`
      b. Action: Respond with `"HELO <TEXT>\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"`
  2. KILL_SERVICE:
      a. Command: `"KILL_SERVICE\n"`
      b. Action: None, kill the server
  3. <ANY OTHER COMMAND>
      a. Command: `"<ANY TEXT>"`
      b. Action: Do nothing

## Prerequisits: 

Server code is written in Go. **compile.sh** will install Go for you.

## Running the server:
  1. Run **compile.sh** to install prerequisits and compile the code
  2. Run **start.sh <PORT NUMBER>** to start the server
