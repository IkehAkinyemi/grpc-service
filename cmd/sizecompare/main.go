// Provides functionality for comparing the sizes of data serialized
// using different wire data serialization protocols. It includes functions for
// serializing data using JSON, XML, and Protocol Buffer protocols, and comparing the
// resulting data sizes.

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/IkehAkinyemi/grpc-service/gen"
	"google.golang.org/protobuf/proto"
)


var genCreateUserRequest = &gen.CreateUserRequest{
	Name: `
	ICDU80SIZW
	TZ3VOGW03E
	U9A09CWOI2
	0MXR23DTX3
	H7JYD0ID25
	6BXFRBZTKE
	OUISRY3RE8
	VAESTUVFNV
	JXQWO5HT90
	O9D7DHA4H2
	7Y2NWRSE0T
	DS6U2D4QTO
	X22LO4WL86
	IKVMAHBZRH
	NW0I2ICQXK
	M9OI173MCJ
	14GY4QSEP2
	ZYXVJ59D6U
	T5C3MIIZSG
	PITUA5F81P
	24BDEY5R68
	E9JVTDMCD7
	TC77ES64ND
	CZXSZFFWC2
	3P3BFIH8UG
	IU0CO4DW23
	8BBVH4AVTO
	7SJD13DCIU
	UOAOROQGNV
	FMJE4SAA5R	
	`,
	Age: 999999999,
}

func main() {
	jsonBytes, err := serializeToJSON(genCreateUserRequest)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serializeToXML(genCreateUserRequest)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genCreateUserRequest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size:\t%dB\n", len(jsonBytes))
	fmt.Printf("XML size:\t%dB\n", len(xmlBytes))
	fmt.Printf("proto size:\t%dB\n", len(protoBytes))
}

// serializeToJSON uses the JSON data serialization protocol for wire data
func serializeToJSON(u *gen.CreateUserRequest) ([]byte, error) {
	return json.Marshal(u)
}

// serializeToXML uses the XML data serialization protocol for wire data
func serializeToXML(u *gen.CreateUserRequest) ([]byte, error) {
	return xml.Marshal(u)
}
// serializeToProto uses the Protocol Buffer data serialization protocol for wire data
func serializeToProto(u *gen.CreateUserRequest) ([]byte, error) {
	return proto.Marshal(u)
}
