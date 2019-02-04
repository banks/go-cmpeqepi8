CASMS=$(patsubst %.cpp,%.s,$(wildcard cpp/*.cpp))
GOASMS=$(patsubst cpp/%.s,%_amd64.s,$(CASMS))

goasm: $(CASMS) $(GOASMS)

cpp/%.s: cpp/%.cpp
	c++ -O3 -masm=intel -mno-red-zone -mstackrealign -mllvm \
		-inline-threshold=1000 -fno-asynchronous-unwind-tables \
		-fno-exceptions -fno-rtti -S -mavx -mfma -o $@ $<

casm: $(CASMS)

install-deps:
	go get -u github.com/minio/c2goasm
	go get -u github.com/minio/asm2plan9s
	go get -u github.com/klauspost/asmfmt/cmd/asmfmt

%_amd64.s: cpp/%.s
	c2goasm -a -f -c $< $@

clean:
	rm $(CASMS)
	rm *_amd64.s

test: goasm
	go test ./...

bench: goasm
	go test -benchmem -bench . ./...


.PHONY: install-deps casm goasm clean