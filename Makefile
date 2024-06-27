PLUGIN_DIRS=$(wildcard ./plugin/*)

all: build-plugins

clean: clean-plugins

build-plugins: $(PLUGIN_DIRS)

clean-plugins:
	rm -f ./bin/*

$(PLUGIN_DIRS):
	make -C $@

.PHONY: all $(PLUGIN_DIRS)
