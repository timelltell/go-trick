package PopeFS

import (
	"fmt"
	"git.xiaojukeji.com/falcon/pope-fs/global"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"git.xiaojukeji.com/gobiz/logger"
	"github.com/BurntSushi/toml"
	"golang.org/x/net/context"
)

func main() {
	var err error
	//init main config
	_, err = toml.DecodeFile(global.GetMainFile(), &global.Config)
	if err != nil {
		fmt.Printf("failed to parse config file||errno:%d||errmsg:%s\n", errutil.ErrConfigParseFailed, err.Error())
		return
	}

	//setup logger
	if err := initLogger(global.GetModeFile(global.LoggerConfigFile)); err != nil {
		fmt.Println("logger init failed ", err)
		return
	}

	// init and start http server
	if err := initHTTPServer(global.Config.Nereus.Port); err != nil {
		logger.Errorf(context.TODO(), logger.DLTagUndefined, "init http server fail||errno:%d||errmsg:service_init_fail||err=%s",
			errutil.ErrInitServiceFailed, err.Error())
		return
	}
}
