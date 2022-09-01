PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
IMAGE_VERSION = v0.0.6

buildx-build-image-deviceshifu-http-http:
	docker buildx build --platform=linux/$(shell go env GOARCH) -f ${PROJECT_ROOT}/deviceshifu/Dockerfile.deviceshifu \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t edgehub/deviceshifu-http-http:${IMAGE_VERSION} --load