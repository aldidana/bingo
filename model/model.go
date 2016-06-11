package model

type Cat struct {
  ID string `json:"id" gorethink:"id,omitempty"`
  Name string `json:"name" gorethink:"name"`
  Type string `json:"type" gorethink:"type"`
}
