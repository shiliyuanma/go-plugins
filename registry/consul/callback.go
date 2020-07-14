package consul

import consul "github.com/hashicorp/consul/api"

var (
	listeners []func(*consul.AgentServiceRegistration)
)

func AddRegistryListener(l func(*consul.AgentServiceRegistration)) {
	listeners = append(listeners, l)
}

func onRegistry(asr *consul.AgentServiceRegistration) {
	for _, l := range listeners {
		l(asr)
	}
}
