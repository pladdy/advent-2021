day-%.go:
	cp day-template.go $@

day-%:
	go run $@.go