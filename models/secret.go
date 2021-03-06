package models

import (
	"fmt"
	"reflect"
	"time"

	specV1 "github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/jinzhu/copier"
)

type SecretList struct {
	Total       int             `json:"total"`
	ListOptions *ListOptions    `json:"listOptions"`
	Items       []specV1.Secret `json:"items"`
}

func FromSecret(s *specV1.Secret) *Registry {
	if v, ok := s.Labels[specV1.SecretLabel]; !ok || v != specV1.SecretRegistry {
		return nil
	}
	res := &Registry{}
	err := copier.Copy(res, s)
	if err != nil {
		panic(fmt.Sprintf("copier exception: %s", err.Error()))
	}
	if v, ok := s.Data["address"]; ok {
		res.Address = string(v)
	}
	if v, ok := s.Data["password"]; ok {
		res.Password = string(v)
	}
	if v, ok := s.Data["username"]; ok {
		res.Username = string(v)
	}
	return res
}

func FromSecretList(s *SecretList) *RegistryList {
	res := &RegistryList{
		Total:       s.Total,
		ListOptions: s.ListOptions,
		Items:       []Registry{},
	}
	for _, sd := range s.Items {
		r := FromSecret(&sd)
		if r == nil {
			res.Total--
		} else {
			res.Items = append(res.Items, *r)
		}
	}

	return res
}

type SecretView struct {
	Name              string            `json:"name,omitempty" validate:"omitempty,resourceName"`
	Namespace         string            `json:"namespace,omitempty"`
	Data              map[string]string `json:"data,omitempty" binding:"required"`
	CreationTimestamp time.Time         `json:"createTime,omitempty"`
	UpdateTimestamp   time.Time         `json:"updateTime,omitempty"`
	Description       string            `json:"description"`
	Version           string            `json:"version,omitempty"`
}

func (s *SecretView) Equal(target *SecretView) bool {
	return reflect.DeepEqual(s.Data, target.Data) &&
		reflect.DeepEqual(s.Description, target.Description)
}

type SecretViewList struct {
	Total       int          `json:"total"`
	ListOptions *ListOptions `json:"listOptions"`
	Items       []SecretView `json:"items"`
}

func (s *SecretView) ToSecret() *specV1.Secret {
	res := &specV1.Secret{
		Labels: map[string]string{
			specV1.SecretLabel: specV1.SecretConfig,
		},
	}
	err := copier.Copy(res, s)
	if err != nil {
		panic(fmt.Sprintf("copier exception: %s", err.Error()))
	}
	res.Data = map[string][]byte{}
	for k, v := range s.Data {
		res.Data[k] = []byte(v)
	}
	return res
}

func FromSecretToView(s *specV1.Secret) *SecretView {
	if v, ok := s.Labels[specV1.SecretLabel]; !ok || v != specV1.SecretConfig {
		return nil
	}
	res := &SecretView{}
	err := copier.Copy(res, s)
	if err != nil {
		panic(fmt.Sprintf("copier exception: %s", err.Error()))
	}
	res.Data = map[string]string{}
	for k, v := range s.Data {
		res.Data[k] = string(v)
	}
	return res
}

func FromSecretListToView(s *SecretList) *SecretViewList {
	res := &SecretViewList{
		Total:       s.Total,
		ListOptions: s.ListOptions,
		Items:       []SecretView{},
	}
	for _, sd := range s.Items {
		r := FromSecretToView(&sd)

		if r == nil {
			res.Total--
		} else {
			r.Data = nil
			res.Items = append(res.Items, *r)
		}
	}
	return res
}
