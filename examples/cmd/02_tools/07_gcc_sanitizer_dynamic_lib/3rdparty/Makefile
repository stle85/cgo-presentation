LIB := libfoo.so
INCLUDES := -Iinclude
LIBS :=

# --------------------------------------------------

BUILD_DIR ?= target/build

CC ?= gcc
CC_COMPILE = $(CC) -fPIC -c $(INCLUDES) $(LIBS)

RM = rm -fr
COPYDIR = cp -r
MKDIR = mkdir -p

# --------------------------------------------------

clean:
	$(RM) $(BUILD_DIR)

build: clean
	$(MKDIR) $(BUILD_DIR)/lib
	$(COPYDIR) include $(BUILD_DIR)/include

	$(CC_COMPILE) -shared src/foo.c -o $(BUILD_DIR)/lib/$(LIB)
