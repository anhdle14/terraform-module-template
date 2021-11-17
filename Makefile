# Terraform modules should not have provider configuration block.
asset:
	zip module.zip -r *.tf modules

check-provider:
	grep -q "^provider" *.tf && exit 1 || exit 0

fmt:
	terraform fmt -check -recursive -no-color

lint:
	tflint

validate:
	terraform validate -json

test:
	cd tests && go test -v ./... -timeout 30m

clean:
	rm -rf *.zip
	find . -type f -name 'terraform.*' -exec rm -vrf {} +
	find . -type d -name '.terraform' -exec rm -vrf {} +

delete-all:
	@git submodule deinit -f --all
	@git rm -rf modules/*
	@rm -rf .git/modules/modules/*

add:
	@echo Fetching module: $(module) with $(module_url)
	@git submodule add $(module_url) modules/$(module)
	@git submodule update --init --recursive

update:
	@echo "Updating module: $(module) with specific version: $(version)"
	$(shell cd modules/$(module); git fetch; git checkout $(version); git pull)

delete:
	@echo "Deleting module: $(module)"
	@git submodule deinit -f modules/$(module)
	@git rm -f modules/$(module)
	@rm -rf .git/modules/modules/$(module)

.PHONY: asset check-provider fmt lint validate test clean add update delete delete-all
