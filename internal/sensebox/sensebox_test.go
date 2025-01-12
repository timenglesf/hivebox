package sensebox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeseboxData(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		expected      SenseboxJSON
		errorExpected bool
	}{
		{
			name:     "dummy sensebox 1",
			id:       SENSEBOX_DUMMY_ID_1,
			expected: dummySensebox1,
		},
		{
			name:     "dummy sensebox 2",
			id:       SENSEBOX_DUMMY_ID_2,
			expected: dummySensebox2,
		},
		{
			name:     "dummy sensebox 3",
			id:       SENSEBOX_DUMMY_ID_3,
			expected: dummySensebox3,
		},
		{
			name:          "invalid sensebox id",
			errorExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sbs := &SenseboxServiceMock{}
			data, err := sbs.GetSenseboxData(tt.id)
			if err != nil {
				if tt.errorExpected {
					return
				}
				t.Fatal(err)
			}
			assert.Equal(t, tt.expected, data)
		})
	}
}
