package utils

type Configuration struct {
	Database DatabaseSetting
}

type DatabaseSetting struct {
	Url string
}
