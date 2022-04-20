package blizzapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefaultConfig(t *testing.T) {
	expected := Config{
		BaseUrl:      "us.api.blizzard.com",
		ClientID:     "abcd-1234",
		ClientSecret: "foobar-fizzbuzz",
		Locale:       "en_US",
		Region:       "us",
		UserAgent:    "unknown golang application",
	}

	actual := Config{
		ClientID:     "abcd-1234",
		ClientSecret: "foobar-fizzbuzz",
		Region:       "us",
	}
	if err := actual.Init(); err != nil {
		t.Fatalf("Config created with only ClientID and ClientSecret failed with error %s", err.Error())
	}

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected config structure does not match actual config structure.")
	}
}

func TestMissingClientID(t *testing.T) {
	c := Config{ClientSecret: "foobar-fizzbuzz", Region: "us"}
	if err := c.Init(); err == nil {
		t.Fatalf("Config successfully created without requiring ClientID. Expected failure.")
	}
}

func TestMissingClientSecret(t *testing.T) {
	c := Config{ClientID: "foobar-fizzbuzz", Region: "us"}
	if err := c.Init(); err == nil {
		t.Fatalf("Config successfully created without requiring ClientSecret. Expected failure.")
	}
}

func TestMissingRegion(t *testing.T) {
	c := Config{ClientID: "foobar-fizzbuzz", ClientSecret: "foobar-fizzbuzz"}
	if err := c.Init(); err == nil {
		t.Fatalf("Config successfully created without requiring Region. Expected failure.")
	}
}
