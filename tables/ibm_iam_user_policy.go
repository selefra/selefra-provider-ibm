package tables

import (
	"context"
	"github.com/selefra/selefra-provider-ibm/ibm_client"

	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableIbmIamUserPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableIbmIamUserPolicyGenerator{}

func (x *TableIbmIamUserPolicyGenerator) GetTableName() string {
	return "ibm_iam_user_policy"
}

func (x *TableIbmIamUserPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableIbmIamUserPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableIbmIamUserPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableIbmIamUserPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := ibm_client.IamPolicyManagementService(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userData := task.ParentRawResult.(usermanagementv2.UserInfo)

			accountID, err := ibm_client.GetAccountId(ctx, taskClient.(*ibm_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &iampolicymanagementv1.ListPoliciesOptions{
				AccountID: core.StringPtr(accountID.(string)),
				Type:      core.StringPtr("access"),
				IamID:     core.StringPtr(userData.IamID),
			}

			result, _, err := conn.ListPoliciesWithContext(ctx, opts)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range result.Policies {
				resultChannel <- userAccessPolicy{i, userData.IamID}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type userAccessPolicy struct {
	iampolicymanagementv1.Policy
	IamID string
}

func (x *TableIbmIamUserPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableIbmIamUserPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the IAM access group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by_id").ColumnType(schema.ColumnTypeString).Description("The iam ID of the entity that created the policy.").
			Extractor(column_value_extractor.StructSelector("CreatedByID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("href").ColumnType(schema.ColumnTypeString).Description("The href link back to the policy.").
			Extractor(column_value_extractor.StructSelector("Href")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_at").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the policy was last modified.").
			Extractor(ibm_client.ExtractorTimestamp("LastModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_by_id").ColumnType(schema.ColumnTypeString).Description("The iam ID of the entity that last modified the policy.").
			Extractor(column_value_extractor.StructSelector("LastModifiedByID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("roles").ColumnType(schema.ColumnTypeJSON).Description("A set of role cloud resource names (CRNs) granted by the policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("ID of the account that this policy belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_id").ColumnType(schema.ColumnTypeString).Description("An alphanumeric value identifying the user's IAM ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The policy type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time when the policy was created.").
			Extractor(ibm_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).Description("The resources associated with a policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subjects").ColumnType(schema.ColumnTypeJSON).Description("The subjects associated with a policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the IAM user policy.").Build(),
	}
}

func (x *TableIbmIamUserPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
