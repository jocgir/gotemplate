package customxml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshaling(t *testing.T) {
	type d = map[string]interface{}
	type l = []interface{}
	type x int
	tests := []struct {
		name   string
		pretty bool
		data   interface{}
		want   string
		err    error
	}{
		{"Empty", false, nil, "", nil},
		{"Empty String", false, "", "<string/>", nil},
		{"Integer", false, 10, "<int>10</int>", nil},
		{"Integer x", false, x(10), "<x>10</x>", nil},
		{"UInt8", false, uint8(8), "<uint8>8</uint8>", nil},
		{"String", false, "hello", "<string>hello</string>", nil},
		{"List", false, []int{1, 2, 3}, "<int>1</int><int>2</int><int>3</int>", nil},
		{"List2", false, xmlList{"Hello", "world", "!", 2019}, "<string>Hello</string><string>world</string><string>!</string><int>2019</int>", nil},
		{"Map", false,
			d{
				"int": 1,
				"str": "Hello",
				"map": d{
					"str":   "world!",
					"int64": int64(2019),
					"list": l{
						1,
						3.14,
						"Hello",
					},
				},
			},
			`<xmlDict><Key name="int" type="int" value="1"/>1</Key><Key name="map"><xmlDict><Key name="int64"><int64>2019</int64></Key><Key name="list"><int>1</int><float64>3.14</float64><string>Hello</string></Key><Key name="str"><string>world!</string></Key></xmlDict></Key><Key name="str"><string>Hello</string></Key></xmlDict>`,
			nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Marshal(tt.data)
			assert.Equal(t, tt.want, string(out))
			if tt.err == nil {
				assert.NoError(t, err)
				var value interface{}
				err = Unmarshal(out, &value)
				assert.NoError(t, err, "Unmarshal")
				assert.Equal(t, tt.data, value)
			} else {
				assert.EqualError(t, err, tt.err.Error())
			}
		})
	}
}
