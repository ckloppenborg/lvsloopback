###lvsloopback
The goal of the project is to allow keepalived the ability to do health checking of servers based on the VIP loopback being up.

http://192.168.1.105:8888/192.168.1.113
 
This will send an HTTP request to the real server 192.168.1.105 with the uri 192.168.1.113 being the VIP. When the request is received by the server it verifies the uri is an IP address and then checks to see if that IP address is active on the server.

