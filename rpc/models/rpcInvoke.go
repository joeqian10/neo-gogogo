package models

type InvokeResult struct {
	Script      string        `json:"script"`
	State       string        `json:"state"`
	GasConsumed string        `json:"gas_consumed"`
	Stack       []InvokeStack `json:"stack"`
	Tx          string        `json:"tx"`
}

type InvokeStack struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Convert converts interface{} "Value" to string or []InvokeStack depending on the "Type"
func (s *InvokeStack) Convert()  {
	if s.Type != "Array" {
		s.Value = s.Value.(string)
	} else {
		vs := s.Value.([]interface{})
		result := make([]InvokeStack, len(vs))
		for i, v := range vs {
			m := v.(map[string]interface{})
			s2 := InvokeStack{
				Type:  m["type"].(string),
				Value: m["value"],
			}
			s2.Convert()
			result[i] = s2
		}
		s.Value = result
	}
}
