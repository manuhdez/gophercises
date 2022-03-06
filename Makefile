SOURCE := $(PWD)

include $(SOURCE)/scripts/make/build.mk
include $(SOURCE)/scripts/make/run.mk
include $(SOURCE)/scripts/make/start.mk

deps:
	go mod download
	go mod tidy
	go mod verify

