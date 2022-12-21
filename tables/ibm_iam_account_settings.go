package tables

import (
	"context"
	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/selefra/selefra-provider-ibm/ibm_client"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIamAccountSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamAccountSettingsGenerator{}

func (x *TableIbmIamAccountSettingsGenerator) GetTableName() string {
	return "ibm_iam_account_settings"
}

func (x *TableIbmIamAccountSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamAccountSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamAccountSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamAccountSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.IamService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &iamidentityv1.GetAccountSettingsOptions{
				AccountID: core.StringPtr(accountID.(string)),
			}

			accountSettings, _, err := conn.GetAccountSettings(opts)

			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- accountSettings

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmIamAccountSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamAccountSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("session_expiration_in_seconds").ColumnType(schema.ColumnTypeString).Description("Defines the session expiration in seconds for the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restrict_create_service_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restrict_create_platform_api_key").ColumnType(schema.ColumnTypeString).Description("Indicates whether creating a platform API key is access controlled, or not.").
			Extractor(column_value_extractor.StructSelector("RestrictCreatePlatformApikey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa").ColumnType(schema.ColumnTypeString).Description("Defines the MFA trait for the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("session_invalidation_in_seconds").ColumnType(schema.ColumnTypeString).Description("Defines the period of time in seconds in which a session will be invalidated due  to inactivity.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("history").ColumnType(schema.ColumnTypeJSON).Description("History of the Account Settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("An unique ID of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_ip_addresses").ColumnType(schema.ColumnTypeString).Description("The IP addresses and subnets from which IAM tokens can be created for the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entity_tag").ColumnType(schema.ColumnTypeString).Description("Version of the account settings.").Build(),
	}
}

func (x *TableIbmIamAccountSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
