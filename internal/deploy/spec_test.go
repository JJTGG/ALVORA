package deploy

import "testing"

func TestSpecValidate(t *testing.T) {
	valid := Spec{Application: "demo", SourceDir: "/tmp/demo", StartCommand: "./demo", Port: 8080}
	if err := valid.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	tests := []struct {
		name string
		edit func(*Spec)
	}{
		{"missing application", func(s *Spec) { s.Application = "" }},
		{"missing source", func(s *Spec) { s.SourceDir = "" }},
		{"relative source", func(s *Spec) { s.SourceDir = "demo" }},
		{"missing command", func(s *Spec) { s.StartCommand = "" }},
		{"invalid port", func(s *Spec) { s.Port = 0 }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := valid
			tt.edit(&value)
			if err := value.Validate(); err == nil {
				t.Fatal("Validate() error = nil, want validation error")
			}
		})
	}
}

func TestSpecCommand(t *testing.T) {
	value := Spec{Application: "demo", SourceDir: "/tmp/demo", StartCommand: "./demo", Args: []string{"--port", "8080"}, Port: 8080}
	command, err := value.Command()
	if err != nil {
		t.Fatalf("Command() error = %v", err)
	}
	if len(command) != 3 || command[0] != "./demo" || command[1] != "--port" || command[2] != "8080" {
		t.Fatalf("Command() = %#v, want executable and args", command)
	}
}
