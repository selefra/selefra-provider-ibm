package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"strings"

	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmKmsKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmKmsKeyGenerator{}

func (x *TableIbmKmsKeyGenerator) GetTableName() string {
	return "ibm_kms_key"
}

func (x *TableIbmKmsKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmKmsKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmKmsKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmKmsKeyGenerator) GetDataSource() *schema.DataSource {
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

			maxResult := int64(100)

			data, err := conn.GetKeys(ctx, int(maxResult), 0)
			if err != nil {
				if strings.Contains(err.Error(), "key_ring does not exist") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range data.Keys {
				resultChannel <- i
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableIbmKmsKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildServiceInstanceList()
}

func (x *TableIbmKmsKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the key material was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A text field used to provide a more detailed description of the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payload").ColumnType(schema.ColumnTypeString).Description("Specifies the key payload.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The key state based on NIST SP 800-57. States are integers and correspond to the Pre-activation = 0, Active = 1, Suspended = 2, Deactivated = 3, and Destroyed = 5 values.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extractable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the key material can leave the service, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_ring_id").ColumnType(schema.ColumnTypeString).Description("An ID that identifies the key ring.").
			Extractor(column_value_extractor.StructSelector("KeyRingID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aliases").ColumnType(schema.ColumnTypeJSON).Description("A list of key aliases.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imported").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the key was originally imported or generated in Key Protect.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encrypted_nonce").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dual_auth_delete").ColumnType(schema.ColumnTypeJSON).Description("Metadata that indicates the status of a dual authorization policy on the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("algorithm_type").ColumnType(schema.ColumnTypeString).Description("Specifies the key algorithm.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Specifies the MIME type that represents the key resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the key has been deleted, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration").ColumnType(schema.ColumnTypeTimestamp).Description("The date the key material will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date when the key metadata was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A human-readable name assigned to your key for convenience.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date the key material was destroyed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_version").ColumnType(schema.ColumnTypeJSON).Description("Properties associated with a specific key version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_policy").ColumnType(schema.ColumnTypeJSON).Description("Key rotation policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this key.").
			Extractor(column_value_extractor.StructSelector("CRN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this key.").
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
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The key protect instance GUID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getServiceInstanceID(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_by").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the resource that deleted the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_algorithm").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_rotate_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date when the key was last rotated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the resource that created the key.").Build(),
	}
}

func (x *TableIbmKmsKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
