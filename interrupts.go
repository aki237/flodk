package flodk

import "fmt"

// InterruptID is used to identify the interrupt against the Node which
// threw the interrupt.
type InterruptID struct {
	NodeID string `json:"node_id"`
	ID     string `json:"id"`
}

// String returns the string representation of the interrupt.
func (i InterruptID) String() string {
	return i.NodeID + ":" + i.ID
}

// RequirementTypes defines the type of the requirement.
type RequirementTypes string

const (
	// Enum requirements can only choose values from the provided suggestions.
	Enum RequirementTypes = "enum"
	// Custom requirements can input any text as the interrupt value.
	Custom RequirementTypes = "custom"
	// CustomWithSuggestions requirements can input any with a given suggestions as hints.
	CustomWithSuggestions RequirementTypes = "custom_with_suggestions"
)

// Requirement defines the constaints for the interrupt requirements.
type Requirement struct {
	// Defines the type of requirmenet.
	Type RequirementTypes `json:"type"`
	// Suggestions for the values of the requirement.
	Suggestions []string `json:"suggestions"`
}

// Requirements is hash map of all the requirements which needs input from the user.
type Requirements map[string]Requirement

// Validate method does a basic validation on the provided values against the constrainsts
// defined in the Requirements.
func (r Requirements) Validate(values map[string]string) error {
	for k := range r {
		data, ok := values[k]
		if !ok {
			return RequirementKeyNotFound(k)
		}

		if data == "" {
			return RequirementKeyNotFound(k)
		}
	}

	return nil
}

// HITLInterrupt is used to return a invoke a human in the loop
// routine as a part of the flow.
type HITLInterrupt struct {
	Reason          string       `json:"reason"`
	Message         string       `json:"message"`
	ValidationError error        `json:"validation_error"`
	Requirements    Requirements `json:"requirements"`
	InterruptID     InterruptID  `json:"interrupt_id"`
}

// Error implements the error interface for the task interrupt.
func (it HITLInterrupt) Error() string {
	return fmt.Sprintf("flow interrupted: %s", it.Reason)
}

// ConditionalInterrupt is used to direct the execution of a flow
// using a alias value. This value will then be used to choose the
// next edge of the graph.
type ConitionalInterrupt struct {
	Value string
}

// Error implements the error interface for the conditional interrupt.
func (ci ConitionalInterrupt) Error() string {
	return fmt.Sprintf("conditional interrupt: directing to %s", ci.Value)
}
