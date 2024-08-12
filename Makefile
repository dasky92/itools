BIN_NAME=itools
BIN_PATH=~/.bin/

COMPLETION_FILES=bash zsh

.PHONY: pre build complete clean deploy install

install: build complete deploy
	@echo 'Install successfully.'

pre:
	@for file in ${COMPLETION_FILES}; do \
		if [ ! -d "$${HOME}/.$${file}_completion.d" ]; then \
			mkdir -p "$${HOME}/.$${file}_completion.d/"; \
		fi; \
	done

build:
	@CGO_ENABLED=0 go build -ldflags "-s -w" -o ${BIN_NAME}
	@chmod +x ${BIN_NAME}

complete: build
	@for file in ${COMPLETION_FILES}; do \
		${BIN_NAME} completion $${file} > ${BIN_NAME}_$${file}_completion; \
		chmod +x ${BIN_NAME}_$${file}_completion; \
	done

clean:
	@rm -f ${BIN_NAME}
	@for file in ${COMPLETION_FILES}; do \
		rm -f ${BIN_NAME}_$${file}_completion; \
	done

deploy: pre
	@mv ${BIN_NAME} ${BIN_PATH}
	@for file in ${COMPLETION_FILES}; do \
		mv ${BIN_NAME}_$${file}_completion "$${HOME}/.$${file}_completion.d/"; \
	done
