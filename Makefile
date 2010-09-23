include $(GOROOT)/src/Make.inc

TARG=backup
GOFILES=\
	$(TARG).go\
	error.go\
	prolog.go\
	os_helper.go\

include $(GOROOT)/src/Make.cmd

run:
	./$(TARG)
