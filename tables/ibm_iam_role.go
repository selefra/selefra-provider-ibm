package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableIbmIamRoleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamRoleGenerator{}

func (x *TableIbmIamRoleGenerator) GetTableName() string {
	return "ibm_iam_role"
}

func (x *TableIbmIamRoleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamRoleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamRoleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamRoleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.Connect(ctx, taskClient.(*ibm_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			svc, err := iampapv2.New(conn)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := svc.IAMRoles()

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := iampapv2.RoleQuery{AccountID: accountID.(string)}
			roles, err := client.ListAll(opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range roles {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableIbmIamRoleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamRoleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("actions").ColumnType(schema.ColumnTypeJSON).Description("The actions of the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The role ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crn").ColumnType(schema.ColumnTypeString).Description("The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by_id").ColumnType(schema.ColumnTypeString).Description("The IAM ID of the entity that created the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the role was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the account ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name of the role that is shown in the console.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the role was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_by_id").ColumnType(schema.ColumnTypeString).Description("The IAM ID of the entity that last modified the policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the role that is used in the CRN.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).Description("The service name.").Build(),
	}
}

func (x *TableIbmIamRoleGenerator) GetSubTables() []*schema.Table {
	return nil
}
