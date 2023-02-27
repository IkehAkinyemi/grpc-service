package main

import "testing"

func BenchmarkSerializeToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToJSON(genCreateUserRequest)
	}
}

func BenchmarkSerializeToXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToXML(genCreateUserRequest)
	}
}

func BenchmarkSerializeToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToProto(genCreateUserRequest)
	}
}
