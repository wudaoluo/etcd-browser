package serverRoom

import (
	"fmt"
	"github.com/wudaoluo/etcd-browser/internal"
)

type version struct {
	ver string  //版本号
	softname string
}

var Version = &version{
	ver:		internal.VERSION,
	softname:	internal.SERVICE_NAME,
}


func (v *version) String() string{
	return fmt.Sprintln("%s version: %s", internal.SERVICE_NAME,internal.VERSION)
}

