package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"strings"

	"github.com/IBM/networking-go-sdk/zonessettingsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmCisDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmCisDomainGenerator{}

func (x *TableIbmCisDomainGenerator) GetTableName() string {
	return "ibm_cis_domain"
}

func (x *TableIbmCisDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmCisDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmCisDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmCisDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.CisZoneService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			result, _, err := conn.ListZones(&zonesv1.ListZonesOptions{})
			if err != nil {
				if strings.Contains(err.Error(), "Not Found") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Result {
				resultChannel <- i
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func getTlsMinimumVersion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	id := result.(zonesv1.ZoneDetails).ID

	conn, err := ibm_client.CisZoneSettingService(ctx, taskClient.(*ibm_client.Client).Config, *id)
	if err != nil {
		return nil, err
	}

	tls, _, err := conn.GetMinTlsVersion(&zonessettingsv1.GetMinTlsVersionOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		}
		return nil, err
	}

	return tls.Result, nil
}

func getWebApplicationFirewall(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	id := result.(zonesv1.ZoneDetails).ID

	conn, err := ibm_client.CisZoneSettingService(ctx, taskClient.(*ibm_client.Client).Config, *id)
	if err != nil {

		return nil, err
	}

	firewall, _, err := conn.GetWebApplicationFirewall(&zonessettingsv1.GetWebApplicationFirewallOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		}
		return nil, err
	}
	return firewall.Result, nil
}

func (x *TableIbmCisDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmCisDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The zone name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The zone status.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_load_balancer").ColumnType(schema.ColumnTypeJSON).Description("The global load balancer of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The zone id.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("An unique ID of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_name_servers").ColumnType(schema.ColumnTypeJSON).Description("The original name servers of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_servers").ColumnType(schema.ColumnTypeJSON).Description("The name servers of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_dnshost").ColumnType(schema.ColumnTypeString).Description("The original DNS host of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paused").ColumnType(schema.ColumnTypeBool).Description("Whether the zone is in paused state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the zone was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_on").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the zone was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_tls_version").ColumnType(schema.ColumnTypeString).Description("The tls version of the zone.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getTlsMinimumVersion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Value")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_registrar").ColumnType(schema.ColumnTypeString).Description("The original registrar of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_application_firewall").ColumnType(schema.ColumnTypeString).Description("The web application firewall state.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getWebApplicationFirewall(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Value")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_records").ColumnType(schema.ColumnTypeJSON).Description("DNS records for the domain.").Build(),
	}
}

func (x *TableIbmCisDomainGenerator) GetSubTables() []*schema.Table {
	return nil
}
