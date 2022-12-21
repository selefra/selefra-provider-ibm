package ibm_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"path"
	"sort"
	"strings"
)

func ConfigApiKey(_ context.Context, config *Config) (string, error) {
	if config.APIKey != "" {
		return config.APIKey, nil
	}

	// No key, cannot proceed
	return "", errors.New("api_key must be configured")
}

// UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	userAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
	generation  int    `default:"2"`
}

func fetchUserDetails(sess *session.Session, generation int) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}
	var bluemixToken string
	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}
	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.userEmail = email.(string)
	}
	user.userID = claims["id"].(string)
	user.userAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"
	user.generation = generation
	return &user, nil
}

// GetAccountId Get current user account
func GetAccountId(ctx context.Context, config *Config) (interface{}, error) {
	conn, err := Connect(ctx, config)
	if err != nil {
		return "", err
	}

	userInfo, err := fetchUserDetails(conn, 2)
	if err != nil {
		return nil, err
	}

	return userInfo.userAccount, nil
}

func BuildServiceInstanceList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		// get all the service instances in the account
		serviceInstances, err := listAllServiceInstances(ctx, taskClient.(*Client).Config)
		if err != nil {
			panic(err)
		}

		slice := make([]*schema.ClientTaskContext, 0)
		for _, instance := range serviceInstances {
			splitID := strings.Split(*instance, ":")
			slice = append(slice, &schema.ClientTaskContext{
				Task:   task.Clone(),
				Client: taskClient.(*Client).CopyWithService(splitID[7], *instance, splitID[5], splitID[4]),
			})
		}
		return slice
	}
}

func listAllServiceInstances(ctx context.Context, config *Config) ([]*string, error) {
	// Create Session
	session, err := ResourceControllerService(ctx, config)
	if err != nil {
		return nil, err
	}

	var serviceInstanceCRNs []*string

	opts := &resourcecontrollerv2.ListResourceInstancesOptions{
		Type: core.StringPtr("service_instance"),
	}

	response, _, err := session.ListResourceInstances(opts)
	if err != nil {
		return nil, err
	}

	for _, i := range response.Resources {
		serviceInstanceCRNs = append(serviceInstanceCRNs, i.CRN)
	}

	return serviceInstanceCRNs, err
}

// Regions is the current known list of valid regions
func Regions() []string {
	return []string{
		"au-syd",
		"br-sao",
		"ca-tor",
		"eu-de",
		"eu-gb",
		"jp-osa",
		"jp-tok",
		"us-east",
		"us-south",
	}
}

// var pluginQueryData *plugin.QueryData

// func init() {
// 	pluginQueryData = &plugin.QueryData{
// 		ConnectionManager: connection.NewManager(),
// 	}
// }

func BuildRegionList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		var allRegions []string

		if taskClient.(*Client).Config.Regions != nil {
			regions := Regions()
			for _, pattern := range taskClient.(*Client).Config.Regions {
				for _, validRegion := range regions {
					if ok, _ := path.Match(pattern, validRegion); ok {
						allRegions = append(allRegions, validRegion)
					}
				}
			}
		}

		// Build regions matrix using config regions
		if len(allRegions) > 0 {
			uniqueRegions := unique(allRegions)

			if len(getInvalidRegions(uniqueRegions)) > 0 {
				panic("\n\nConnection config have invalid regions: " + strings.Join(getInvalidRegions(uniqueRegions), ","))
			}

			// validate regions list
			matrix := make([]map[string]interface{}, len(uniqueRegions))
			for i, region := range uniqueRegions {
				matrix[i] = map[string]interface{}{"region": region}
			}

			// validate regions list
			slice := make([]*schema.ClientTaskContext, 0)
			for _, region := range uniqueRegions {
				slice = append(slice, &schema.ClientTaskContext{
					Task:   task.Clone(),
					Client: taskClient.(*Client).CopyWithRegion(region),
				})
			}
			return slice
		}

		// Search for region configured using env, or use default region (i.e. us-south)
		defaultIBMRegion := GetDefaultIBMRegion(taskClient.(*Client).Config)

		// validate regions list
		slice := make([]*schema.ClientTaskContext, 0)
		slice = append(slice, &schema.ClientTaskContext{
			Task:   task.Clone(),
			Client: taskClient.(*Client).CopyWithRegion(defaultIBMRegion),
		})
		return slice
	}
}

// Return invalid regions from a region list
func getInvalidRegions(regions []string) []string {
	invalidRegions := []string{}
	for _, region := range regions {
		if !in(region, Regions()) {
			invalidRegions = append(invalidRegions, region)
		}
	}
	return invalidRegions
}

// GetDefaultIBMRegion returns the default region for IBM account
// if not set by Env variable
func GetDefaultIBMRegion(config *Config) string {
	// have we already created and cached the service?
	allIBMRegions := Regions()

	var regions []string
	var region string

	if config.Regions != nil {
		regions = config.Regions
		region = regions[0]
	}

	validPatterns := []string{}
	invalidPatterns := []string{}
	for _, namePattern := range regions {
		validRegions := []string{}
		for _, validRegion := range allIBMRegions {
			if ok, _ := path.Match(namePattern, validRegion); ok {
				validRegions = append(validRegions, validRegion)
			}
		}
		if len(validRegions) == 0 {
			invalidPatterns = append(invalidPatterns, namePattern)
		} else {
			validPatterns = append(validPatterns, namePattern)
		}
	}

	if len(validPatterns) == 0 {
		panic("\nconnection config have invalid \"regions\": " + strings.Join(invalidPatterns, ", ") + ". Edit your connection configuration file and then restart Steampipe")
	}

	if !in(region, allIBMRegions) {
		region = "us-south"
	}

	return region
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// Returns a list of unique items
func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ExtractorTimestamp(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector(path).Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		ts, ok := v.(strfmt.DateTime)
		if !ok {
			return nil, schema.NewDiagnosticsAddErrorMsg(fmt.Sprintf("unextected type, wanted \"strfmt.DateTime\", have \"%T\"", v))
		}
		return ts.String(), nil
	})
}
