package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIamAccessGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamAccessGroupGenerator{}

func (x *TableIbmIamAccessGroupGenerator) GetTableName() string {
	return "ibm_iam_access_group"
}

func (x *TableIbmIamAccessGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamAccessGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamAccessGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamAccessGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.IamAccessGroupService(ctx, taskClient.(*ibm_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &iamaccessgroupsv2.ListAccessGroupsOptions{
				AccountID: core.StringPtr(accountID.(string)),
			}

			maxResult := int64(100)

			opts.SetLimit(maxResult)

			result, _, err := conn.ListAccessGroups(opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Groups {
				resultChannel <- i

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableIbmIamAccessGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamAccessGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the IAM access group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("Specifies the date and time, the group las modified.").
			Extractor(ibm_client.ExtractorTimestamp("LastModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("An url to the given group resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the access group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_federated").ColumnType(schema.ColumnTypeBool).Description("This is set to true if rules exist for the group.").
			Extractor(column_value_extractor.StructSelector("IsFederated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp the group was created at.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the IAM access group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by_id").ColumnType(schema.ColumnTypeString).Description("The iam_id of the entity that created the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("ID of the account that this group belongs to.").Build(),
	}
}

func (x *TableIbmIamAccessGroupGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableIbmIamAccessGroupPolicyGenerator{}),
	}
}
