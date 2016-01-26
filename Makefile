all:
	docker run -v `pwd`:/go/src/takis -v `pwd`/bin:/go/bin -e CGO_ENABLED=0 google/golang go get -a -ldflags '-s' takis
