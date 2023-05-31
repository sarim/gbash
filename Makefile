all: gbash
gbash: gbash.go
	go build gbash.go

clean:
	rm -f gbash gbash.o a.out
install: gbash
	sudo cp ./gbash /bin/gbash
	sudo chown root:root /bin/gbash
	sudo chmod u+sx /bin/gbash
	sudo sed -i '/gbash/d' /etc/shells
	echo '/bin/gbash' | sudo tee -a /etc/shells
run: install
	/bin/gbash
config: install
	chsh --shell /bin/gbash
	sudo /opt/distrod/bin/distrod enable
unconfig: clean
	chsh --shell /bin/zsh
	sudo sed -i '/gbash/d' /etc/shells
	sudo /opt/distrod/bin/distrod enable
