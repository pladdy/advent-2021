day-%.go:
	cp day-template.go $@

run-%:
	go run day-$*/day-$*.go