package pkg

import (
	"gopkg.in/yaml.v2"
)

/*Struct for policy file*/
type Policy struct {
	APIVersion string                 `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Kind       string                 `json:"kind,omitempty" yaml:"kind,omitempty"`
	Metadata   map[string]string      `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec       map[string]interface{} `json:"spec,omitempty" yaml:"spec,omitempty"`
}

/*get Policy Spec in String format*/
func (p *Policy) GetSpec() string {
	data, _ := yaml.Marshal(p.Spec)

	return string(data)
}

/*get Policy Metadata in String format*/
func (p *Policy) GetMetadata() string {
	data, _ := yaml.Marshal(p.Metadata)
	return string(data)
}

/*Creates new Policy struct from given data*/
func NewPolicyFrom(data []byte) Policy {
	p := Policy{}
	yaml.Unmarshal(data, &p)
	return p
}
