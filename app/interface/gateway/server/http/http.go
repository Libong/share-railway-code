package http

import (
	"libong/common/server/http"
	loginConf "libong/login/app/interface/login/conf"
	loginHttp "libong/login/app/interface/login/server/http"
	loginService "libong/login/app/interface/login/service"
	"share-railway-code/app/interface/gateway/conf"
	//recordInterfaceConf "for-share/app/interface/record/conf"
	//recordInterfaceHttp "for-share/app/interface/record/server/http"
	//recordInterfaceService "for-share/app/interface/record/service"
)

func New(conf *conf.Config) *http.Server {
	server := http.New(conf.Server.HTTP)
	//
	//recordInterfaceHttp.ConfigHttp(recordInterfaceService.New(&recordInterfaceConf.Service{
	//	Dao: &recordInterfaceConf.Dao{
	//		Redis: conf.Service.Dao.Redis,
	//		Mysql: conf.Service.Dao.Mysql,
	//	},
	//}), server)
	loginHttp.ConfigHttp(loginService.New(&loginConf.Service{
		//LoginRSAKey: &loginConf.RSAKey{
		//	PublicKeyPath:  conf.Service.LoginRSAKey.LoginPublic,
		//	PrivateKeyPath: conf.Service.LoginRSAKey.LoginPrivate,
		//},
		Dao: &loginConf.Dao{
			Redis: conf.Service.Dao.Redis,
			Mysql: conf.Service.Dao.Mysql,
		},
		RBACService: conf.Service.RBACService,
	}), server)
	return server
}
