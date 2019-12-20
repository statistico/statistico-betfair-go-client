package betfair

type ErrorResponse struct {
	FaultCode   string `json:"faultcode"`
	FaultString string `json:"faultstring"`
	Detail      struct {
		ExceptionName         string `json:"exceptionname"`
		AccountAPINGException struct {
			RequestUUID  string `json:"requestUUID"`
			ErrorCode    string `json:"errorCode"`
			ErrorDetails string `json:"errorDetails"`
		} `json:"AccountAPINGException"`
	} `json:"detail,omitempty"`
}
