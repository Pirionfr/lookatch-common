package events

type (
	Offset struct {
		Database string `json:"database"`
		Agent    string `json:"agent"`
	}
	/**
	 * SQL events
	 */
	//type stream
	SqlEvent struct {
		Tenant       string  `json:"Tenant,omitempty"`
		Environment  string  `json:"Environment"`
		Timestamp    string  `json:"Timestamp"`
		Database     string  `json:"Database"`
		Schema       string  `json:"schema"`
		Table        string  `json:"Table"`
		Method       string  `json:"Method"`
		PrimaryKey   string  `json:"PrimaryKey,omitempty"`
		Offset       *Offset `json:"Offset,omitempty"`
		Statement    string  `json:"Statement"`
		StatementOld string  `json:"StatementOld,omitempty"`
	}

	//type Generic
	GenericEvent struct {
		Tenant      string `json:"Tenant"`
		Environment string `json:"Environment"`
		Timestamp   string `json:"Timestamp"`
		AgentId     string `json:"AgentId"`
		Value       string `json:"value"`
	}

	/**
	 * Global events
	 */
	LookatchTenantInfo struct {
		Id  string `json:"id"`
		Env string `json:"env"`
	}

	LookatchHeader struct {
		EventType string
		Tenant    LookatchTenantInfo
	}

	LookatchEvent struct {
		Header  *LookatchHeader
		Payload interface{}
	}

	ErrorEvent struct {
		Timestamp int64
		Level     int
		Host      string
		Short_msg string
		Full_msg  string
	}
)
