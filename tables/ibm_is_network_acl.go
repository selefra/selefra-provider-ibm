package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsNetworkAclGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsNetworkAclGenerator{}

func (x *TableIbmIsNetworkAclGenerator) GetTableName() string {
	return "ibm_is_network_acl"
}

func (x *TableIbmIsNetworkAclGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsNetworkAclGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsNetworkAclGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsNetworkAclGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			region := taskClient.(*ibm_client.Client).Region

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(100)
			start := ""

			opts := &vpcv1.ListNetworkAclsOptions{
				Limit: &maxResult,
			}

			for {
				if start != "" {
					opts.Start = &start
				}

				result, _, err := conn.ListNetworkAclsWithContext(ctx, opts)

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range result.NetworkAcls {
					resultChannel <- i

				}
				start = GetNext(result.Next)
				if start == "" {
					break
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmIsNetworkAclGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsNetworkAclGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for this network ACL").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc").ColumnType(schema.ColumnTypeJSON).Description("he VPC this network ACL is a part of.").
			Extractor(column_value_extractor.StructSelector("VPC")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this subnet.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group").ColumnType(schema.ColumnTypeJSON).Description("The resource group for this network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeJSON).Description("The subnets to which this network ACL is attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Description("The ordered rules for this network ACL. If no rules exist, all traffic will be denied.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this subnet.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user-defined name for this network ACL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The CRN for this network ACL.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the network ACL was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
	}
}

func (x *TableIbmIsNetworkAclGenerator) GetSubTables() []*schema.Table {
	return nil
}
