package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/IBM-Cloud/bluemix-go/api/certificatemanager"
)

type TableIbmCertificateManagerCertificateGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmCertificateManagerCertificateGenerator{}

func (x *TableIbmCertificateManagerCertificateGenerator) GetTableName() string {
	return "ibm_certificate_manager_certificate"
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			instanceCRN := taskClient.(*ibm_client.Client).InstanceCrn
			serviceType := taskClient.(*ibm_client.Client).ServiceType

			if serviceType != "cloudcerts" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			conn, err := ibm_client.Connect(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			svc, err := certificatemanager.New(conn)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := svc.Certificate()

			certificates, err := client.ListCertificates(instanceCRN)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range certificates {
				resultChannel <- i

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

// Transform used to get the instance_crn column
func getServiceInstanceCRN(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	InstanceCrn := taskClient.(*ibm_client.Client).InstanceCrn
	return InstanceCrn, nil
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return ibm_client.BuildServiceInstanceList()
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("begins_on").ColumnType(schema.ColumnTypeTimestamp).Description("The creation date of the certificate.").
			Extractor(ibm_client.ExtractorTimestamp("BeginsOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the certificate that is managed in certificate manager.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_manager_instance_id").ColumnType(schema.ColumnTypeString).Description("The CRN of the certificate manager service instance.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getServiceInstanceCRN(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("algorithm").ColumnType(schema.ColumnTypeString).Description("The Algorithm of a certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_renew_enabled").ColumnType(schema.ColumnTypeBool).Description("The automatic renewal status of the certificate.").
			Extractor(column_value_extractor.StructSelector("OrderPolicy.AutoRenewEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("order_policy_name").ColumnType(schema.ColumnTypeString).Description("The order policy name of the certificate.").
			Extractor(column_value_extractor.StructSelector("OrderPolicy.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("akas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires_on").ColumnType(schema.ColumnTypeTimestamp).Description("The expiration date of the certificate.").
			Extractor(ibm_client.ExtractorTimestamp("ExpiresOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_previous").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a certificate has a previous version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("imported").ColumnType(schema.ColumnTypeBool).Description("Indicates whether a certificate has imported or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The display name of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer").ColumnType(schema.ColumnTypeString).Description("The issuer of the certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuance_info").ColumnType(schema.ColumnTypeJSON).Description("The issuance information of a certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the account ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("The account ID of this certificate.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The region of this certificate.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := getRegion(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of a certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial_number").ColumnType(schema.ColumnTypeString).Description("The serial number of a certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domains").ColumnType(schema.ColumnTypeJSON).Description("An array of valid domains for the issued certificate. The first domain is the primary domain, extra domains are secondary domains.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotate_keys").ColumnType(schema.ColumnTypeBool).Description("Rotate keys.").Build(),
	}
}

func (x *TableIbmCertificateManagerCertificateGenerator) GetSubTables() []*schema.Table {
	return nil
}
