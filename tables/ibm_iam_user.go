package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIamUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamUserGenerator{}

func (x *TableIbmIamUserGenerator) GetTableName() string {
	return "ibm_iam_user"
}

func (x *TableIbmIamUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.Connect(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			svc, err := usermanagementv2.New(conn)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := svc.UserInvite()

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			users, err := client.ListUsers(accountID.(string))
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range users {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmIamUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the user profile.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("realm").ColumnType(schema.ColumnTypeString).Description("The realm of the user. The value is either IBMid or SL.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("The first name of the user.").
			Extractor(column_value_extractor.StructSelector("Firstname")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("The last name of the user.").
			Extractor(column_value_extractor.StructSelector("Lastname")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The state of the user. Possible values are PROCESSING, PENDING, ACTIVE, DISABLED_CLASSIC_INFRASTRUCTURE, and VPN_ONLY.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phonenumber").ColumnType(schema.ColumnTypeString).Description("The phone number of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("User settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_id").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the user's IAM ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The user ID used for login.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The email of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alt_phonenumber").ColumnType(schema.ColumnTypeString).Description("The alternative phone number of the user.").
			Extractor(column_value_extractor.StructSelector("Altphonenumber")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo").ColumnType(schema.ColumnTypeString).Description("A link to a photo of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the account ID.").Build(),
	}
}

func (x *TableIbmIamUserGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableIbmIamUserPolicyGenerator{}),
	}
}
