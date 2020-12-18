package inmemory

import "testing"

type TestUpdater struct {
	Updater
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
	api, err := NewInMemoryClient(WithUpdater(TestUpdater{}), WithInit())
	if err != nil {
		t.Fatal(err)
	}

	if string(api.Items) != "items" {
		t.Errorf("expected %s, got %s", "items", string(api.Items))
	}

	if string(api.Monsters) != "monsters" {
		t.Errorf("expected %s, got %s", "items", string(api.Monsters))
	}

	if string(api.Prayers) != "prayers" {
		t.Errorf("expected %s, got %s", "items", string(api.Prayers))
	}
}
