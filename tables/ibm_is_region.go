package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsRegionGenerator{}

func (x *TableIbmIsRegionGenerator) GetTableName() string {
	return "ibm_is_region"
}

func (x *TableIbmIsRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			region := ibm_client.GetDefaultIBMRegion(taskClient.(*ibm_client.Client).Config)

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &vpcv1.ListRegionsOptions{}

			result, _, err := conn.ListRegionsWithContext(ctx, opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Regions {
				resultChannel <- i
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableIbmIsRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIsRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The unique user-defined name for this region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this region.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).Description("The API endpoint for this region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of this region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this region.").Build(),
	}
}

func (x *TableIbmIsRegionGenerator) GetSubTables() []*schema.Table {
	return nil
}
