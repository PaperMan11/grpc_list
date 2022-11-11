package discovery

import (
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"
)

const schema = "etcd"

// Server 服务实例
type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`    // 服务地址
	Version string `json:"version"` // 服务版本
	Weight  int64  `json:"weight"`  // 服务权重
}

func BuildPrefix(info Server) string {
	if info.Version == "" {
		return fmt.Sprintf("/%s/", info.Name)
	}
	return fmt.Sprintf("/%s/%s/", info.Name, info.Version)
}

// eg. /srvName/srvVersion/srvAddr
func BuildRegPath(info Server) string {
	return fmt.Sprintf("%s%s", BuildPrefix(info), info.Addr)
}

func ParseValue(value []byte) (Server, error) {
	info := Server{}
	if err := json.Unmarshal(value, &info); err != nil {
		return info, err
	}
	return info, nil
}

func SplitPath(path string) (Server, error) {
	info := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return info, fmt.Errorf("invalid path")
	}
	// `/srvName/srvVersion/srvAddr`
	info.Addr = strs[len(strs)-1]
	return info, nil
}

// Exist 判断服务是否存在
func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}

// Remove 移除服务
func Remove(l []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range l {
		if l[i].Addr == addr.Addr {
			l[i] = l[len(l)-1]
			return l[:len(l)-1], true
		}
	}
	return nil, false
}

func BuildResolver(app string) string {
	return schema + ":///" + app
}
