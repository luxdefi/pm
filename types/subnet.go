package types

import "github.com/ava-labs/avalanchego/chains"

var _ Plugin = &Subnet{}

type Subnet struct {
	ID_          string              `yaml:"id"`
	Alias_       string              `yaml:"alias"`
	Homepage_    string              `yaml:"homepage"`
	Description_ string              `yaml:"description"`
	Maintainers_ []string            `yaml:"maintainers"`
	VMs_         []string            `yaml:"vms"`
	Config_      chains.SubnetConfig `yaml:"config"`
}

func (s *Subnet) ID() string {
	return s.ID_
}

func (s *Subnet) Alias() string {
	return s.Alias_
}

func (s *Subnet) Homepage() string {
	return s.Homepage_
}

func (s *Subnet) Description() string {
	return s.Description_
}

func (s *Subnet) Maintainers() []string {
	return s.Maintainers_
}