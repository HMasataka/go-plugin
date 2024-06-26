FROM tinygo/tinygo:0.32.0

WORKDIR /go/src/app

COPY . .

RUN tinygo build -o plugin.wasm -scheduler=none -target=wasi --no-debug plugin.go
