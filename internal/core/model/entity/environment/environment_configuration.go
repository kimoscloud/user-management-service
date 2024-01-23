package environment

type EnvironmentConfiguration struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
}
