package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// awsTagsResourceTypeFilterFile returns the file path that contains a list of
// AWS resource types. Terraform would inject tags specified by apply command's
// `-aws-tag` option only to these resource types.
func awsTagsResourceTypeFilter() map[string]bool {
	dir, err := ConfigDir()
	if err != nil {
		log.Printf("[ERROR] Error finding global config directory: %s", err)
	}

	os.MkdirAll(dir, 0755)
	file := filepath.Join(dir, "aws_tags_resource_type_filter")
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		content, _ := json.MarshalIndent(defaultAwsResourceTypeFilter(), "", "  ")
		ioutil.WriteFile(file, content, 0644)
	}

	resourceTypes := make(map[string]bool)
	if raw, err := ioutil.ReadFile(file); err != nil {
		log.Printf("[ERROR] Error reading AWS tags resource type filter file %s. Error: %s", file, err)
	} else {
		var content []string
		json.Unmarshal(raw, &content)
		for _, resourceType := range content {
			resourceTypes[resourceType] = true
		}
	}
	return resourceTypes
}

func defaultAwsResourceTypeFilter() []string {
	return []string{
		"aws_ami",
		"aws_autoscaling_group",
		"aws_batch_compute_environment",
		"aws_cloudformation_stack",
		"aws_cloudfront_distribution",
		"aws_cloudtrail",
		"aws_cloudwatch_log_group",
		"aws_codebuild_project",
		"aws_cognito_user_pool",
		"aws_customer_gateway",
		"aws_db_event_subscription",
		"aws_db_instance",
		"aws_db_option_group",
		"aws_db_parameter_group",
		"aws_db_security_group",
		"aws_db_subnet_group",
		"aws_default_network_acl",
		"aws_default_route_table",
		"aws_directory_service_directory",
		"aws_dms_endpoint",
		"aws_dms_replication_instance",
		"aws_dms_replication_subnet_group",
		"aws_dms_replication_task",
		"aws_dynamodb_table",
		"aws_ebs_snapshot",
		"aws_ebs_volume",
		"aws_efs_file_system",
		"aws_eip",
		"aws_elastic_beanstalk_environment",
		"aws_elasticache_cluster",
		"aws_elasticache_replication_group",
		"aws_elasticsearch_domain",
		"aws_elb",
		"aws_emr_cluster",
		"aws_glacier_vault",
		"aws_inspector_resource_group",
		"aws_instance",
		"aws_internet_gateway",
		"aws_kinesis_stream",
		"aws_kms_key",
		"aws_lambda_function",
		"aws_lb",
		"aws_lb_target_group",
		"aws_nat_gateway",
		"aws_network_acl",
		"aws_network_interface",
		"aws_opsworks_stack",
		"aws_rds_cluster",
		"aws_rds_cluster_instance",
		"aws_rds_cluster_parameter_group",
		"aws_redshift_cluster",
		"aws_redshift_subnet_group",
		"aws_route53_health_check",
		"aws_route53_zone",
		"aws_route_table",
		"aws_s3_bucket",
		"aws_s3_bucket_object",
		"aws_security_group",
		"aws_servicecatalog_portfolio",
		"aws_spot_fleet_request",
		"aws_spot_instance_request",
		"aws_sqs_queue",
		"aws_subnet",
		"aws_vpc",
		"aws_vpc_dhcp_options",
		"aws_vpc_peering_connection",
		"aws_vpc_peering_connection_accepter",
		"aws_vpn_connection",
		"aws_vpn_gateway",
	}
}
