package responses

type GetJSConfigResult struct {
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
	NonceStr  string `json:"nonce_str"`
	AppID     string `json:"app_id"`
	URL       string `json:"url"`
}

type GetJSAgentConfigResult struct {
	CorpID    string `json:"corp_id"`
	AgentID   int64  `json:"agent_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
	URL       string `json:"url"`
}
