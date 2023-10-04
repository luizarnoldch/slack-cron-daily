build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc
	zip lambda.zip bootstrap
	rm -rf bootstrap
deploy:
	sam deploy --stack-name hello-test --guided --resolve-s3
