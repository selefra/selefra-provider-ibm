package ibm_client

type Client struct {
	Region      string
	InstanceId  string
	InstanceCrn string
	ServiceType string
	Config      *Config
}

func NewClients(configs Configs) ([]*Client, error) {
	var clients []*Client
	for i := range configs.Providers {
		clients = append(clients, &Client{Config: &configs.Providers[i]})
	}
	return clients, nil
}

func (x *Client) CopyWithRegion(region string) *Client {
	return &Client{
		Region: region,
		Config: x.Config,
	}
}

func (x *Client) CopyWithService(instanceId string, instanceCrn string, region string, serviceType string) *Client {
	return &Client{
		Region:      region,
		InstanceId:  instanceId,
		InstanceCrn: instanceCrn,
		ServiceType: serviceType,
		Config:      x.Config,
	}
}
