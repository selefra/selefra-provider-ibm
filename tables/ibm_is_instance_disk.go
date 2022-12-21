package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"strings"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsInstanceDiskGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsInstanceDiskGenerator{}

func (x *TableIbmIsInstanceDiskGenerator) GetTableName() string {
	return "ibm_is_instance_disk"
}

func (x *TableIbmIsInstanceDiskGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsInstanceDiskGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsInstanceDiskGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsInstanceDiskGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*ibm_client.Client).Region

			instanceData := task.ParentRawResult.(vpcv1.Instance)
			var instanceRegion string
			splitCRN := strings.Split(*instanceData.CRN, ":")
			if len(splitCRN) > 5 {
				instanceRegion = strings.Split(*instanceData.CRN, ":")[5]
			}

			if !strings.Contains(instanceRegion, region) {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &vpcv1.ListInstanceDisksOptions{
				InstanceID: instanceData.ID,
			}

			result, _, err := conn.ListInstanceDisksWithContext(ctx, opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Disks {
				resultChannel <- instanceDiskInfo{i, *instanceData.ID}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type instanceDiskInfo = struct {
	vpcv1.InstanceDisk
	InstanceId string
}

func (x *TableIbmIsInstanceDiskGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsInstanceDiskGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).Description("The resource type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The size of the disk in GB (gigabytes).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this instance disk.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the instance disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the disk was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this instance disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interface_type").ColumnType(schema.ColumnTypeString).Description("The disk interface used for attaching the disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user defined name for this disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The instance identifier.").
			Extractor(column_value_extractor.StructSelector("InstanceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this instance disk.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableIbmIsInstanceDiskGenerator) GetSubTables() []*schema.Table {
	return nil
}
