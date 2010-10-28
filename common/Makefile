include $(GOROOT)/src/Make.inc

myall: all
	cp _obj/common.a .

TARG=common
GOFILES=\
	$(TARG).go\
	backup.go\

CLEANFILES=\
	common.a\

include $(GOROOT)/src/Make.pkg
