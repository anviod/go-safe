package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

// 构建a记录查询的dns服务器
// dig @localhost baidu.com 测试  systemctl stop systemd-resolved.service 停止原有的dns服务，解决53端口占用问题
func main() {

	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) { // req请求本身
		var resp dns.Msg
		resp.SetReply(req)

		for _, q := range req.Question {
			a := dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    0,
				},
				A: net.ParseIP("127.0.0.1").To4(),
			}
			resp.Answer = append(resp.Answer, &a)
		}
		w.WriteMsg(&resp)
	})
	log.Fatal(dns.ListenAndServe(":53", "udp", nil))
}
