# Raft Solution

This is a basic example of a key value store using raft as the consensus layer.

## How to start your cluster:

Install dependencies with `dep ensure`

### Start your first node:
```sh
raft -httpaddr localhost:8180 --raftaddr localhost:8186 /tmp/raft1
```

### Start your second node:
```sh
raft --httpaddr localhost:8280 --raftaddr localhost:8286 --join localhost:8180 /tmp/raft2
```

### Start your third node:
```sh
raft --httpaddr localhost:8380 --raftaddr localhost:8386 --join localhost:8180 /tmp/raft3
```

### Write your first value
```sh
curl -XPOST localhost:8180/key -d '{"foo": "bar"}'
```

### Read back the value
```sh
curl -XGET localhost:8180/key/foo
```

### Reading a value from a non-leader

This is an error, as you can only get a consensus read from the leader.
However, we are nice enough to put the new leader in the header so you could
write a client to redirect if needed.

```
$ curl -v -XGET localhost:8380/key/foo
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8380 (#0)
> GET /key/foo HTTP/1.1
> Host: localhost:8380
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 500 Internal Server Error
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< X-Raft-Leader: 127.0.0.1:8186
< Date: Sun, 30 Oct 2016 19:12:59 GMT
< Content-Length: 11
<
not leader
* Connection #0 to host localhost left intact
```

## Things that we didn't do:

- We did not handle writing or reading to a non-leader.  We still need to
  respond with the leader address, or allow our handlers to forward the request
  to the leader via redirect.
- Tests... yes, we need a lot of testing for this.
- Add a close method to clean up and shut down our raft and httpd services
  nicely
