package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"strings"

	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmCosBucketGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmCosBucketGenerator{}

func (x *TableIbmCosBucketGenerator) GetTableName() string {
	return "ibm_cos_bucket"
}

func (x *TableIbmCosBucketGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmCosBucketGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmCosBucketGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmCosBucketGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			region := ibm_client.GetDefaultIBMRegion(taskClient.(*ibm_client.Client).Config)

			serviceType := taskClient.(*ibm_client.Client).ServiceType

			if serviceType != "cloud-object-storage" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			conn, err := ibm_client.CosService(ctx, taskClient.(*ibm_client.Client).Config, region)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opt := &s3.ListBucketsExtendedInput{}

			data, err := conn.ListBucketsExtended(opt)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range data.Buckets {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func getBucketLifecycle(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	serviceType := taskClient.(*ibm_client.Client).ServiceType

	if serviceType != "cloud-object-storage" {
		return nil, nil
	}
	bucket := result.(*s3.BucketExtended)

	location := strings.TrimSuffix(*bucket.LocationConstraint, "-smart")

	conn, err := ibm_client.CosService(ctx, taskClient.(*ibm_client.Client).Config, location)
	if err != nil {
		return nil, err
	}

	params := &s3.GetBucketLifecycleConfigurationInput{
		Bucket: bucket.Name,
	}

	lifecycle, err := conn.GetBucketLifecycleConfiguration(params)
	if err != nil {
		if strings.Contains(err.Error(), "lifecycle configuration does not exist") {
			return nil, nil
		}
		return nil, err
	}
	return lifecycle, nil
}

func getBucketVersioning(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	serviceType := taskClient.(*ibm_client.Client).ServiceType

	if serviceType != "cloud-object-storage" {
		return nil, nil
	}
	bucket := result.(*s3.BucketExtended)

	location := strings.TrimSuffix(*bucket.LocationConstraint, "-smart")

	conn, err := ibm_client.CosService(ctx, taskClient.(*ibm_client.Client).Config, location)
	if err != nil {
		return nil, err
	}

	params := &s3.GetBucketVersioningInput{
		Bucket: bucket.Name,
	}

	versioning, err := conn.GetBucketVersioning(params)
	if err != nil {

		if strings.Contains(err.Error(), "bucket does not exist") {
			return nil, nil
		}
		return nil, err
	}

	return versioning, nil
}

func headBucket(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	serviceType := taskClient.(*ibm_client.Client).ServiceType

	if serviceType != "cloud-object-storage" {
		return nil, nil
	}
	bucket := result.(*s3.BucketExtended)

	location := strings.TrimSuffix(*bucket.LocationConstraint, "-smart")

	conn, err := ibm_client.CosService(ctx, taskClient.(*ibm_client.Client).Config, location)
	if err != nil {
		return nil, err
	}

	params := &s3.HeadBucketInput{
		Bucket: bucket.Name,
	}

	lifecycle, err := conn.HeadBucket(params)
	if err != nil {

		if strings.Contains(err.Error(), "Not Found") {
			return nil, nil
		}
		return nil, err
	}
	return lifecycle, nil
}

func (x *TableIbmCosBucketGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildServiceInstanceList()
}

func (x *TableIbmCosBucketGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Description("The date when the bucket was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acl").ColumnType(schema.ColumnTypeJSON).Description("The access control list (ACL) of a bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_rules").ColumnType(schema.ColumnTypeJSON).Description("The lifecycle configuration information of the bucket.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getBucketLifecycle(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Rules")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sse_kp_enabled").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the Bucket has Key Protect enabled.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := headBucket(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("IBMSSEKPEnabled")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_enabled").ColumnType(schema.ColumnTypeBool).Description("The versioning state of a bucket.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getBucketVersioning(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Status")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_mfa_delete").ColumnType(schema.ColumnTypeBool).Description("The MFA Delete status of the versioning state.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getBucketVersioning(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("MFADelete")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_access_block_configuration").ColumnType(schema.ColumnTypeJSON).Description("The public access block configuration information of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention").ColumnType(schema.ColumnTypeJSON).Description("The retention configuration information of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("website").ColumnType(schema.ColumnTypeJSON).Description("The lifecycle configuration information of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of the bucket.").
			Extractor(column_value_extractor.StructSelector("LocationConstraint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sse_kp_customer_root_key_crn").ColumnType(schema.ColumnTypeString).Description("The root key used by Key Protect to encrypt this bucket. This value must be the full CRN of the root key.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := headBucket(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("IBMSSEKPCrkId")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableIbmCosBucketGenerator) GetSubTables() []*schema.Table {
	return nil
}
