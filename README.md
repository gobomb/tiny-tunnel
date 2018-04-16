	
# the route:
|----------public network------------|----------nat networt-------------|

**User**--pubAddr--**Server**--ctlAddr--**Client**--locAddr--**LocalServer**

# server	
## control address
	
server listen for the client, wait for the control connection
	
`ctlAddr = "127.0.0.1:9991"`
	
## public address

server listen the request from the public network

`pubAddr = "127.0.0.1:9992"`

# client side

## control address
    
client will dial the remote server, establish the control connection(and proxy connection)
    
TODO: separate the cltConn and the pxyConn

`ctlAddr = "127.0.0.1:9991"`
   
## local address
    
client will dial the local server(the target server), establish the local connection

`locAddr = "127.0.0.1:22"`
    