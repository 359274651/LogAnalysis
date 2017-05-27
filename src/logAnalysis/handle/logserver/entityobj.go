package logserver

//存储下拉的列表用的
type MenuData map[DB]Collections

type NodeCollection struct {
	Nodename        string
	NLlog           string
	NLErrlog        string
	Atslog          string
	AtsErrlog       string
	HttpsNLlog      string
	HttpsErrorNllog string
}

type DB string
type Collections []string
