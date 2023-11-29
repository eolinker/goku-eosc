package extra_params

import (
	"github.com/eolinker/eosc"

	"github.com/eolinker/apinto/drivers"
)

const (
	Name = "extra_params"
)

func Register(register eosc.IExtenderDriverRegister) {
	register.RegisterExtenderDriver(Name, NewFactory())
}
func NewFactory() eosc.IExtenderDriverFactory {
	return drivers.NewFactory[Config](Create)
}
