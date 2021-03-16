/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package ns_service

import (
	"mykg.ai/rua/config"
	"mykg.ai/rua/domain/entity"
)

func CreateNamespace(namespace *entity.Namespace) (entity.Namespace, error) {
	result := config.DB.Create(&namespace)
	if result.Error != nil {
		return *namespace, result.Error
	}
	return *namespace, nil
}

func FindNamespaces() ([]entity.Namespace, error) {
	var namespaces []entity.Namespace
	err := config.DB.Find(&namespaces).Error
	if err != nil {
		return []entity.Namespace{}, err
	}
	return namespaces, nil
}

func FindNamespace(id uint64) (entity.Namespace, error) {
	var namespace entity.Namespace
	err := config.DB.Where("id = ?", id).First(&namespace).Error
	if err != nil {
		return entity.Namespace{}, err
	}
	return namespace, nil
}

func DeleteNamespace(id uint64) error {
	var namespace entity.Namespace
	err := config.DB.Where("id = ?", id).Delete(&namespace).Error
	if err != nil {
		return err
	}
	return nil
}
