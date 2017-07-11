# context rpc

The `server` and `client` folders implement a simple Go rpc service using gorilla/rpc.


## Exercise

Start a single consul node:

$ docker run -p 8400:8400 -p 8500:8500 -p 8600:53/udp -h `hostname` progrium/consul -server -bootstrap -ui-dir /ui


Register the service in consul:

Get local IP address with `ifconfig`  Find the public ip.

run the service:
SERVER_IP=x.x.x.x ./server


Run the `dig` command to see it in Consul's DNS:

dig @127.0.0.1  hello.service.consul SRV    

Run `curl` to see the consul API's registry version of the same record:

curl http://127.0.0.1:8500/v1/catalog/services

Modify the client to read the server address from the consul registry, then call the service
