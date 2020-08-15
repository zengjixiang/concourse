package resources

import (
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
)

// Configurator loops over all resources and resource types and updates the
// resource config (and scope) on each resource.
//
// Resource configs are only updated once they have been successfully checked.
//
// Resources are simply skipped if any of their type hierarchy is missing a
// version.
type Configurator struct {
}

func (thing Configurator) Run() error {
	var resources db.Resources
	var resourceTypes db.ResourceTypes

	vrts := resourceTypes.Deserialize()

	for _, resource := range resources {
		// XXX: vars
		source, err := creds.NewSource(nil, resource.Source()).Evaluate()
		if err != nil {
			return err
		}

		_, err = resource.SetResourceConfig(source, vrts)
		if err != nil {
			return err
		}
	}

	return nil
}
