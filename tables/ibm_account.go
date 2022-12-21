package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableIbmAccountGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmAccountGenerator{}

func (x *TableIbmAccountGenerator) GetTableName() string {
	return "ibm_account"
}

func (x *TableIbmAccountGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmAccountGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmAccountGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmAccountGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.Connect(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			svc, err := accountv2.New(conn)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := svc.Accounts()

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			data, err := client.Get(accountID.(string))
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- *data

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmAccountGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmAccountGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The current state of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("currency_code").ColumnType(schema.ColumnTypeString).Description("Specifies the currency type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_guid").ColumnType(schema.ColumnTypeString).Description("An unique Id of the account owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_unique_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the account owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizations").ColumnType(schema.ColumnTypeJSON).Description("A list of organizations the account is associated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("members").ColumnType(schema.ColumnTypeJSON).Description("A list of members associated with this account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Specifies the name of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guid").ColumnType(schema.ColumnTypeString).Description("An unique ID of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country_code").ColumnType(schema.ColumnTypeString).Description("Specifies the country code.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_id").ColumnType(schema.ColumnTypeString).Description("The customer ID of the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_user_id").ColumnType(schema.ColumnTypeString).Description("The owner user ID used for login.").Build(),
	}
}

func (x *TableIbmAccountGenerator) GetSubTables() []*schema.Table {
	return nil
}
