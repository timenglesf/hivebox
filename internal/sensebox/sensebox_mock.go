package sensebox

import "fmt"

type SenseboxServiceMock struct{}

const (
	SENSEBOX_DUMMY_ID_1 = "dummy_id_1"
	SENSEBOX_DUMMY_ID_2 = "dummy_id_2"
	SENSEBOX_DUMMY_ID_3 = "dummy_id_3"
)

func (sbs *SenseboxServiceMock) GetSenseboxData(id string) (SenseboxJSON, error) {
	switch id {
	case SENSEBOX_DUMMY_ID_1:
		return dummySensebox1, nil
	case SENSEBOX_DUMMY_ID_2:
		return dummySensebox2, nil
	case SENSEBOX_DUMMY_ID_3:
		return dummySensebox3, nil
	default:
		return SenseboxJSON{}, fmt.Errorf("sensebox with ID %s not found", id)
	}
}
