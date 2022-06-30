package data

import "testing"

func TestValidate(t *testing.T) {
	p := &Product{
		Name:  "AD",
		Price: 1,
		SKU:   "ab-ab-ab",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateFail(t *testing.T) {
	p := &Product{}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
