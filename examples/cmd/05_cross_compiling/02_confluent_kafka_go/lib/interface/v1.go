package main

import "C"
import "github.com/confluentinc/confluent-kafka-go/kafka"

//export test_v1_get_version
func test_v1_get_version() C.int {
	version, _ := kafka.LibraryVersion()
	return C.int(version)
}
