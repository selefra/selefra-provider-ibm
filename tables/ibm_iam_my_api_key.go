package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIamMyApiKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamMyApiKeyGenerator{}

func (x *TableIbmIamMyApiKeyGenerator) GetTableName() string {
	return "ibm_iam_my_api_key"
}

func (x *TableIbmIamMyApiKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamMyApiKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamMyApiKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamMyApiKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			conn, err := ibm_client.IamService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &iamidentityv1.ListAPIKeysOptions{}

			result, _, err := conn.ListAPIKeys(opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Apikeys {
				resultChannel <- i

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableIbmIamMyApiKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamMyApiKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the API key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entity_tag").ColumnType(schema.ColumnTypeString).Description("Version of the API Key details object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("ID of the account that this API key authenticates for.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the date and time, the API key las modified.").
			Extractor(ibm_client.ExtractorTimestamp("ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Specifies the name of the API key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of this API Key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_id").ColumnType(schema.ColumnTypeString).Description("The iam_id that this API key authenticates.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the date and time, the API key is created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("Cloud Resource Name of the API key.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_key").ColumnType(schema.ColumnTypeString).Description("The API key value. This property only contains the API key value for the following cases: create an API key, update a service ID API key that stores the API key value as retrievable, or get a service ID API key that stores the API key value as retrievable.").
			Extractor(column_value_extractor.StructSelector("Apikey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("history").ColumnType(schema.ColumnTypeJSON).Description("History of the API key.").Build(),
	}
}

func (x *TableIbmIamMyApiKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
