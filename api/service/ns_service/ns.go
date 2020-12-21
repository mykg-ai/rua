/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package ns_service

import (
	"mykg.ai/rua/config"
	"mykg.ai/rua/domain/entity"
)

func CreateNs(name string) (entity.Namespace, error) {
	ns := entity.Namespace{Name: name}
	result := config.DB.Create(&ns)
	if result.Error != nil {
		return ns, result.Error
	}
	return ns, nil
}

func GetNs() {

}