package serverRoom

import (
	"fmt"
)

type version struct {
	ver string  //版本号
}

var Version = &version{
	ver:		VERSION,
}


func (v *version) String() string{
	cnf := GetConfigInstance()
	return fmt.Sprintln("%s version: %s",cnf.GetString("service_name"), VERSION)
}

