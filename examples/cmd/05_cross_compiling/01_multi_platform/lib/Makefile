# sudo apt install gcc-mingw-w64-x86-64

# --------------------------------------------------

TARGET ?= libsum_glibc_x64.a
INCLUDES :=
LIBS :=

# --------------------------------------------------

CC ?= gcc
AR ?= ar

# --------------------------------------------------

build:
	$(CC) -c libsum.c $(INCLUDES) $(LIBS) -o libsum.o
	$(AR) rcs  $(TARGET) libsum.o
	$(RM) libsum.o


build-mingw: CC=/usr/bin/x86_64-w64-mingw32-gcc-win32
build-mingw: AR=/usr/bin/x86_64-w64-mingw32-ar
build-mingw: build
