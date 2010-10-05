include $(GOROOT)/src/Make.inc

TARG=backup
GOFILES=\
	$(TARG).go\
	error.go\
	prologue.go\
	helpers.go\
	pipeline.go\

include $(GOROOT)/src/Make.cmd

run:
	./$(TARG) -c test/config
