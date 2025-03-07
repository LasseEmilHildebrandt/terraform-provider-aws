// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists apigatewayv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn apigatewayv2iface.ApiGatewayV2API, identifier string) (tftags.KeyValueTags, error) {
	input := &apigatewayv2.GetTagsInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.GetTagsWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) (tftags.KeyValueTags, error) {
	return ListTags(ctx, meta.(*conns.AWSClient).APIGatewayV2Conn(), identifier)
}

// map[string]*string handling

// Tags returns apigatewayv2 service tags.
func Tags(tags tftags.KeyValueTags) map[string]*string {
	return aws.StringMap(tags.Map())
}

// KeyValueTags creates KeyValueTags from apigatewayv2 service tags.
func KeyValueTags(ctx context.Context, tags map[string]*string) tftags.KeyValueTags {
	return tftags.New(ctx, tags)
}

// UpdateTags updates apigatewayv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn apigatewayv2iface.ApiGatewayV2API, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &apigatewayv2.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &apigatewayv2.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).APIGatewayV2Conn(), identifier, oldTags, newTags)
}
