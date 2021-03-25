# --------------------------------------------------

THEME_FILE ?= ./assets/css/theme.scss

NPX ?= npx
MARP_CLI_URL ?= @marp-team/marp-cli
MARP_CLI = $(NPX) $(MARP_CLI_URL)

# --------------------------------------------------

all: build

clean:

build: clean
	$(MARP_CLI) --theme $(THEME_FILE) cgo_ru.md -o cgo_ru.html --alow-local-files 
