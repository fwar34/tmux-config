# http://c.biancheng.net/makefile/
all:
	go build

.PHONY:install uninstall
install:all
	sudo cp collect-go /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/collect-go

.PHONY:clean
clean:
	rm collect-go
