package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"net/url"
	"reflect"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIsFlowLogGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIsFlowLogGenerator{}

func (x *TableIbmIsFlowLogGenerator) GetTableName() string {
	return "ibm_is_flow_log"
}

func (x *TableIbmIsFlowLogGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIsFlowLogGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIsFlowLogGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIsFlowLogGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*ibm_client.Client).Region

			conn, err := ibm_client.VpcService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			maxResult := int64(100)
			start := ""

			opts := &vpcv1.ListFlowLogCollectorsOptions{
				Limit: &maxResult,
			}

			for {
				if start != "" {
					opts.Start = &start
				}

				result, _, err := conn.ListFlowLogCollectorsWithContext(ctx, opts)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, i := range result.FlowLogCollectors {
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

func GetNext(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}
	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return ""
	}
	q := u.Query()
	return q.Get("start")
}

// Transform used to get the region column
func getRegion(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	region := taskClient.(*ibm_client.Client).Region
	return region, nil
}

func (x *TableIbmIsFlowLogGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildRegionList()
}

func (x *TableIbmIsFlowLogGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The unique user-defined name for this flow log collector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).Description("Indicates whether this collector is active.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group").ColumnType(schema.ColumnTypeJSON).Description("The resource group for this flow log collector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc").ColumnType(schema.ColumnTypeJSON).Description("The VPC this flow log collector is associated with.").
			Extractor(column_value_extractor.StructSelector("VPC")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this flow log collector.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for this flow log collector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The URL for this flow log collector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The CRN for this flow log collector.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_delete").ColumnType(schema.ColumnTypeBool).Description("If set to `true`, this flow log collector will be automatically deleted when the target is deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The lifecycle state of the flow log collector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time that the flow log collector was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_bucket").ColumnType(schema.ColumnTypeJSON).Description("The Cloud Object Storage bucket where the collected flows are logged.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeJSON).Description("The target this collector is collecting flow logs for.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this flow log collector.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
	}
}

func (x *TableIbmIsFlowLogGenerator) GetSubTables() []*schema.Table {
	return nil
}
