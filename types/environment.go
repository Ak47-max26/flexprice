package types

// EnvironmentType defines the type of environment.
type EnvironmentType string

const (
	EnvironmentDevelopment EnvironmentType = "development"
	EnvironmentProduction  EnvironmentType = "production"
)

func (e EnvironmentType) String() string {
	return string(e)
}

func (e EnvironmentType) DisplayTitle() string {
	if e == EnvironmentDevelopment {
		return "Sandbox"
	}

	return "Production"
}
