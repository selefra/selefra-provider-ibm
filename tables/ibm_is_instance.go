package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsInstanceGenerator{}

func (x *TableIbmIsInstanceGenerator) GetTableName() string {
	return "ibm_is_instance"
}

func (x *TableIbmIsInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*ibm_client.Client).Region

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(100)
			start := ""

			opts := &vpcv1.ListInstancesOptions{
				Limit: &maxResult,
			}

			for {
				if start != "" {
					opts.Start = &start
				}

				result, _, err := conn.ListInstancesWithContext(ctx, opts)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, i := range result.Instances {
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

func (x *TableIbmIsInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("floating_ips").ColumnType(schema.ColumnTypeJSON).Description("Floating IPs allow inbound and outbound traffic from the Internet to an instance").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile").ColumnType(schema.ColumnTypeJSON).Description("The profile for this virtual server instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc").ColumnType(schema.ColumnTypeJSON).Description("The VPC this virtual server instance resides in.").
			Extractor(column_value_extractor.StructSelector("VPC")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeJSON).Description("The zone this virtual server instance resides in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The CRN for this virtual server instance.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the virtual server instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the virtual server instance was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user-defined name for this virtual server instance (and default system hostname).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gpu").ColumnType(schema.ColumnTypeJSON).Description("The virtual server instance GPU configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcpu").ColumnType(schema.ColumnTypeJSON).Description("The virtual server instance VCPU configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interfaces").ColumnType(schema.ColumnTypeJSON).Description("A collection of the virtual server instance's network interfaces, including the primary network interface.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary_network_interface").ColumnType(schema.ColumnTypeJSON).Description("Specifies the primary network interface.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group").ColumnType(schema.ColumnTypeJSON).Description("The resource group for this instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_attachments").ColumnType(schema.ColumnTypeJSON).Description("A collection of the virtual server instance's volume attachments, including the boot volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bandwidth").ColumnType(schema.ColumnTypeInt).Description("The total bandwidth (in megabits per second) shared across the virtual server instance's network interfaces.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disks").ColumnType(schema.ColumnTypeJSON).Description("A collection of the instance's disks.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeJSON).Description("The image the virtual server instance was provisioned from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_volume_attachment").ColumnType(schema.ColumnTypeJSON).Description("Specifies the boot volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this instance.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for this virtual server instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this virtual server instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeInt).Description("The amount of memory, truncated to whole gibibytes.").Build(),
	}
}

func (x *TableIbmIsInstanceGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableIbmIsInstanceDiskGenerator{}),
	}
}
