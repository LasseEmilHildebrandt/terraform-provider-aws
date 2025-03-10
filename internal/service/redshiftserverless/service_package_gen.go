// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceCredentials,
			TypeName: "aws_redshiftserverless_credentials",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceEndpointAccess,
			TypeName: "aws_redshiftserverless_endpoint_access",
		},
		{
			Factory:  ResourceNamespace,
			TypeName: "aws_redshiftserverless_namespace",
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_redshiftserverless_resource_policy",
		},
		{
			Factory:  ResourceSnapshot,
			TypeName: "aws_redshiftserverless_snapshot",
		},
		{
			Factory:  ResourceUsageLimit,
			TypeName: "aws_redshiftserverless_usage_limit",
		},
		{
			Factory:  ResourceWorkgroup,
			TypeName: "aws_redshiftserverless_workgroup",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RedshiftServerless
}

var ServicePackage = &servicePackage{}
