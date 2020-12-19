package inmemory

import "testing"

type TestUpdater struct {
	source
}

func (tu TestUpdater) Items() ([]byte, error) {
	return []byte("items"), nil
}

func (tu TestUpdater) Monsters() ([]byte, error) {
	return []byte("monsters"), nil
}

func (tu TestUpdater) Prayers() ([]byte, error) {
	return []byte("prayers"), nil
}

func Test_WithInit(t *testing.T) {
	api := NewAPI()
	api.RunOptions(WithSource(TestUpdater{}), WithInit())

	if string(api.items) != "items" {
		t.Errorf("expected %s, got %s", "items", string(api.items))
	}

	if string(api.monsters) != "monsters" {
		t.Errorf("expected %s, got %s", "items", string(api.monsters))
	}

	if string(api.prayers) != "prayers" {
		t.Errorf("expected %s, got %s", "items", string(api.prayers))
	}
}
