package config

import (
	"fmt"
	"testing"

	"comma/pkg/library/setting"
	config2 "comma/pkg/library/util/config"
)

func Test_LoadConfig(t *testing.T) {
	setting.Server.RootPath = "/home/xxx/dev/xxx/gateway/cmd/gateway"
	config := config2.GetXmlInstance()
	// s, _ := json.MarshalIndent(config, "", "\t")
	// fmt.Print(string(s))
	fmt.Println(config.GetEsRoute("search_all_v2", "all"))
	fmt.Println(config.GetEsFilterType("search_all_v2", "all"))
	fmt.Println(config.GetEsProjectId("search_all_v2", "all"))
}
