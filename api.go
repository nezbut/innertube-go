package innertubego

import "strings"

func GetContext(clientName string) ClientContext {
	for _, clientContext := range config.Clients {
		if strings.EqualFold(strings.ToUpper(clientContext.ClientName), strings.ToUpper(clientName)) {
			return clientContext
		}
	}
	return config.Clients[0]
}
