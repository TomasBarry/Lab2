// Import node socket module and socket connection variables
const socket = require('net').Socket();
const socketVariables = require('./socketVariables.js');

   
// Set up socket connection to the server
socket.connect(socketVariables.serverPort, socketVariables.serverIPAddress);


/**
 * Create handlers for the response from the server
 */

// Print the response from the server to the console and close the socket
socket.on('data', (data) => {
    console.log(data.toString());
    socket.end();
});

// Handle error in connecting to the server
socket.on('error', (err) => {
    console.log('Error connecting to server');
    socket.end();
});

// Close the socket connection
socket.on('end', () => {
    console.log('Bytes written on the socket: ' + socket.bytesWritten);
    console.log('Bytes read on the socket: ' + socket.bytesRead);
    console.log('Disconnected');
});


// Send message to the echo server through the socket connection
socket.write('GET http://' + socketVariables.serverIPAddress + '/echo.php?message=hello HTTP/1.1\n\n');
