package requester

type Requester struct {
	Url string
	Method string
	Concurrency int
	Requests int
	Body map[string]interface{}
	Headers map[string]interface{}
}

