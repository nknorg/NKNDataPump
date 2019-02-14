package config

import (
	. "NKNDataPump/common"
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	CFG_SERVICE_TYPE_FULL   = "full"
	CFG_SERVICE_TYPE_PUMP   = "pump"
	CFG_SERVICE_TYPE_SERVER = "server"

	NODE_RPC_SERVER_V1 = "1.0"
	NODE_RPC_SERVER_V2 = "2.0"
)

type DataPumpConfig struct {
	HttpScheme  	string `json:"HttpScheme"`
	NodeServer  	string `json:"NodeServer"`
	NodePort    	string `json:"NodePort"`
	NodeRootURI 	string `json:"NodeRootURI"`
	NodeRPCVersion  string `json:"NodeRPCVersion"`

	APIServerPort uint16 `json:"APIServerPort"`
	WebServerPort uint16 `json:"WebServerPort"`
	WebDir        string `json:"WebDir"`

	ServiceType    string `json:"ServiceType"`
	ServiceBaseURI string `json:"ServiceBaseURI"`

	ServiceDBUser  string `json:"ServiceDBUser"`
	ServiceDBPwd   string `json:"ServiceDBPwd"`
	ServiceDBName  string `json:"ServiceDBName"`

	Logfile  string `json:"Logfile"`
	LogLevel uint32 `json:"LogLevel"`
}

var PumpConfig = DataPumpConfig{
	HttpScheme:  "http",
	NodeServer:  "127.0.0.1",
	NodePort:    "",
	NodeRootURI: "",
	NodeRPCVersion: NODE_RPC_SERVER_V2,

	APIServerPort: 7890,

	WebServerPort: 7891,
	WebDir:        "web",

	ServiceType:    "",
	ServiceBaseURI: "/",

	ServiceDBUser: "root",
	ServiceDBPwd : "password",
	ServiceDBName: "chain_data",

	Logfile:  "gw_log.log",
	LogLevel: uint32(logrus.DebugLevel),
}

func (c *DataPumpConfig) updateBaseUrl() {
	if "" == c.NodePort {
		c.NodeRootURI = ""
		return
	}
	c.NodeRootURI = c.HttpScheme + "://" + c.NodeServer + ":" + c.NodePort
}

func (c *DataPumpConfig) SetNodePort(port string) {
	PumpConfig.NodePort = port
	c.updateBaseUrl()
}

func NewConfig(configFile string) (gwErr *GatewayError) {

	cfg := DataPumpConfig{}

	if !FileExist(configFile) {
		return &GatewayError{Code: GW_ERR_FILE_NOT_FOUND, UserData: configFile}
	}

	cfgStr, err := ioutil.ReadFile(configFile)

	if nil != err {
		return &GatewayError{Code: GW_ERR_CAN_NOT_READ_FILE, UserData: err}
	}

	cfgStr = bytes.TrimPrefix(cfgStr, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(cfgStr, &cfg)

	if nil != err {
		return &GatewayError{Code: GW_ERR_JSON_UNMARSHAL,
			UserData: "unmarshal config file" + configFile + " failed" + err.Error()}
	}

	if "" == cfg.NodePort &&
		"" == PumpConfig.NodePort {
		return &GatewayError{Code: GW_ERR_CONFIG,
			UserData: "none config file nor cli parameter has node port value setting"}
	}

	cfg.updateBaseUrl()

	StructDataMerge(&PumpConfig, &cfg, &DataPumpConfig{})

	return
}
