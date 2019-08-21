.PHONY: cmds
ECHO:  
	@echo $@
cmds:ECHO 
	@echo $(GOPATH)
	rm -rf $(GOPATH)/bin/cmds
	go build -o $(GOPATH)/bin/cmds cmds
.PHONY: libs 
libs:
	cd src/cmds && dep ensure
