build:
	GOOS=linux go build -o app
	docker image build -t sdomula/say .
	rm -f app
