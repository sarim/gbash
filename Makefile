all: gbash
gbash: gbash.go
	go build gbash.go

clean:
	rm -f gbash gbash.o a.out
install: gbash
	sudo cp ./gbash /bin/gbash
	sudo chown root:root /bin/gbash
	sudo chmod u+sx /bin/gbash
run: install
	/bin/gbash
config: install
	chsh --shell /bin/gbash
	sudo /opt/distrod/bin/distrod enable
unconfig:
	chsh --shell /bin/bash
	sudo /opt/distrod/bin/distrod enable
