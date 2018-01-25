guile-shell: *.go *.h *.c
	go build

clean:
	go clean

run-example-script: guile-shell
	./guile-shell -s ./log_now.scm

.PHONY: clean run-example-script
