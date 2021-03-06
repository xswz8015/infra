package common

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	structpb "github.com/golang/protobuf/ptypes/struct"
	structbuilder "google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// JSONPBUnmarshaler unmarshals JSON into proto messages.
var JSONPBUnmarshaler = jsonpb.Unmarshaler{AllowUnknownFields: true}

// MapToStruct constructs a Struct from the given map[string]interface{}. The
// map keys must be valid UTF-8. The map values can be any of Go's basic types
// (bool, string, number type, byte, or rune), a proto message (in the form
// protoreflect.ProtoMessage), or a nested map[string]interface{} that fulfils
// the same requirements recursively.
//
// NOTE: This function is just a modified version of structpb.NewStruct(), with
// added logic to handle the case where the map value is a proto message. This
// is necessary because Buildbucket request interfaces are almost always
// implemented as proto messages at some level.
func MapToStruct(m map[string]interface{}) (*structpb.Struct, error) {
	s := &structpb.Struct{Fields: make(map[string]*structpb.Value, len(m))}
	for key, val := range m {
		if !utf8.ValidString(key) {
			return nil, fmt.Errorf("invalid UTF-8 in string: %q", key)
		}
		var err error
		var newStructVal *structpb.Value

		switch val := val.(type) {
		case protoreflect.ProtoMessage:
			newStructVal, err = protoToStructVal(val)
			if err != nil {
				return nil, fmt.Errorf("error converting proto %v to *structpb.Value: %s", val, err)
			}
		case map[string]interface{}:
			// Recursively call MapToStruct. The default case of
			// calling structbuilder.NewValue() below also attempts to handle
			// this case recursively, but would throw an error if the inner map
			// contains a proto.
			innerStruct, err := MapToStruct(val)
			newStructVal = structbuilder.NewStructValue(innerStruct)
			if err != nil {
				return nil, err
			}
		default:
			newStructVal, err = structbuilder.NewValue(val)
			if err != nil {
				return nil, fmt.Errorf("error callling structbuilder.NewValue(%v): %s", val, err)
			}
		}
		s.Fields[key] = newStructVal
	}

	return s, nil
}

// protoToStructVal converts a proto message to a *structpb.Value.
func protoToStructVal(msg protoreflect.ProtoMessage) (*structpb.Value, error) {
	m := jsonpb.Marshaler{}
	msgJSON, err := m.MarshalToString(msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	msgStruct := &structpb.Struct{}
	err = JSONPBUnmarshaler.Unmarshal(strings.NewReader(msgJSON), msgStruct)
	if err != nil {
		return nil, err
	}
	return &structpb.Value{
		Kind: &structpb.Value_StructValue{StructValue: msgStruct},
	}, nil
}

// OffsetTimestamp returns a timestamp offset from the current time by the
// given number of minutes.
func OffsetTimestamp(minutes int64) *timestamppb.Timestamp {
	return timestamppb.New(time.Now().Add(time.Duration(minutes) * time.Minute))
}
