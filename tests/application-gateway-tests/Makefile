APP_NAME = application-gateway-tests
APP_PATH = tests/application-gateway-tests
BUILDPACK = eu.gcr.io/kyma-project/test-infra/buildpack-golang-toolbox:v20191011-51ed45a
SCRIPTS_DIR = $(realpath $(shell pwd)/../..)/common/makefiles

include $(SCRIPTS_DIR)/generic-make-go.mk

VERIFY_IGNORE := /vendor\|/mocks

test:
	@echo "Tests are not to be run at this stage."

.PHONY: path-to-referenced-charts
path-to-referenced-charts:
	@echo "resources/application-connector"
