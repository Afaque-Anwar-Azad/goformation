package networkfirewall

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// FirewallPolicy_StatelessActions AWS CloudFormation Resource (AWS::NetworkFirewall::FirewallPolicy.StatelessActions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessactions.html
type FirewallPolicy_StatelessActions struct {

	// StatelessActions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessactions.html#cfn-networkfirewall-firewallpolicy-statelessactions-statelessactions
	StatelessActions []string `json:"StatelessActions,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *FirewallPolicy_StatelessActions) AWSCloudFormationType() string {
	return "AWS::NetworkFirewall::FirewallPolicy.StatelessActions"
}
