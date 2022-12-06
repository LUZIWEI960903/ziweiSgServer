package middleware

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
)

func CheckRole() net.MiddlewareFunc {
	return func(next net.HandlerFunc) net.HandlerFunc {
		return func(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
			log.Println("进入到CheckRole()中间件")
			_, err := req.Conn.GetProperty("role")
			if err != nil {
				rsp.Body.Code = constant.SessionInvalid
				return
			}
			next(req, rsp)
		}
	}
}
