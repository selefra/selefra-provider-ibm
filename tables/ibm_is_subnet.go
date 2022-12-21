package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsSubnetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsSubnetGenerator{}

func (x *TableIbmIsSubnetGenerator) GetTableName() string {
	return "ibm_is_subnet"
}

func (x *TableIbmIsSubnetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsSubnetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsSubnetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsSubnetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			region := taskClient.(*ibm_client.Client).Region

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(100)
			start := ""

			opts := &vpcv1.ListSubnetsOptions{
				Limit: &maxResult,
			}

			for {
				if start != "" {
					opts.Start = &start
				}

				result, _, err := conn.ListSubnetsWithContext(ctx, opts)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range result.Subnets {
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

func (x *TableIbmIsSubnetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsSubnetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("available_ipv4_address_count").ColumnType(schema.ColumnTypeInt).Description("The number of IPv4 addresses in this subnet that are not in-use, and have not been reserved by the user or the provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_acl").ColumnType(schema.ColumnTypeJSON).Description("The network ACL for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("routing_table").ColumnType(schema.ColumnTypeJSON).Description("The routing table for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeJSON).Description("The zone this subnet resides in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this subnet.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The CRN for this subnet.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_version").ColumnType(schema.ColumnTypeString).Description("The IP version(s) supported by this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc").ColumnType(schema.ColumnTypeJSON).Description("The VPC this subnet is a part of.").
			Extractor(column_value_extractor.StructSelector("VPC")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the subnet was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv4_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The IPv4 range of the subnet, expressed in CIDR format.").
			Extractor(column_value_extractor.StructSelector("Ipv4CIDRBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group").ColumnType(schema.ColumnTypeJSON).Description("The resource group for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The unique user-defined name for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_gateway").ColumnType(schema.ColumnTypeJSON).Description("The public gateway to handle internet bound traffic for this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_ipv4_address_count").ColumnType(schema.ColumnTypeInt).Description("The total number of IPv4 addresses in this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this subnet.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
	}
}

func (x *TableIbmIsSubnetGenerator) GetSubTables() []*schema.Table {
	return nil
}
