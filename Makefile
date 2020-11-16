#!/usr/bin/env make

bench:
	taskset --cpu-list 0-3 hey \
		-n 1000000 \
		-c 100 \
		-m POST \
		-T "application/json" \
		-d '{"name":"some_string"}' \
		http://localhost:50051

server-node-express: build-node-express start-node-express

build-node-express:
		cd ./node/rest-express && \
		npm install && \
		npm run build

start-node-express:
		cd ./node/rest-express && \
		npm run start

server-go-std: build-go-std start-go-std

build-go-std:
		cd ./go/rest-std && \
		easyjson data/data.go && \
		go build .

start-go-std:
		cd ./go/rest-std && \
		./rest-std
