
all: build

build: build-stamp
build-stamp: main.go
	@echo start build target
	@echo GOPATH: $(GOPATH)
	# always format code
	# GOPATH=$(GOPATH) go fmt $^
	# binary
	GOPATH=$(GOPATH) go build -v jupiter
	touch $@

check:
	@echo check tartget
	@echo GOBIN: $(GOBIN)
	@echo GOPATH: $(GOPATH)
	@echo PATH: $(PATH)
	GOPATH=$(GOPATH) go get -v github.com/astaxie/beego
	#GOPATH=$(GOPATH) go get -v github.com/beego/bee
	GOPATH=$(GOPATH) go get -v gopkg.in/couchbaselabs/gocb.v1
	GOPATH=$(GOPATH) go get -v github.com/dghubble/sling
	GOPATH=$(GOPATH) go get -v github.com/google/uuid
	GOPATH=$(GOPATH) go get -v gopkg.in/jmcvetta/napping.v3
	GOPATH=$(GOPATH) go get -v github.com/denisenkom/go-mssqldb
	
clean:
	# remove to go get fresh new ones in every build
	rm -rf ./github.com

install:
	@echo deploy target
	@echo env is $(env)
	ssh root@10.8.145.6 service supervisord stop
	scp ./jupiter root@10.8.145.6:/home/jupiter_shop/
ifeq "$(env)" "prod"
	scp -r ../conf root@10.8.145.6:/home/jupiter_shop/
	#scp -r ./conf root@10.8.145.6:/home/jupiter_shop/
else
	scp -r ./conf root@10.8.145.6:/home/jupiter_shop/
endif
	scp -r ./static root@10.8.145.6:/home/jupiter_shop/
	scp -r ./views root@10.8.145.6:/home/jupiter_shop/
	ssh root@10.8.145.6 service supervisord start
