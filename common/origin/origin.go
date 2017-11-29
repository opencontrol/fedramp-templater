package origin

import (
	"github.com/opencontrol/fedramp-templater/common/source"
	"gopkg.in/fatih/set.v0"
	"log"
	"reflect"
	"strings"
)

// Key is a unique value that represents the different types of possible control origination value.
type Key uint8

// Origination prefixes.
const (
	NoOrigin Key = iota
	ServiceProviderCorporateOrigination
	ServiceProviderSystemSpecificOrigination
	ServiceProviderHybridOrigination
	ConfiguredByCustomerOrigination
	ProvidedByCustomerOrigination
	SharedOrigination
	InheritedOrigination
)

// SrcMapping is a data structure that represents the text for a particular control origination in a particular source.
type SrcMapping map[source.Source]string

// IsDocMappingASubstrOf is wrapper that checks if the input string contains the SSP mapping.
// This is useful because the input string value may have extra characters so we can't do a == (equal to) check.
func (o SrcMapping) IsDocMappingASubstrOf(value string) bool {
	return strings.Contains(value, o[source.SSP])
}

// IsYAMLMappingEqualTo is a wrapper that checks if the input string equals to the YAML mapping.
func (o SrcMapping) IsYAMLMappingEqualTo(value string) bool {
	return value == o[source.YAML]
}

// GetSourceMappings returns a mapping of each control origination to their respective sources.
func GetSourceMappings() map[Key]SrcMapping {
	return map[Key]SrcMapping{
		ServiceProviderCorporateOrigination: {
			source.YAML: "service_provider_corporate",
			source.SSP:  "Service Provider Corporate",
		},
		ServiceProviderSystemSpecificOrigination: {
			source.YAML: "service_provider_system_specific",
			source.SSP:  "Service Provider System Specific",
		},
		ServiceProviderHybridOrigination: {
			source.YAML: "hybrid",
			source.SSP:  "Service Provider Hybrid",
		},
		ConfiguredByCustomerOrigination: {
			source.YAML: "customer_configured",
			source.SSP:  "Configured by Customer",
		},
		ProvidedByCustomerOrigination: {
			source.YAML: "customer_provided",
			source.SSP:  "Provided by Customer",
		},
		SharedOrigination: {
			source.YAML: "shared",
			source.SSP:  "Shared",
		},
		InheritedOrigination: {
			source.YAML: "inherited",
			source.SSP:  "Inherited",
		},
	}
}

// ConvertSetToKeys will convert the set, which each value is of type interface{} to the Key.
func ConvertSetToKeys(s set.Interface) []Key {
	keys := []Key{}
	for _, item := range s.List() {
		key, isType := item.(Key)
		if isType {
			keys = append(keys, key)
		} else {
			log.Printf("Unable to use value as ControlOrigin 'Key' Type: %v. Value: %v.\n",
				reflect.TypeOf(item), item)
		}
	}
	return keys
}
