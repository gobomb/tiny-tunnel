	
# the route:
|------------public network----------------|**Nat**|----------nat networt-------------|

**User**--pubConn--**Server**--pxyConn--|**Nat**|--**Client**--locConn--**LocalServer**

# server	
## proxy address
	
server listen for the client, wait for the proxy connection
	
`pxyAddr = "127.0.0.1:9991"`
	
## public address

server listen the request from the public network

`pubAddr = "127.0.0.1:9992"`

# client

## proxy address
    
client will dial the remote server, establish the proxy connection(and proxy connection)
    
`pxyAddr = "127.0.0.1:9991"`
   
## local address
    
client will dial the local server(the target server), establish the local connection

`locAddr = "127.0.0.1:22"`
    