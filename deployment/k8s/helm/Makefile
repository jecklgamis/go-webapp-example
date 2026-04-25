APP_NAME := go-webapp-example
VERSION := $(shell yq .appVersion chart/Chart.yaml)
HELM_CHART := $(APP_NAME)-$(VERSION).tgz

.PHONY: package dry-run install uninstall upgrade rollback clean

package:
	rm -f $(HELM_CHART)
	helm package ./chart
dry-run:
	helm install --dry-run --debug $(APP_NAME) ./chart
install: package
	helm install $(APP_NAME) $(HELM_CHART)
uninstall:
	helm uninstall $(APP_NAME)
upgrade: package
	helm upgrade $(APP_NAME) $(HELM_CHART)
rollback:
	helm rollback $(APP_NAME)
clean:
	rm -f $(HELM_CHART)
