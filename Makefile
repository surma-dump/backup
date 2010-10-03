include $(GOROOT)/src/Make.inc

TARG=backup
GOFILES=\
	$(TARG).go\
	error.go\
	prolog.go\
	helpers.go\

include $(GOROOT)/src/Make.cmd

run:
	./$(TARG) -c test/config
