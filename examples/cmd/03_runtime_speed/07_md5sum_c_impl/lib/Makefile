LIB := libhash.so
INCLUDES :=
LIBS := -lcrypto -lpthread

# --------------------------------------------------

CC ?= gcc

# --------------------------------------------------

build:
	$(CC) -shared -fPIC libhash.c $(INCLUDES) $(LIBS) -o $(LIB)
