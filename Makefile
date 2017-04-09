# Disable make's implicit rules, which are not useful for golang, and
# slow down the build considerably.

VPATH=bin:l4lb:l4lb/spartan

#dcos-l4lb
L4LB=github.com/asridharan/dcos-cni-plugins/l4lb
L4LB_SRC=$(wilcard l4lb/*.go) $(wildcard l4lb/spartan/*.go)

PLUGINS=dcos-l4lb

.PHONY: all plugin
default: all

.PHONY: clean
clean:
	rm -rf vendor bin

# To update upstream dependencies, delete the glide.lock file first.
# Use this to populate the vendor directory after checking out the
# repository.
vendor: glide.yaml
	glide install -strip-vendor

dcos-l4lb:$(L4LB_SRC)
	mkdir -p `pwd`/bin
	go build -v -o `pwd`/bin/$@ $(L4LB)

plugin: vendor $(PLUGINS)

all: plugin

