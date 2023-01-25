// Code generated by transpose_ints.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

// when we upgrade to support go1.18, this can be massively simplified by using
// Go Generics, but since we aren't supporting go1.18 yet, I didn't want to use
// them here so we can maintain the backwards compatibility.

func transposeInt8Int8(src []int8, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeInt8Uint8(src []int8, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeInt8Int16(src []int8, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeInt8Uint16(src []int8, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeInt8Int32(src []int8, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeInt8Uint32(src []int8, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeInt8Int64(src []int8, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeInt8Uint64(src []int8, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeUint8Int8(src []uint8, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeUint8Uint8(src []uint8, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeUint8Int16(src []uint8, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeUint8Uint16(src []uint8, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeUint8Int32(src []uint8, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeUint8Uint32(src []uint8, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeUint8Int64(src []uint8, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeUint8Uint64(src []uint8, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeInt16Int8(src []int16, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeInt16Uint8(src []int16, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeInt16Int16(src []int16, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeInt16Uint16(src []int16, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeInt16Int32(src []int16, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeInt16Uint32(src []int16, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeInt16Int64(src []int16, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeInt16Uint64(src []int16, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeUint16Int8(src []uint16, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeUint16Uint8(src []uint16, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeUint16Int16(src []uint16, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeUint16Uint16(src []uint16, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeUint16Int32(src []uint16, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeUint16Uint32(src []uint16, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeUint16Int64(src []uint16, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeUint16Uint64(src []uint16, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeInt32Int8(src []int32, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeInt32Uint8(src []int32, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeInt32Int16(src []int32, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeInt32Uint16(src []int32, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeInt32Int32(src []int32, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeInt32Uint32(src []int32, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeInt32Int64(src []int32, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeInt32Uint64(src []int32, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeUint32Int8(src []uint32, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeUint32Uint8(src []uint32, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeUint32Int16(src []uint32, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeUint32Uint16(src []uint32, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeUint32Int32(src []uint32, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeUint32Uint32(src []uint32, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeUint32Int64(src []uint32, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeUint32Uint64(src []uint32, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeInt64Int8(src []int64, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeInt64Uint8(src []int64, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeInt64Int16(src []int64, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeInt64Uint16(src []int64, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeInt64Int32(src []int64, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeInt64Uint32(src []int64, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeInt64Int64(src []int64, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeInt64Uint64(src []int64, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}

func transposeUint64Int8(src []uint64, dest []int8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int8(transposeMap[s])
	}
}

func transposeUint64Uint8(src []uint64, dest []uint8, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint8(transposeMap[s])
	}
}

func transposeUint64Int16(src []uint64, dest []int16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int16(transposeMap[s])
	}
}

func transposeUint64Uint16(src []uint64, dest []uint16, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint16(transposeMap[s])
	}
}

func transposeUint64Int32(src []uint64, dest []int32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int32(transposeMap[s])
	}
}

func transposeUint64Uint32(src []uint64, dest []uint32, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint32(transposeMap[s])
	}
}

func transposeUint64Int64(src []uint64, dest []int64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = int64(transposeMap[s])
	}
}

func transposeUint64Uint64(src []uint64, dest []uint64, transposeMap []int32) {
	for i, s := range src {
		dest[i] = uint64(transposeMap[s])
	}
}