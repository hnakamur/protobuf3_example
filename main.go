package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/hnakamur/protobuf3_example/example"
)

func main() {
	test := &example.Test{
		Label: "hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.Label != newTest.Label {
		log.Fatalf("data mismatch %q != %q", test.Label, newTest.Label)
	}
	// etc.
}
