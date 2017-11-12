default: target

.PHONY: target strip

target: http-go

http-go: http-go.go
	go build http-go.go

strip: target
	strip http-go

clean:
	rm -f http-go
