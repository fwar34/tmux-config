# http://c.biancheng.net/makefile/
all:
	go build

.PHONY:install uninstall clean
install:all
	sudo cp collect-go /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/collect-go

clean:
	rm collect-go
