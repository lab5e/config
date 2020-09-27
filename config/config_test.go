package config

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	const id = uint32(1)
	const command = uint32(2)

	setValues := &Request{
		Id:      id,
		Command: command,
		Values: []*Value{
			{Id: 1, Int32Val: int32(1234)},
			{Id: 2, Int64Val: int64(5678)},
			{Id: 3, DoubleVal: float64(3.14)},
			{Id: 4, StringVal: "this is a test"},
			{Id: 5, BytesVal: []byte{1, 2, 3, 4}},
		},
	}

	// Marshal the message into a byte array
	out, err := proto.Marshal(setValues)
	assert.Nil(t, err)
	assert.NotZero(t, out)

	// Then unmarshal
	unmarshalled := &Request{}
	err = proto.Unmarshal(out, unmarshalled)
	assert.Nil(t, err)

	// Make sure that the values that should be there are present
	assert.Equal(t, id, unmarshalled.Id)
	assert.Equal(t, command, unmarshalled.Command)
	assert.Equal(t, uint32(1), unmarshalled.Values[0].Id)
	assert.Equal(t, int32(1234), unmarshalled.Values[0].Int32Val)

	assert.Equal(t, uint32(2), unmarshalled.Values[1].Id)
	assert.Equal(t, int64(5678), unmarshalled.Values[1].Int64Val)

	assert.Equal(t, uint32(3), unmarshalled.Values[2].Id)
	assert.Equal(t, float64(3.14), unmarshalled.Values[2].DoubleVal)

	assert.Equal(t, uint32(4), unmarshalled.Values[3].Id)
	assert.Equal(t, "this is a test", unmarshalled.Values[3].StringVal)

	assert.Equal(t, uint32(5), unmarshalled.Values[4].Id)
	assert.Equal(t, []byte{1, 2, 3, 4}, unmarshalled.Values[4].BytesVal)

	// Make sure that the values that shouldn't be there are Zero
	// values for that type
	assert.Zero(t, unmarshalled.Values[0].Int64Val)
	assert.Zero(t, unmarshalled.Values[0].DoubleVal)
	assert.Zero(t, unmarshalled.Values[0].StringVal)
	assert.Zero(t, unmarshalled.Values[0].BytesVal)
}

func TestResponse(t *testing.T) {
	const id = uint32(1)
	const command = uint32(2)
	const sequence = uint32(3)

	response := &Response{
		Id:       id,
		Command:  command,
		Sequence: sequence,
		Values: []*Value{
			{Id: 1, Int32Val: int32(1234)},
			{Id: 2, Int64Val: int64(5678)},
			{Id: 3, DoubleVal: float64(3.14)},
			{Id: 4, StringVal: "this is a test"},
			{Id: 5, BytesVal: []byte{1, 2, 3, 4}},
		},
	}

	// Marshal the message into a byte array
	out, err := proto.Marshal(response)
	assert.Nil(t, err)
	assert.NotZero(t, out)

	// Then unmarshal
	unmarshalled := &Response{}
	err = proto.Unmarshal(out, unmarshalled)
	assert.Nil(t, err)

	// Make sure that the values that should be there are present
	assert.Equal(t, id, unmarshalled.Id)
	assert.Equal(t, command, unmarshalled.Command)
	assert.Equal(t, sequence, unmarshalled.Sequence)
	assert.Equal(t, uint32(1), unmarshalled.Values[0].Id)
	assert.Equal(t, int32(1234), unmarshalled.Values[0].Int32Val)

	assert.Equal(t, uint32(2), unmarshalled.Values[1].Id)
	assert.Equal(t, int64(5678), unmarshalled.Values[1].Int64Val)

	assert.Equal(t, uint32(3), unmarshalled.Values[2].Id)
	assert.Equal(t, float64(3.14), unmarshalled.Values[2].DoubleVal)

	assert.Equal(t, uint32(4), unmarshalled.Values[3].Id)
	assert.Equal(t, "this is a test", unmarshalled.Values[3].StringVal)

	assert.Equal(t, uint32(5), unmarshalled.Values[4].Id)
	assert.Equal(t, []byte{1, 2, 3, 4}, unmarshalled.Values[4].BytesVal)

	// Make sure that the values that shouldn't be there are Zero
	// values for that type
	assert.Zero(t, unmarshalled.Values[0].Int64Val)
	assert.Zero(t, unmarshalled.Values[0].DoubleVal)
	assert.Zero(t, unmarshalled.Values[0].StringVal)
	assert.Zero(t, unmarshalled.Values[0].BytesVal)
}
