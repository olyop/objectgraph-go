package jsonbuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JSONBuilder struct {
	data *bytes.Buffer
}

func New() *JSONBuilder {
	return &JSONBuilder{
		data: &bytes.Buffer{},
	}
}

func (jb *JSONBuilder) Reset() {
	jb.data.Reset()
}

func (jb *JSONBuilder) Bytes() []byte {
	return jb.data.Bytes()
}

func (jb *JSONBuilder) WriteObject(object JSONObject) {
	jb.data.WriteRune('{')

	for i, pair := range object {
		jb.writeKey(pair.Key)
		jb.writeValue(pair.Value)
		jb.writeComma(i, len(object))
	}

	jb.data.WriteRune('}')
}

func (jb *JSONBuilder) WriteArray(objects []JSONObject) {
	jb.data.WriteRune('[')

	for i, object := range objects {
		jb.WriteObject(object)
		jb.writeComma(i, len(objects))
	}

	jb.data.WriteRune(']')
}

func (jb *JSONBuilder) writeKey(key string) {
	jb.data.WriteRune('"')
	jb.data.WriteString(key)
	jb.data.WriteRune('"')
	jb.data.WriteRune(':')
}

func (jb *JSONBuilder) writeValue(value *JSONValue) {
	if value.JSONType == 0 {
		jb.writeData(value.Data)
	} else if value.JSONType == 1 {
		jb.WriteObject(value.Object)
	} else if value.JSONType == 2 {
		jb.WriteArray(value.Array)
	} else {
		panic(fmt.Sprintf("unsupported JSON type: %d", value.JSONType))
	}
}

func (qp *JSONBuilder) writeData(data any) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	qp.data.Write(b)
}

func (jb *JSONBuilder) writeComma(i, length int) {
	if i == length-1 {
		return
	}

	jb.data.WriteRune(',')
}

type JSONValue struct {
	JSONType int
	Data     any
	Object   JSONObject
	Array    []JSONObject
}

type JSONObject []JSONPair

type JSONPair struct {
	Key   string
	Value *JSONValue
}
