package server

import (
	"sync"
)

var (
	dataFile = saveDir + "/data.json"
	rcd      *record
)

type parameters struct {
	ServerName    string `json:"serverName"`
	ProjectName   string `json:"projectName"`
	ModuleName    string `json:"moduleName"`
	RepoAddr      string `json:"repoAddr"`
	ProtobufFile  string `json:"-"`
	YamlFile      string `json:"-"`
	DbDriver      string `json:"dbDriver"`
	Dsn           string `json:"dsn"`
	TableName     string `json:"tableName"`
	Embed         bool   `json:"embed"`
	IncludeInitDB bool   `json:"includeInitDB"`
	UpdateAt      string `json:"updateAt"`

	TemplateDir string `json:"templateDir"`
	Fields      string `json:"fields"`
	DepProtoDir string `json:"depProtoDir"`
	OnlyPrint   bool   `json:"onlyPrint"`

	LLMType       string `json:"llmType"`
	LLMModel      string `json:"llmModel"`
	APIKey        string `json:"apiKey"`
	GoDir         string `json:"goDir"`
	GoFile        string `json:"goFile"`
	TargetType    string `json:"targetType"`
	IsCleanAICode bool   `json:"isCleanAICode"`

	Protocol      string   `json:"protocol"`
	URL           string   `json:"url"`
	Method        string   `json:"method"`
	Body          string   `json:"body"`
	Headers       []string `json:"headers"`
	Worker        int      `json:"worker"`
	TestType      string   `json:"testType"`
	TotalRequests uint64   `json:"totalRequests"`
	Duration      string   `json:"duration"`
	PushType      string   `json:"pushType"`
	PushURL       string   `json:"pushUrl"`
	PrometheusURL string   `json:"prometheusUrl"`
	JobName       string   `json:"jobName"`

	SuitedMonoRepo bool `json:"suitedMonoRepo"`
}

type record struct {
	mux        *sync.Mutex
	HostRecord map[string]*parameters // [ip + "-" + commandType]:parameters
}

func initRecord() { _ = "STUB: not implemented"; return }

func recordObj() *record { _ = "STUB: not implemented"; return nil }

func (r *record) set(ip string, commandType string, params *parameters) {
	_ = "STUB: not implemented"
	return
}

func getKey(ip string, commandType string) string { _ = "STUB: not implemented"; return "" }

func (r *record) get(ip string, commandType string) *parameters {
	_ = "STUB: not implemented"
	return nil
}

// nolint
func parseCommandArgs(args []string) *parameters { _ = "STUB: not implemented"; return nil }

//nolint

//nolint
