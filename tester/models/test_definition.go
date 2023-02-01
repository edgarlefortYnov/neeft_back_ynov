package models

type TestRequestHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TestDefinition struct {
	Name         string              `json:"name"`
	Url          string              `json:"url"`
	Method       string              `json:"method"`
	Headers      []TestRequestHeader `json:"headers"`
	ExpectedCode int                 `json:"expected_code"`
}
