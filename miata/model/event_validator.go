package model

type EventReq struct {
	Name string `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}

type TFReq struct {
	A bool `json:"a" binding:"required"`
	B *bool `json:"b" binding:"required"`
	C *bool `json:"c" binding:"omitempty"`

}

type TFFormReq struct {
	A bool `form:"a" binding:"required"`
	B *bool `form:"b" binding:"required"`
	C bool `form:"c"`
	D *bool `form:"d"`
}

type Annotation struct {
	Summary string `json:"summary"`
	Owner string `json:"owner" binding:"required"`
}

type Annotations []Annotation

type NestEventReq struct {
	Name string `json:"name" validate:"required"`
	Annotations []Annotation `json:"annotations" binding:"dive,required"`
}
