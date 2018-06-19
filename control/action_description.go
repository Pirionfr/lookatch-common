package control

import "reflect"

type ParametersDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
}

type ActionDescription struct {
	Description string                   `json:"description"`
	Parameters  []*ParametersDescription `json:"parameters"`
}

func DeclareNewAction(class interface{}, description string) (action *ActionDescription) {
	action = &ActionDescription{

		Description: description,
	}

	if class == nil {
		return
	}
	t := reflect.TypeOf(class)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		action.Parameters = append(action.Parameters, &ParametersDescription{
			Name:        field.Name,
			Description: field.Tag.Get("description"),
			Type:        field.Type.Name(),
			Required:    field.Tag.Get("required") == "true",
		})
	}

	return
}
