package req

import (
	"http_5/pkg/res"
	"net/http"
)

// error handlin is done in this function, so when we call it after
// we just need to check err and return from function
func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}

	return &body, nil
}
