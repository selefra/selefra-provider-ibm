package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmResourceGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmResourceGroupGenerator{}

func (x *TableIbmResourceGroupGenerator) GetTableName() string {
	return "ibm_resource_group"
}

func (x *TableIbmResourceGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmResourceGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmResourceGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmResourceGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.ResourceManagerService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &resourcemanagerv2.ListResourceGroupsOptions{
				AccountID: core.StringPtr(accountID.(string)),
			}

			result, _, err := conn.ListResourceGroupsWithContext(ctx, opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Resources {
				resultChannel <- i
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmResourceGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmResourceGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this resource group is default of the account or not.").
			Extractor(column_value_extractor.StructSelector("Default")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_id").ColumnType(schema.ColumnTypeString).Description("An alpha-numeric value identifying the quota ID associated with the resource group.").
			Extractor(column_value_extractor.StructSelector("QuotaID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("quota_url").ColumnType(schema.ColumnTypeString).Description("The URL to access the quota details that associated with the resource group.").
			Extractor(column_value_extractor.StructSelector("QuotaURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An alpha-numeric value identifying the resource group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The human-readable name of the resource group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The full CRN (cloud resource name) associated with the resource group.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The state of the resource group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date when the resource group was last updated.").
			Extractor(ibm_client.ExtractorTimestamp("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_linkages").ColumnType(schema.ColumnTypeString).Description("An array of the resources that linked to the resource group.").
			Extractor(column_value_extractor.StructSelector("ResourceLinkages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date when the resource group was initially created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payment_methods_url").ColumnType(schema.ColumnTypeString).Description("The URL to access the payment methods details that associated with the resource group.").
			Extractor(column_value_extractor.StructSelector("PaymentMethodsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams_url").ColumnType(schema.ColumnTypeString).Description("The URL to access the team details that associated with the resource group.").
			Extractor(column_value_extractor.StructSelector("TeamsURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("An alpha-numeric value identifying the account ID.").
			Extractor(column_value_extractor.StructSelector("AccountID")).Build(),
	}
}

func (x *TableIbmResourceGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
