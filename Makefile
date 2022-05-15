build_docker:
	docker build --build-arg GOARCH=amd64 --platform linux/amd64 --progress plain --tag orlangure/gnomock-test-image:amd64 .
	docker build --build-arg GOARCH=arm64 --platform linux/arm64 --progress plain --tag orlangure/gnomock-test-image:arm64 .
	docker push orlangure/gnomock-test-image:amd64
	docker push orlangure/gnomock-test-image:arm64
	docker manifest create --amend orlangure/gnomock-test-image:latest orlangure/gnomock-test-image:amd64 orlangure/gnomock-test-image:arm64
	docker manifest push orlangure/gnomock-test-image:latest

