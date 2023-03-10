package provider

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "ibm",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var ibmConfig ibm_client.Configs

				err := config.Unmarshal(&ibmConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(ibmConfig.Providers) == 0 {
					ibmConfig.Providers = append(ibmConfig.Providers, ibm_client.Config{})
				}

				if ibmConfig.Providers[0].APIKey == "" {
					ibmConfig.Providers[0].APIKey = os.Getenv("IBM_API_KEY")
				}

				if ibmConfig.Providers[0].APIKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing api_key in configuration")
				}

				if len(ibmConfig.Providers[0].Regions) == 0 {
					regionData := os.Getenv("IBM_REGIONS")

					var regionList []string

					if regionData != "" {
						regionList = strings.Split(regionData, ",")
					}

					ibmConfig.Providers[0].Regions = regionList
				}

				if len(ibmConfig.Providers[0].Regions) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing regions in configuration")
				}

				clients, err := ibm_client.NewClients(ibmConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# api_key: "<YOUR_API_KEY>"
# regions:
#   - us-south`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var client_config ibm_client.Configs
				err := config.Unmarshal(&client_config.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
