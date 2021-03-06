package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	student "github.com/nadirbasalamah/protodemo/src/model"
)

func main() {
	// create a student entity
	var newStudent student.Student = student.Student{
		Name:      "Nathan Mckane",
		StudentId: "RVN2021",
		CardNumber: &student.Student_IdCardNumber{
			IdCardNumber: "12345321",
		},
	}

	// add some courses
	newStudent.Courses = []*student.Course{
		{
			Code: "C001",
			Name: "Algorithm",
		},
		{
			Code: "C002",
			Name: "Data Structure",
		},
	}

	// print courses
	fmt.Println("Courses")
	for _, v := range newStudent.Courses {
		fmt.Println(v)
	}

	// convert to JSON
	var result string = convertToJSON(&newStudent)
	fmt.Println("to JSON: ", result)

	// convert to protocol buffers from JSON data
	var jsonData string = `{"name":"Ryan Cooper","studentId":"RPC2007","idCardNumber":"98753443"}`
	var protoResult student.Student = student.Student{}

	// call the function
	convertToProto(jsonData, &protoResult)
	fmt.Println("to proto: ", &protoResult)

}

func convertToJSON(pb proto.Message) string {
	var marshaler jsonpb.Marshaler = jsonpb.Marshaler{}
	result, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}
	return result
}

func convertToProto(data string, pb proto.Message) {
	err := jsonpb.UnmarshalString(data, pb)
	if err != nil {
		log.Fatalln("Cant convert to proto, ", err)
	}
}
