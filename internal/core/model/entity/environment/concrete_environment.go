package environment

// You share the codebase with the other concrete environments, but the information about the machine, datadog, and other things that are used it's not the same
type ConcreteEnvironment struct {
	ID            string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EnvironmentID string `gorm:"type:uuid;not null"`
}
