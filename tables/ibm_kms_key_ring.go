package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmKmsKeyRingGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmKmsKeyRingGenerator{}

func (x *TableIbmKmsKeyRingGenerator) GetTableName() string {
	return "ibm_kms_key_ring"
}

func (x *TableIbmKmsKeyRingGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmKmsKeyRingGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmKmsKeyRingGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmKmsKeyRingGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			instanceID := taskClient.(*ibm_client.Client).InstanceCrn
			serviceType := taskClient.(*ibm_client.Client).ServiceType

			if serviceType != "kms" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			conn, err := ibm_client.KmsService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			conn.Config.InstanceID = instanceID

			data, err := conn.GetKeyRings(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range data.KeyRings {
				resultChannel <- i

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

// Transform used to get the instance_id column
func getServiceInstanceID(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	instanceID := taskClient.(*ibm_client.Client).InstanceCrn
	return instanceID, nil
}

func (x *TableIbmKmsKeyRingGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildServiceInstanceList()
}

func (x *TableIbmKmsKeyRingGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the key ring.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The key protect instance GUID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getServiceInstanceID(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when the key ring was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The creator of the key ring.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this key ring.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this key ring.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
	}
}

func (x *TableIbmKmsKeyRingGenerator) GetSubTables() []*schema.Table {
	return nil
}
