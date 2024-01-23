package environment

type EnvironmentConfiguration struct {
	ID              string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EnableWaitTimer bool   `gorm:"type:boolean;default:false"`
	WaitTimer       int    `gorm:"type:integer;default:0"`
}
