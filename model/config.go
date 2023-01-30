package model

type Config struct {
	Package    string `json:"package"`
	Connection *struct {
		MaxIdleConnections *uint   `json:"maxIdleConnections"`
		ConnMaxLifetime    *string `json:"connMaxLifetime"`
		MaxOpenConnections *uint   `json:"maxOpenConnections"`
	} `json:"connection,omitempty"`
}
