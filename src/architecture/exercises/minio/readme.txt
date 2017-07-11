
Run a Minio server:
docker run -v /mnt/minio:/minio -d -p 9000:9000 minio/minio:latest server /minio

Get the token/accessKey

	docker ps (get minio container id)
	docker logs (minio container id)

Run a Go app to talk to minio:

	edit $GOPATH/src/architecture/exampmles/minio/main.go to use YOUR keys/ID
	go build
	./minio

Now see the file stored in minio:
	browse minio at http://server.IP:9000
