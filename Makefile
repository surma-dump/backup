all:
	make -C common all
	make -C client all
	make -C server all
	cp client/backupc .
	cp server/backups .

clean:
	rm backupc backups
	make -C common clean
	make -C client clean
	make -C server clean

