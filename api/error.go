package api

// import "net/http"

// type APIError struct {
// 	Id            string `json:"id"`
// 	StatusCode    int    `json:"status_code,omitempty"`
// 	Message       string `json:"message"`
// 	DetailedError string `json:"detailed_error"`
// 	Where         string `json:"-"`
// 	params        map[string]interface{}
// }

type APIError struct {
	Message string
}

func (c *APIError) Error() string {
	return c.Message
}

// func (c *APIServer) newAPIError(where string, id string, params map[string]interface{}, details string, status int) *APIError {
// 	ap := &APIError{}
// 	ap.Id = id
// 	ap.params = params
// 	ap.Message = id
// 	ap.Where = where
// 	ap.DetailedError = details
// 	ap.StatusCode = status
// 	return ap
// }

// func (c *APIServer) setInvalidParam(w http.ResponseWriter, parameter string) {
// 	err := c.newAPIError("Context", "api.context.invalid_body_param.app_error", map[string]interface{}{"Name": parameter}, "", http.StatusBadRequest)
// 	return err
// }
