APP := app.out
INCLUDES := -I../lib
LIBS := -L../lib -Wl,-rpath=../lib -lreverse

# --------------------------------------------------

BUILD_DIR ?= target/build

CC ?= gcc

RM = rm -fr
COPYDIR = cp -r
MKDIR = mkdir -p

# --------------------------------------------------

clean:
	$(RM) $(BUILD_DIR)

build: clean
	$(MKDIR) $(BUILD_DIR)/bin
	$(CC) main.c $(INCLUDES) $(LIBS) -o $(BUILD_DIR)/bin/$(APP)

run: build
	$(BUILD_DIR)/bin/$(APP)
