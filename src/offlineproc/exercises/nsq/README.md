# Offline processing with NSQ

## Install NSQ

http://nsq.io/deployment/installing.html

Mac:  brew install nsq is painless
Pre-built binaries also available.

## Start NSQ

```
> nsqlookupd 
> nsqd --lookupd-tcp-address=127.0.0.1:4160 
> nsqadmin --lookupd-http-address=127.0.0.1:4161
```

## NSQ Admin

You can connect to the NSQ admin via your browser on here: [http://127.0.0.1:4171/](http://127.0.0.1:4171/)

## Get the go-nsq client

```sh
go get github.com/bitly/go-nsq
```

## Run the client
```sh
client
```

## Run the producer
This will add a message to the channel/topic

```sh
producer --website http://www.yahoo.com
```

You will see the client process the message

## View the output

By default, the output of the consumers are written to `/tmp/nsqoutput`

```sh
$ ls -d -1 /tmp/nsqoutput/*.* | xargs cat | jq
{
  "website": "http://www.yahoo.com",
  "statusCode": 200,
  "duration": "232.222045ms"
}
{
  "website": "bad-request",
  "error": "Get bad-request: unsupported protocol scheme \"\""
}
```
