package environment

type EnvironmentSecrets struct {
	ID            string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EnvironmentID string `gorm:"type:uuid;not null"`
	Key           string `gorm:"type:varchar(255);not null"`
	Value         string `gorm:"type:text;not null"` //This should be hashed and salted
}
