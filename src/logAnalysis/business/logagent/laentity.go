package logagent

type NodeCondition struct {
	Nodename string
}

type NodeCollection struct {
	Nodename        string
	NLlog           string
	NLErrlog        string
	Atslog          string
	AtsErrlog       string
	HttpsNLlog      string
	HttpsErrorNllog string
}
