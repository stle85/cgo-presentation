TARGET ?= app.out

# --------------------------------------------------

GO ?= go
GOCMD := $(GO)
GOBUILD := $(GO) build -x

TEST_NAME ?= .
GOBENCHMARK_RUN := $(GO) test -bench=$(TEST_NAME)
GOTEST_RUN := $(GO) test -v -run $(TEST_NAME)

RM ?= rm -fr
TIME ?= time -f '\nExecution time: %E'

PERF ?= perf
PERF_RECORD_RUN := $(PERF) record
PERF_REPORT_RUN := $(PERF) report -T

VALGRIND ?= valgrind
VALGRIND_RUN := $(VALGRIND) -s \
	--tool=memcheck \
	--leak-check=full \
	--show-leak-kinds=all \
	--track-origins=yes

CALLGRIND ?= $(VALGRIND) --tool=callgrind
CALLGRIND_RUN := $(CALLGRIND)

CALLGRIND_ANNOTATE ?= callgrind_annotate
CALLGRIND_ANNOTATE_RUN := $(CALLGRIND_ANNOTATE)

# --------------------------------------------------

.PHONY: $(TARGET)

default: all

all: build

clean: $(TARGET)
	$(RM) $(TARGET)

build: clean
	$(GOBUILD) -o $(TARGET) .

build-time: clean
	$(TIME) $(GOBUILD) -o $(TARGET) .

build-c-static-lib:
	$(GOBUILD) -buildmode=c-archive -o $(TARGET) .

build-c-shared-lib:
	$(GOBUILD) -buildmode=c-shared -o $(TARGET) .

run: build
	./$(TARGET)

run-less: build
	./$(TARGET) 2>&1 | head -n 30

go-benchmark:
	$(GOBENCHMARK_RUN)

go-test:
	$(GOTEST_RUN)

perf: build
	$(PERF_RECORD_RUN) ./$(TARGET)
	$(PERF_REPORT_RUN)

show-perf-report:
	$(PERF_REPORT_RUN)

callgrind: build
	$(CALLGRIND_RUN) --callgrind-out-file=callgrind.out ./$(TARGET)
	$(CALLGRIND_ANNOTATE_RUN) callgrind.out

show-callgrind-report:
	$(CALLGRIND_ANNOTATE_RUN) callgrind.out

valgrind: build
	$(VALGRIND_RUN) ./$(TARGET)

vet:
	$(GO) vet .
