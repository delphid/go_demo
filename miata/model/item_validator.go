package model

type ItemReq struct {
	Name string `json:"name"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}