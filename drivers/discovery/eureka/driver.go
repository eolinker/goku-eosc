package eureka

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/eolinker/eosc"
	"github.com/eolinker/apinto/discovery"
)

const (
	driverName = "eureka"
)

//driver 实现github.com/eolinker/eosc.eosc.IProfessionDriver接口
type driver struct {
	profession string
	name       string
	driver     string
	label      string
	desc       string
	configType reflect.Type
}

//ConfigType 返回eureka驱动配置的反射类型
func (d *driver) ConfigType() reflect.Type {
	return d.configType
}

//Create 创建eureka驱动实例
func (d *driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]interface{}) (eosc.IWorker, error) {
	cfg, ok := v.(*Config)
	if !ok {
		return nil, fmt.Errorf("need %s,now %s", eosc.TypeNameOf((*Config)(nil)), eosc.TypeNameOf(v))
	}
	return &eureka{
		id:       id,
		name:     name,
		client:   newClient(cfg.getAddress(), cfg.getParams()),
		nodes:    discovery.NewNodesData(),
		services: discovery.NewServices(),
		locker:   sync.RWMutex{},
	}, nil
}
