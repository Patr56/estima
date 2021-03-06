package conf

import (
	"encoding/json"
	"os"
	"fmt"
)

type Ldap struct {
	Protocol string
	Host string
	Port uint32
	Dn string
}

type Database struct {
	Url	string
	User	string
	Password string
	Log bool
	Name string
}

type Auth struct {
	CookieName string
	MaxAge int
}

type Profile struct {
	Name string
	Secret string
	Ldap	Ldap
	Database Database
	Auth Auth
}

type Configuration struct {
	Active   string
	Profiles []Profile
}

func (cf Configuration) ActiveProfile ()(Profile) {
	var profile Profile
	for i:=0; i<len(cf.Profiles); i++ {
		if cf.Profiles[i].Name == cf.Active {
			profile = cf.Profiles[i]
		}
	}

	return profile
}

var config Configuration
func LoadConfig() (Profile) {
	if config.Active != "" {
		return config.ActiveProfile()
	}

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(config.ActiveProfile())
	return config.ActiveProfile()
}


