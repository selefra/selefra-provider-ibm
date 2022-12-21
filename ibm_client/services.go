package ibm_client

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	CloudSession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	kp "github.com/IBM/keyprotect-go-client"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	"github.com/IBM/networking-go-sdk/zonessettingsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	gohttp "net/http"
)

func Connect(ctx context.Context, config *Config) (*CloudSession.Session, error) {
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}

	conf := &bluemix.Config{
		BluemixAPIKey: apiKey,
	}

	conn, err := CloudSession.New(conf)
	if err != nil {
		return nil, err
	}

	err = authenticateAPIKey(conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func authenticateAPIKey(sess *CloudSession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

// KmsService return the service for IBM KMS service
func KmsService(ctx context.Context, config *Config) (*kp.Client, error) {
	// Create region endpoint
	defaultIBMRegion := GetDefaultIBMRegion(config)

	endpoint := fmt.Sprintf("https://%s.kms.cloud.ibm.com", defaultIBMRegion)
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	opts := kp.ClientConfig{
		BaseURL:  endpoint,
		APIKey:   apiKey,
		TokenURL: kp.DefaultTokenURL,
	}
	service, err := kp.New(opts, kp.DefaultTransport())
	if err != nil {
		return nil, err
	}
	return service, nil
}

// VpcService returns the service for IBM VPC Infrastructure service
func VpcService(ctx context.Context, config *Config, region string) (*vpcv1.VpcV1, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed vpcService")
	}

	// Create region endpoint
	endpoint := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region)

	// Fetch API key from config
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}

	// Instantiate the service with an API key based IAM authenticator
	service, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		URL: endpoint,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// CisZoneService returns the service for IBM CIS Zone service
func CisZoneService(ctx context.Context, config *Config) (*zonesv1.ZonesV1, error) {
	serviceInstanceID := "crn:v1:bluemix:public:internet-svcs:global:a/76aa4877fab6436db86f121f62faf221:3e5dc1e0-3aea-4699-986e-e3b8f117c51d::"
	endpoint := "https://api.cis.cloud.ibm.com"

	// Fetch API key from config
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	// Instantiate the service with an API key based IAM authenticator
	service, err := zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
		Crn: &serviceInstanceID,
		URL: endpoint,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// CisZoneSettingService returns the service for IBM CIS Zone Setting service
func CisZoneSettingService(ctx context.Context, config *Config, zoneId string) (*zonessettingsv1.ZonesSettingsV1, error) {
	serviceInstanceID := "crn:v1:bluemix:public:internet-svcs:global:a/76aa4877fab6436db86f121f62faf221:3e5dc1e0-3aea-4699-986e-e3b8f117c51d::"
	endpoint := "https://api.cis.cloud.ibm.com"

	// Fetch API key from config
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	// Instantiate the service with an API key based IAM authenticator
	service, err := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
		Crn:            &serviceInstanceID,
		ZoneIdentifier: &zoneId,
		URL:            endpoint,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// CisGlobalLoadBalancerService returns the service for IBM CIS Global Load Balancer service
func CisGlobalLoadBalancerService(ctx context.Context, config *Config, zoneId string) (*globalloadbalancerv1.GlobalLoadBalancerV1, error) {
	serviceInstanceID := "crn:v1:bluemix:public:internet-svcs:global:a/76aa4877fab6436db86f121f62faf221:3e5dc1e0-3aea-4699-986e-e3b8f117c51d::"
	endpoint := "https://api.cis.cloud.ibm.com"

	// Fetch API key from config
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	// Instantiate the service with an API key based IAM authenticator
	service, err := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
		Crn:            &serviceInstanceID,
		ZoneIdentifier: &zoneId,
		URL:            endpoint,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// CisDnsRecordService returns the service for IBM CIS DNS service
func CisDnsRecordService(ctx context.Context, config *Config, zoneId string) (*dnsrecordsv1.DnsRecordsV1, error) {
	serviceInstanceID := "crn:v1:bluemix:public:internet-svcs:global:a/76aa4877fab6436db86f121f62faf221:3e5dc1e0-3aea-4699-986e-e3b8f117c51d::"
	endpoint := "https://api.cis.cloud.ibm.com"

	// Fetch API key from config
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	// Instantiate the service with an API key based IAM authenticator
	service, err := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
		Crn:            &serviceInstanceID,
		ZoneIdentifier: &zoneId,
		URL:            endpoint,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}

func IamService(ctx context.Context, config *Config) (*iamidentityv1.IamIdentityV1, error) {
	// Load connection from cache, which preserves throttling protection etc
	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	serviceClientOptions := &iamidentityv1.IamIdentityV1Options{Authenticator: &core.IamAuthenticator{
		ApiKey: apiKey,
	}}
	// Instantiate the service with an API key based IAM authenticator
	service, err := iamidentityv1.NewIamIdentityV1UsingExternalConfig(serviceClientOptions)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func IamAccessGroupService(ctx context.Context, config *Config) (*iamaccessgroupsv2.IamAccessGroupsV2, error) {
	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	serviceClientOptions := &iamaccessgroupsv2.IamAccessGroupsV2Options{Authenticator: &core.IamAuthenticator{
		ApiKey: apiKey,
	}}
	service, err := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(serviceClientOptions)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func IamPolicyManagementService(ctx context.Context, config *Config) (*iampolicymanagementv1.IamPolicyManagementV1, error) {
	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	serviceClientOptions := &iampolicymanagementv1.IamPolicyManagementV1Options{Authenticator: &core.IamAuthenticator{
		ApiKey: apiKey,
	}}
	service, err := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(serviceClientOptions)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func TagService(ctx context.Context, config *Config) (*globaltaggingv1.GlobalTaggingV1, error) {

	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	opts := &globaltaggingv1.GlobalTaggingV1Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	service, err := globaltaggingv1.NewGlobalTaggingV1(opts)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func ResourceControllerService(ctx context.Context, config *Config) (*resourcecontrollerv2.ResourceControllerV2, error) {
	apiKey, err := ConfigApiKey(ctx, config)
	if err != nil {
		return nil, err
	}
	opts := &resourcecontrollerv2.ResourceControllerV2Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	service, err := resourcecontrollerv2.NewResourceControllerV2(opts)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func ResourceManagerService(ctx context.Context, config *Config) (*resourcemanagerv2.ResourceManagerV2, error) {
	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	opts := &resourcemanagerv2.ResourceManagerV2Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	service, err := resourcemanagerv2.NewResourceManagerV2(opts)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func CosService(ctx context.Context, config *Config, region string) (*s3.S3, error) {
	serviceInstanceID := `plugin.GetMatrixItem(ctx)["instance_crn"].(string)`

	apiKey, err := ConfigApiKey(ctx, config)

	if err != nil {
		return nil, err
	}
	authEndpoint := "https://iam.cloud.ibm.com/identity/token"
	serviceEndpoint := fmt.Sprintf("s3.%s.cloud-object-storage.appdomain.cloud", region)

	conf := aws.NewConfig().
		WithEndpoint(serviceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			authEndpoint, apiKey, serviceInstanceID)).
		WithS3ForcePathStyle(true)

	sess := session.Must(session.NewSession())

	service := s3.New(sess, conf)

	return service, nil
}
