package request

type CreateHouseRequest struct {
	ID   int32  `json:"id" validate:"nonnil"`
	Code string `json:"code"`
	Name string `json:"name"`
}
