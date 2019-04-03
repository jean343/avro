package avro_test

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/stretchr/testify/assert"
)

func TestParse_InvalidType(t *testing.T) {
	schemas := []string{
		`123`,
		`{"type": 123}`,
	}

	for _, schm := range schemas {
		_, err := avro.Parse(schm)

		assert.Error(t, err)
	}
}

func TestMustParse(t *testing.T) {
	s := avro.MustParse("null")

	assert.Equal(t, avro.Null, s.Type())
}

func TestMustParse_PanicsOnError(t *testing.T) {
	assert.Panics(t, func() {
		avro.MustParse("123")
	})
}

func TestParseFiles(t *testing.T) {
	s, err := avro.ParseFiles("testdata/schema.avsc")

	assert.NoError(t, err)
	assert.Equal(t, avro.String, s.Type())
}

func TestParseFiles_FileDoesntExist(t *testing.T) {
	_, err := avro.ParseFiles("test.something")

	assert.Error(t, err)
}

func TestParseFiles_InvalidSchema(t *testing.T) {
	_, err := avro.ParseFiles("testdata/bad-schema.avsc")

	assert.Error(t, err)
}

func TestNullSchema(t *testing.T) {
	schemas := []string{
		`null`,
		`{"type":"null"}`,
	}

	for _, schm := range schemas {
		schema, err := avro.Parse(schm)

		assert.NoError(t, err)
		assert.Equal(t, avro.Null, schema.Type())
		want := [32]byte{0xf0, 0x72, 0xcb, 0xec, 0x3b, 0xf8, 0x84, 0x18, 0x71, 0xd4, 0x28, 0x42, 0x30, 0xc5, 0xe9, 0x83, 0xdc, 0x21, 0x1a, 0x56, 0x83, 0x7a, 0xed, 0x86, 0x24, 0x87, 0x14, 0x8f, 0x94, 0x7d, 0x1a, 0x1f}
		assert.Equal(t, want, schema.Fingerprint())
	}
}

func TestPrimitiveSchema(t *testing.T) {
	tests := []struct {
		schema          string
		want            avro.Type
		wantFingerprint [32]byte
	}{
		{
			schema:          "string",
			want:            avro.String,
			wantFingerprint: [32]byte{0xe9, 0xe5, 0xc1, 0xc9, 0xe4, 0xf6, 0x27, 0x73, 0x39, 0xd1, 0xbc, 0xde, 0x7, 0x33, 0xa5, 0x9b, 0xd4, 0x2f, 0x87, 0x31, 0xf4, 0x49, 0xda, 0x6d, 0xc1, 0x30, 0x10, 0xa9, 0x16, 0x93, 0xd, 0x48},
		},
		{
			schema:          `{"type":"string"}`,
			want:            avro.String,
			wantFingerprint: [32]byte{0xe9, 0xe5, 0xc1, 0xc9, 0xe4, 0xf6, 0x27, 0x73, 0x39, 0xd1, 0xbc, 0xde, 0x7, 0x33, 0xa5, 0x9b, 0xd4, 0x2f, 0x87, 0x31, 0xf4, 0x49, 0xda, 0x6d, 0xc1, 0x30, 0x10, 0xa9, 0x16, 0x93, 0xd, 0x48},
		},
		{
			schema:          "bytes",
			want:            avro.Bytes,
			wantFingerprint: [32]byte{0x9a, 0xe5, 0x7, 0xa9, 0xdd, 0x39, 0xee, 0x5b, 0x7c, 0x7e, 0x28, 0x5d, 0xa2, 0xc0, 0x84, 0x65, 0x21, 0xc8, 0xae, 0x8d, 0x80, 0xfe, 0xea, 0xe5, 0x50, 0x4e, 0xc, 0x98, 0x1d, 0x53, 0xf5, 0xfa},
		},
		{
			schema:          `{"type":"bytes"}`,
			want:            avro.Bytes,
			wantFingerprint: [32]byte{0x9a, 0xe5, 0x7, 0xa9, 0xdd, 0x39, 0xee, 0x5b, 0x7c, 0x7e, 0x28, 0x5d, 0xa2, 0xc0, 0x84, 0x65, 0x21, 0xc8, 0xae, 0x8d, 0x80, 0xfe, 0xea, 0xe5, 0x50, 0x4e, 0xc, 0x98, 0x1d, 0x53, 0xf5, 0xfa},
		},
		{
			schema:          "int",
			want:            avro.Int,
			wantFingerprint: [32]byte{0x3f, 0x2b, 0x87, 0xa9, 0xfe, 0x7c, 0xc9, 0xb1, 0x38, 0x35, 0x59, 0x8c, 0x39, 0x81, 0xcd, 0x45, 0xe3, 0xe3, 0x55, 0x30, 0x9e, 0x50, 0x90, 0xaa, 0x9, 0x33, 0xd7, 0xbe, 0xcb, 0x6f, 0xba, 0x45},
		},
		{
			schema:          `{"type":"int"}`,
			want:            avro.Int,
			wantFingerprint: [32]byte{0x3f, 0x2b, 0x87, 0xa9, 0xfe, 0x7c, 0xc9, 0xb1, 0x38, 0x35, 0x59, 0x8c, 0x39, 0x81, 0xcd, 0x45, 0xe3, 0xe3, 0x55, 0x30, 0x9e, 0x50, 0x90, 0xaa, 0x9, 0x33, 0xd7, 0xbe, 0xcb, 0x6f, 0xba, 0x45},
		},
		{
			schema:          "long",
			want:            avro.Long,
			wantFingerprint: [32]byte{0xc3, 0x2c, 0x49, 0x7d, 0xf6, 0x73, 0xc, 0x97, 0xfa, 0x7, 0x36, 0x2a, 0xa5, 0x2, 0x3f, 0x37, 0xd4, 0x9a, 0x2, 0x7e, 0xc4, 0x52, 0x36, 0x7, 0x78, 0x11, 0x4c, 0xf4, 0x27, 0x96, 0x5a, 0xdd},
		},
		{
			schema:          `{"type":"long"}`,
			want:            avro.Long,
			wantFingerprint: [32]byte{0xc3, 0x2c, 0x49, 0x7d, 0xf6, 0x73, 0xc, 0x97, 0xfa, 0x7, 0x36, 0x2a, 0xa5, 0x2, 0x3f, 0x37, 0xd4, 0x9a, 0x2, 0x7e, 0xc4, 0x52, 0x36, 0x7, 0x78, 0x11, 0x4c, 0xf4, 0x27, 0x96, 0x5a, 0xdd},
		},
		{
			schema:          "float",
			want:            avro.Float,
			wantFingerprint: [32]byte{0x1e, 0x71, 0xf9, 0xec, 0x5, 0x1d, 0x66, 0x3f, 0x56, 0xb0, 0xd8, 0xe1, 0xfc, 0x84, 0xd7, 0x1a, 0xa5, 0x6c, 0xcf, 0xe9, 0xfa, 0x93, 0xaa, 0x20, 0xd1, 0x5, 0x47, 0xa7, 0xab, 0xeb, 0x5c, 0xc0},
		},
		{
			schema:          `{"type":"float"}`,
			want:            avro.Float,
			wantFingerprint: [32]byte{0x1e, 0x71, 0xf9, 0xec, 0x5, 0x1d, 0x66, 0x3f, 0x56, 0xb0, 0xd8, 0xe1, 0xfc, 0x84, 0xd7, 0x1a, 0xa5, 0x6c, 0xcf, 0xe9, 0xfa, 0x93, 0xaa, 0x20, 0xd1, 0x5, 0x47, 0xa7, 0xab, 0xeb, 0x5c, 0xc0},
		},
		{
			schema:          "double",
			want:            avro.Double,
			wantFingerprint: [32]byte{0x73, 0xa, 0x9a, 0x8c, 0x61, 0x16, 0x81, 0xd7, 0xee, 0xf4, 0x42, 0xe0, 0x3c, 0x16, 0xc7, 0xd, 0x13, 0xbc, 0xa3, 0xeb, 0x8b, 0x97, 0x7b, 0xb4, 0x3, 0xea, 0xff, 0x52, 0x17, 0x6a, 0xf2, 0x54},
		},
		{
			schema:          `{"type":"double"}`,
			want:            avro.Double,
			wantFingerprint: [32]byte{0x73, 0xa, 0x9a, 0x8c, 0x61, 0x16, 0x81, 0xd7, 0xee, 0xf4, 0x42, 0xe0, 0x3c, 0x16, 0xc7, 0xd, 0x13, 0xbc, 0xa3, 0xeb, 0x8b, 0x97, 0x7b, 0xb4, 0x3, 0xea, 0xff, 0x52, 0x17, 0x6a, 0xf2, 0x54},
		},
		{
			schema:          "boolean",
			want:            avro.Boolean,
			wantFingerprint: [32]byte{0xa5, 0xb0, 0x31, 0xab, 0x62, 0xbc, 0x41, 0x6d, 0x72, 0xc, 0x4, 0x10, 0xd8, 0x2, 0xea, 0x46, 0xb9, 0x10, 0xc4, 0xfb, 0xe8, 0x5c, 0x50, 0xa9, 0x46, 0xcc, 0xc6, 0x58, 0xb7, 0x4e, 0x67, 0x7e},
		},
		{
			schema:          `{"type":"boolean"}`,
			want:            avro.Boolean,
			wantFingerprint: [32]byte{0xa5, 0xb0, 0x31, 0xab, 0x62, 0xbc, 0x41, 0x6d, 0x72, 0xc, 0x4, 0x10, 0xd8, 0x2, 0xea, 0x46, 0xb9, 0x10, 0xc4, 0xfb, 0xe8, 0x5c, 0x50, 0xa9, 0x46, 0xcc, 0xc6, 0x58, 0xb7, 0x4e, 0x67, 0x7e},
		},
	}

	for _, tt := range tests {
		t.Run(tt.schema, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, s.Type())
			assert.Equal(t, tt.wantFingerprint, s.Fingerprint())
		})
	}
}

func TestRecordSchema(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			name:    "Valid",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "doc": "docs", "fields":[{"name": "field", "type": "int"}]}`,
			wantErr: false,
		},
		{
			name:    "Full Name",
			schema:  `{"type":"record", "name":"org.apache.avro.test", "doc": "docs", "fields":[{"name": "field", "type": "int"}]}`,
			wantErr: false,
		},
		{
			name:    "Invalid Name",
			schema:  `{"type":"record", "name":"test+", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "Empty Name",
			schema:  `{"type":"record", "name":"", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "No Name",
			schema:  `{"type":"record", "namespace": "org.apache.avro", "fields":[{"name": "intField", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Namespace",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro+", "fields":[{"name": "field", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "Empty Namespace",
			schema:  `{"type":"record", "name":"test", "namespace": "", "fields":[{"name": "intField", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "No Fields",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Field Type",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":["test"]}`,
			wantErr: true,
		},
		{
			name:    "No Field Name",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Field Name",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field+", "type": "int"}]}`,
			wantErr: true,
		},
		{
			name:    "No Field Type",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field"}]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Field Type",
			schema:  `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "blah"}]}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Record, s.Type())
		})
	}
}

func TestRecordSchema_Default(t *testing.T) {
	tests := []struct {
		name   string
		schema string
		want   interface{}
	}{
		{
			name:   "Normal",
			schema: `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "string", "default": "test"}]}`,
			want:   "test",
		},
		{
			name:   "Int",
			schema: `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "int", "default": 1}]}`,
			want:   int32(1),
		},
		{
			name:   "Long",
			schema: `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "long", "default": 1}]}`,
			want:   int64(1),
		},
		{
			name:   "Float",
			schema: `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "float", "default": 1}]}`,
			want:   float32(1),
		},
		{
			name:   "Double",
			schema: `{"type":"record", "name":"test", "namespace": "org.apache.avro", "fields":[{"name": "field", "type": "double", "default": 1}]}`,
			want:   float64(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, s.(*avro.RecordSchema).Fields()[0].Default())
		})
	}
}

func TestRecordSchema_WithReference(t *testing.T) {
	schm := `
{
   "type": "record",
   "name": "valid_name",
   "namespace": "org.apache.avro",
   "fields": [
       {
           "name": "intField",
           "type": "int"
       },
       {
           "name": "Ref",
           "type": "valid_name"
       }
   ]
}
`

	s, err := avro.Parse(schm)

	assert.NoError(t, err)
	assert.Equal(t, avro.Record, s.Type())
	assert.Equal(t, avro.Ref, s.(*avro.RecordSchema).Fields()[1].Type().Type())
	assert.Equal(t, s.Fingerprint(), s.(*avro.RecordSchema).Fields()[1].Type().Fingerprint())
}

func TestEnumSchema(t *testing.T) {
	tests := []struct {
		name     string
		schema   string
		wantName string
		wantErr  bool
	}{
		{
			name:     "Valid",
			schema:   `{"type":"enum", "name":"test", "namespace": "org.apache.avro", "symbols":["TEST"]}`,
			wantName: "org.apache.avro.test",
			wantErr:  false,
		},
		{
			name:    "Invalid Name",
			schema:  `{"type":"enum", "name":"test+", "namespace": "org.apache.avro", "symbols":["TEST"]}`,
			wantErr: true,
		},
		{
			name:    "Empty Name",
			schema:  `{"type":"enum", "name":"", "namespace": "org.apache.avro", "symbols":["TEST"]}`,
			wantErr: true,
		},
		{
			name:    "No Name",
			schema:  `{"type":"enum", "namespace": "org.apache.avro", "symbols":["TEST"]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Namespace",
			schema:  `{"type":"enum", "name":"test", "namespace": "org.apache.avro+", "symbols":["TEST"]}`,
			wantErr: true,
		},
		{
			name:    "Empty Namespace",
			schema:  `{"type":"enum", "name":"test", "namespace": "", "symbols":["TEST"]}`,
			wantErr: true,
		},
		{
			name:    "No Symbols",
			schema:  `{"type":"enum", "name":"test", "namespace": "org.apache.avro"}`,
			wantErr: true,
		},
		{
			name:    "Empty Symbols",
			schema:  `{"type":"enum", "name":"test", "namespace": "org.apache.avro", "symbols":[]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Symbol",
			schema:  `{"type":"enum", "name":"test", "namespace": "org.apache.avro", "symbols":["TEST+"]}`,
			wantErr: true,
		},
		{
			name:    "Invalid Symbol Type",
			schema:  `{"type":"enum", "name":"test", "namespace": "org.apache.avro", "symbols":[1]}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Enum, schema.Type())
			named := schema.(avro.NamedSchema)
			assert.Equal(t, tt.wantName, named.Name())
		})
	}
}

func TestArraySchema(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			name:    "Valid",
			schema:  `{"type":"array", "items": "int"}`,
			wantErr: false,
		},
		{
			name:    "No Items",
			schema:  `{"type":"array"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Items Type",
			schema:  `{"type":"array", "items": "blah"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Array, s.Type())
		})
	}
}

func TestMapSchema(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			name:    "Valid",
			schema:  `{"type":"map", "values": "int"}`,
			wantErr: false,
		},
		{
			name:    "No Values",
			schema:  `{"type":"map"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Values Type",
			schema:  `{"type":"map", "values": "blah"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Map, s.Type())
		})
	}
}

func TestUnionSchema(t *testing.T) {
	tests := []struct {
		name            string
		schema          string
		wantFingerprint [32]byte
		wantErr         bool
	}{
		{
			name:            "Valid Simple",
			schema:          `["null", "int"]`,
			wantFingerprint: [32]byte{0xb4, 0x94, 0x95, 0xc5, 0xb1, 0xc2, 0x6f, 0x4, 0x89, 0x6a, 0x5f, 0x68, 0x65, 0xf, 0xe2, 0xb7, 0x64, 0x23, 0x62, 0xc3, 0x41, 0x98, 0xd6, 0xbc, 0x74, 0x65, 0xa1, 0xd9, 0xf7, 0xe1, 0xaf, 0xce},
			wantErr:         false,
		},
		{
			name:            "Valid Complex",
			schema:          `{"type":["null", "int"]}`,
			wantFingerprint: [32]byte{0xb4, 0x94, 0x95, 0xc5, 0xb1, 0xc2, 0x6f, 0x4, 0x89, 0x6a, 0x5f, 0x68, 0x65, 0xf, 0xe2, 0xb7, 0x64, 0x23, 0x62, 0xc3, 0x41, 0x98, 0xd6, 0xbc, 0x74, 0x65, 0xa1, 0xd9, 0xf7, 0xe1, 0xaf, 0xce},
			wantErr:         false,
		},
		{
			name:    "Invalid Type",
			schema:  `["null", "blah"]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Union, s.Type())
			assert.Equal(t, tt.wantFingerprint, s.Fingerprint())
		})
	}
}

func TestFixedSchema(t *testing.T) {
	tests := []struct {
		name            string
		schema          string
		wantName        string
		wantFingerprint [32]byte
		wantErr         bool
	}{
		{
			name:            "Valid",
			schema:          `{"type":"fixed", "name":"test", "namespace": "org.apache.avro", "size": 12}`,
			wantName:        "org.apache.avro.test",
			wantFingerprint: [32]byte{0xa8, 0x13, 0xfa, 0xb4, 0xf, 0xd7, 0xe3, 0xc9, 0x3a, 0x98, 0x77, 0x24, 0xaf, 0xa9, 0x36, 0xe6, 0xe9, 0x53, 0xa9, 0x1c, 0x10, 0x70, 0xfe, 0x4e, 0x13, 0x2a, 0x7c, 0x51, 0x6, 0x5f, 0xa4, 0xbc},
			wantErr:         false,
		},
		{
			name:    "Invalid Name",
			schema:  `{"type":"fixed", "name":"test+", "namespace": "org.apache.avro", "size": 12}`,
			wantErr: true,
		},
		{
			name:    "Empty Name",
			schema:  `{"type":"fixed", "name":"", "namespace": "org.apache.avro", "size": 12}`,
			wantErr: true,
		},
		{
			name:    "No Name",
			schema:  `{"type":"fixed", "namespace": "org.apache.avro", "size": 12}`,
			wantErr: true,
		},
		{
			name:    "Invalid Namespace",
			schema:  `{"type":"fixed", "name":"test", "namespace": "org.apache.avro+", "size": 12}`,
			wantErr: true,
		},
		{
			name:    "Empty Namespace",
			schema:  `{"type":"fixed", "name":"test", "namespace": "", "size": 12}`,
			wantErr: true,
		},
		{
			name:    "No Size",
			schema:  `{"type":"fixed", "name":"test", "namespace": "org.apache.avro"}`,
			wantErr: true,
		},
		{
			name:    "Invalid Size Type",
			schema:  `{"type":"fixed", "name":"test", "namespace": "org.apache.avro", "size": "test"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema, err := avro.Parse(tt.schema)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, avro.Fixed, schema.Type())
			named := schema.(avro.NamedSchema)
			assert.Equal(t, tt.wantName, named.Name())
			assert.Equal(t, tt.wantFingerprint, named.Fingerprint())
		})
	}
}

func TestSchema_Interop(t *testing.T) {
	schm := `
{
   "type": "record",
   "name": "Interop",
   "namespace": "org.apache.avro",
   "fields": [
       {
           "name": "intField",
           "type": "int"
       },
       {
           "name": "longField",
           "type": "long"
       },
       {
           "name": "stringField",
           "type": "string"
       },
       {
           "name": "boolField",
           "type": "boolean"
       },
       {
           "name": "floatField",
           "type": "float"
       },
       {
           "name": "doubleField",
           "type": "double"
       },
       {
           "name": "bytesField",
           "type": "bytes"
       },
       {
           "name": "nullField",
           "type": "null"
       },
       {
           "name": "arrayField",
           "type": {
               "type": "array",
               "items": "double"
           }
       },
       {
           "name": "mapField",
           "type": {
               "type": "map",
               "values": {
                   "type": "record",
                   "name": "Foo",
                   "fields": [
                       {
                           "name": "label",
                           "type": "string"
                       }
                   ]
               }
           }
       },
       {
           "name": "unionField",
           "type": [
               "boolean",
               "double",
               {
                   "type": "array",
                   "items": "bytes"
               }
           ]
       },
       {
           "name": "enumField",
           "type": {
               "type": "enum",
               "name": "Kind",
               "symbols": [
                   "A",
                   "B",
                   "C"
               ]
           }
       },
       {
           "name": "fixedField",
           "type": {
               "type": "fixed",
               "name": "MD5",
               "size": 16
           }
       },
       {
           "name": "recordField",
           "type": {
               "type": "record",
               "name": "Node",
               "fields": [
                   {
                       "name": "label",
                       "type": "string"
                   },
                   {
                       "name": "child",
                       "type": {"type": "org.apache.avro.Node"}
                   },
                   {
                       "name": "children",
                       "type": {
                           "type": "array",
                           "items": "Node"
                       }
                   }
               ]
           }
       }
   ]
}`

	_, err := avro.Parse(schm)

	assert.NoError(t, err)
}