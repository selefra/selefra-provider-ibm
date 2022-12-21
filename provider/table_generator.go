package provider

import (
	"github.com/selefra/selefra-provider-ibm/table_schema_generator"
	"github.com/selefra/selefra-provider-ibm/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableIbmKmsKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsVpcGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmResourceGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsSecurityGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsFlowLogGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmKmsKeyRingGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsSubnetGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsNetworkAclGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsRegionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmCertificateManagerCertificateGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsVolumeGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmAccountGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmCisDomainGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamRoleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmCosBucketGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamMyApiKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamApiKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIsInstanceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamAccountSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableIbmIamAccessGroupGenerator{}),
	}
}
