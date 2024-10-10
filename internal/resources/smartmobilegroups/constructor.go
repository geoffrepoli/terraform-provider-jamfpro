// smartcomputergroup_object.go
package smartmobilegroups

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProSmartComputerGroup constructs a ResourceMobileDeviceGroup object from the provided schema data.
func construct(d *schema.ResourceData) (*jamfpro.ResourceMobileDeviceGroup, error) {
	resource := &jamfpro.ResourceMobileDeviceGroup{
		Name:    d.Get("name").(string),
		IsSmart: true,
	}
	/*
		resource.Site = sharedschemas.ConstructSharedResourceSite(d.Get("site_id").(int))

		if v, ok := d.GetOk("criteria"); ok {
			resource.Criteria = constructMobileGroupSubsetContainerCriteria(v.([]interface{}))
		}
	*/
	resourceXML, err := xml.MarshalIndent(resource, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Computer Group '%s' to XML: %v", resource.Name, err)
	}

	log.Printf("[DEBUG] Constructed Jamf Pro Computer Group XML:\n%s\n", string(resourceXML))

	return resource, nil
}

// constructComputerGroupSubsetContainerCriteria constructs a ComputerGroupSubsetContainerCriteria object from the provided schema data.
func constructMobileGroupSubsetContainerCriteria(criteriaList []interface{}) *jamfpro.ComputerGroupSubsetContainerCriteria {
	criteria := &jamfpro.ComputerGroupSubsetContainerCriteria{
		Size:      len(criteriaList),
		Criterion: &[]jamfpro.SharedSubsetCriteria{},
	}

	for _, item := range criteriaList {
		criterionData := item.(map[string]interface{})
		criterion := jamfpro.SharedSubsetCriteria{
			Name:         criterionData["name"].(string),
			Priority:     criterionData["priority"].(int),
			AndOr:        criterionData["and_or"].(string),
			SearchType:   criterionData["search_type"].(string),
			Value:        criterionData["value"].(string),
			OpeningParen: criterionData["opening_paren"].(bool),
			ClosingParen: criterionData["closing_paren"].(bool),
		}
		*criteria.Criterion = append(*criteria.Criterion, criterion)
	}

	return criteria
}
