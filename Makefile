#
# Makefile for mixderdl
#
.PHONY: usage edit build run play clean git

PROG=mixerdl
usage:
	@echo "usage: make [edit|build|run|play|clean]"

edit e:
	vi mixerdl.go

build b:
	go build -o $(PROG) *.go
	@mv $(PROG) $(GOPATH)/bin
	@ls -al $(GOPATH)/bin/$(PROG)

run r:
	$(PROG) -url="https://mixer.com/Kabby?vod=WVKDcVRHNEOFt3o7H0-l5g"

play p:
	ffplay *.mp4

clean:
	rm -f $(PROG) *.mp4
#----------------------------------------------------------------------------------
git g:
	@echo "> make (git:g) [update|store]"
git-update gu:
	make clean
	git add .
	git commit -a -m "clean code"
	git push
git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------

