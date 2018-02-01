src = ./api/pkg
main = ./api/cmd/main.go
pkgDir = $(src)/$(pkg)

.PHONY: build coverage dockerUp fmt goMessageChallenge install start test package unbindPort3000 unbindPort4200 ui-deps ui-dev ui-prod vet

build:
	docker-compose build

dockerUp:
	@docker-compose down
	@docker-compose up -d
	@python -m webbrowser "http://localhost:8080" &> /dev/null

coverage:
	@set -e;
	@echo "mode: set" > acc.out;

	@for Dir in $$(find . -type d); do \
		if ls "$$Dir"/*.go &> /dev/null; then \
			go test -coverprofile=profile.out "$$Dir"; \
			go tool cover -html=profile.out; \
		fi \
	done

	@rm -rf ./profile.out;
	@rm -rf ./acc.out;

fmt:
	@go fmt ./...

goMessageChallenge:
	@mkdir -p bin
	@cd ./cmd/ \
	&& go build -o goMessageChallenge \
	&& mv goMessageChallenge ../bin \
	&& cd ..

install:
	go install $(main)

package:
	@mkdir -p $(pkgDir)
	@echo package $(pkg) | tee $(pkgDir)/$(pkg).go $(pkgDir)/$(pkg)_test.go

start: unbindPort3000
	@rm -f $(main)/main
	go run $(main)

test:
	@go test -v $(src)/...

ui-deps:
	cd ui \
	&& npm i \
	&& cd ..

ui-dev: unbindPort4200
	cd ui \
	&& ng serve --open \
	&& cd ..

ui-prod: unbindPort4200
	cd ui \
	&& ng build --env=prod \
	&& cd ..

unbindPort3000:
	kill -9 $$(lsof -i :3000 | grep main | awk '{ print $$2}' | xargs) || true

unbindPort4200:
	kill -9 $$(lsof -i :4200 | grep ng | awk '{ print $$2}' | xargs) || true

vet:
	@go vet ./...