.PHONY: build_run
build_run_server:
	go build -v ./bin/archiver
	./archiver.exe


.DEFAULT_GOAL := build_run