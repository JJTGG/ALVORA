package store

import "testing"

func TestApplySchemaRejectsNilDatabase(t *testing.T) {
	if err := ApplySchema(nil); err != ErrNilDB {
		t.Fatalf("ApplySchema(nil) error = %v, want %v", err, ErrNilDB)
	}
}

func TestSchemaContainsResourceTable(t *testing.T) {
	for _, want := range []string{"CREATE TABLE IF NOT EXISTS resources", "organization_id", "desired_spec", "observed_status"} {
		if !contains(Schema, want) {
			t.Fatalf("Schema does not contain %q", want)
		}
	}
}

func contains(value, fragment string) bool {
	return len(value) >= len(fragment) && stringIndex(value, fragment) >= 0
}

func stringIndex(value, fragment string) int {
	for i := 0; i+len(fragment) <= len(value); i++ {
		if value[i:i+len(fragment)] == fragment {
			return i
		}
	}
	return -1
}
