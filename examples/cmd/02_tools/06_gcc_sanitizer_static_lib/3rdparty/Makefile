LIB := libfoo.a
INCLUDES := -Iinclude
LIBS :=

# --------------------------------------------------

BUILD_DIR ?= target/build

CC ?= gcc
CC_COMPILE = $(CC) $(INCLUDES) $(LIBS)

AR ?= ar
AR_CREATE = $(AR) rcs


RM = rm -fr
COPYDIR = cp -r
MKDIR = mkdir -p

# --------------------------------------------------

clean:
	$(RM) $(BUILD_DIR)

build: clean
	$(MKDIR) $(BUILD_DIR)/obj
	$(MKDIR) $(BUILD_DIR)/lib
	$(COPYDIR) include $(BUILD_DIR)/include

	$(CC_COMPILE) -c src/foo.c -o $(BUILD_DIR)/obj/foo.o
	$(AR_CREATE) $(BUILD_DIR)/lib/$(LIB) $(BUILD_DIR)/obj/*.o
