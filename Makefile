

# https://tutorialedge.net/golang/makefiles-for-go-developers/

hello:
	echo "hello world! from Makefile!!!"

love:
	go build -o bin/helloworld helloworld.go

run:
	go run helloworld.go




