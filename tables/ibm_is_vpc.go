package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsVpcGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsVpcGenerator{}

func (x *TableIbmIsVpcGenerator) GetTableName() string {
	return "ibm_is_vpc"
}

func (x *TableIbmIsVpcGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsVpcGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsVpcGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsVpcGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			region := taskClient.(*ibm_client.Client).Region

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(100)
			start := ""

			opts := &vpcv1.ListVpcsOptions{
				Limit: &maxResult,
			}

			for {
				if start != "" {
					opts.Start = &start
				}

				result, _, err := conn.ListVpcsWithContext(ctx, opts)

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range result.Vpcs {
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

func (x *TableIbmIsVpcGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsVpcGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the VPC was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classic_access").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this VPC is connected to Classic Infrastructure.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The CRN for this VPC.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_network_acl").ColumnType(schema.ColumnTypeJSON).Description("The default network ACL to use for subnets created in this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_routing_table").ColumnType(schema.ColumnTypeJSON).Description("The default routing table to use for subnets created in this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_security_group").ColumnType(schema.ColumnTypeJSON).Description("The default security group to use for network interfaces created in this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group").ColumnType(schema.ColumnTypeJSON).Description("The resource group for this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cse_source_ips").ColumnType(schema.ColumnTypeJSON).Description("Array of CSE source IP addresses for the VPC. The VPC will have one CSE source IP address per zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this VPC.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this VPC.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The unique user-defined name for this VPC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_prefixes").ColumnType(schema.ColumnTypeJSON).Description("Array of all address pool prefixes for this VPC.").Build(),
	}
}

func (x *TableIbmIsVpcGenerator) GetSubTables() []*schema.Table {
	return nil
}
