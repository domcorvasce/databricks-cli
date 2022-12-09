// Generated from Databricks Terraform provider schema. DO NOT EDIT.

package schema

type ResourcePermissionsAccessControl struct {
	GroupName            string `json:"group_name,omitempty"`
	PermissionLevel      string `json:"permission_level"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`
}

type ResourcePermissions struct {
	Authorization     string                             `json:"authorization,omitempty"`
	ClusterId         string                             `json:"cluster_id,omitempty"`
	ClusterPolicyId   string                             `json:"cluster_policy_id,omitempty"`
	DirectoryId       string                             `json:"directory_id,omitempty"`
	DirectoryPath     string                             `json:"directory_path,omitempty"`
	ExperimentId      string                             `json:"experiment_id,omitempty"`
	Id                string                             `json:"id,omitempty"`
	InstancePoolId    string                             `json:"instance_pool_id,omitempty"`
	JobId             string                             `json:"job_id,omitempty"`
	NotebookId        string                             `json:"notebook_id,omitempty"`
	NotebookPath      string                             `json:"notebook_path,omitempty"`
	ObjectType        string                             `json:"object_type,omitempty"`
	PipelineId        string                             `json:"pipeline_id,omitempty"`
	RegisteredModelId string                             `json:"registered_model_id,omitempty"`
	RepoId            string                             `json:"repo_id,omitempty"`
	RepoPath          string                             `json:"repo_path,omitempty"`
	SqlAlertId        string                             `json:"sql_alert_id,omitempty"`
	SqlDashboardId    string                             `json:"sql_dashboard_id,omitempty"`
	SqlEndpointId     string                             `json:"sql_endpoint_id,omitempty"`
	SqlQueryId        string                             `json:"sql_query_id,omitempty"`
	AccessControl     []ResourcePermissionsAccessControl `json:"access_control,omitempty"`
}