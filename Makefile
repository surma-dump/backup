all:
	make -C common
	make -C client
	make -C server
	cp client/backupc .
	cp server/backups .

clean:
	-@rm backupc backups
	make -C common clean
	make -C client clean
	make -C server clean

