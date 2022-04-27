// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pseudo resource representing the default VPC and associated subnets for an account and region. This does not create any resources. This will be replaced with `getDefaultVpc` in the future.
type DefaultVpc struct {
	pulumi.ResourceState

	PrivateSubnetIds pulumi.StringArrayOutput `pulumi:"privateSubnetIds"`
	PublicSubnetIds  pulumi.StringArrayOutput `pulumi:"publicSubnetIds"`
	// The VPC ID for the default VPC
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultVpc registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpc(ctx *pulumi.Context,
	name string, args *DefaultVpcArgs, opts ...pulumi.ResourceOption) (*DefaultVpc, error) {
	if args == nil {
		args = &DefaultVpcArgs{}
	}

	var resource DefaultVpc
	err := ctx.RegisterRemoteComponentResource("awsx:vpc:DefaultVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type defaultVpcArgs struct {
}

// The set of arguments for constructing a DefaultVpc resource.
type DefaultVpcArgs struct {
}

func (DefaultVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcArgs)(nil)).Elem()
}

type DefaultVpcInput interface {
	pulumi.Input

	ToDefaultVpcOutput() DefaultVpcOutput
	ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput
}

func (*DefaultVpc) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpc)(nil)).Elem()
}

func (i *DefaultVpc) ToDefaultVpcOutput() DefaultVpcOutput {
	return i.ToDefaultVpcOutputWithContext(context.Background())
}

func (i *DefaultVpc) ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcOutput)
}

// DefaultVpcArrayInput is an input type that accepts DefaultVpcArray and DefaultVpcArrayOutput values.
// You can construct a concrete instance of `DefaultVpcArrayInput` via:
//
//          DefaultVpcArray{ DefaultVpcArgs{...} }
type DefaultVpcArrayInput interface {
	pulumi.Input

	ToDefaultVpcArrayOutput() DefaultVpcArrayOutput
	ToDefaultVpcArrayOutputWithContext(context.Context) DefaultVpcArrayOutput
}

type DefaultVpcArray []DefaultVpcInput

func (DefaultVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultVpc)(nil)).Elem()
}

func (i DefaultVpcArray) ToDefaultVpcArrayOutput() DefaultVpcArrayOutput {
	return i.ToDefaultVpcArrayOutputWithContext(context.Background())
}

func (i DefaultVpcArray) ToDefaultVpcArrayOutputWithContext(ctx context.Context) DefaultVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcArrayOutput)
}

// DefaultVpcMapInput is an input type that accepts DefaultVpcMap and DefaultVpcMapOutput values.
// You can construct a concrete instance of `DefaultVpcMapInput` via:
//
//          DefaultVpcMap{ "key": DefaultVpcArgs{...} }
type DefaultVpcMapInput interface {
	pulumi.Input

	ToDefaultVpcMapOutput() DefaultVpcMapOutput
	ToDefaultVpcMapOutputWithContext(context.Context) DefaultVpcMapOutput
}

type DefaultVpcMap map[string]DefaultVpcInput

func (DefaultVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultVpc)(nil)).Elem()
}

func (i DefaultVpcMap) ToDefaultVpcMapOutput() DefaultVpcMapOutput {
	return i.ToDefaultVpcMapOutputWithContext(context.Background())
}

func (i DefaultVpcMap) ToDefaultVpcMapOutputWithContext(ctx context.Context) DefaultVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcMapOutput)
}

type DefaultVpcOutput struct{ *pulumi.OutputState }

func (DefaultVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpc)(nil)).Elem()
}

func (o DefaultVpcOutput) ToDefaultVpcOutput() DefaultVpcOutput {
	return o
}

func (o DefaultVpcOutput) ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput {
	return o
}

type DefaultVpcArrayOutput struct{ *pulumi.OutputState }

func (DefaultVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultVpc)(nil)).Elem()
}

func (o DefaultVpcArrayOutput) ToDefaultVpcArrayOutput() DefaultVpcArrayOutput {
	return o
}

func (o DefaultVpcArrayOutput) ToDefaultVpcArrayOutputWithContext(ctx context.Context) DefaultVpcArrayOutput {
	return o
}

func (o DefaultVpcArrayOutput) Index(i pulumi.IntInput) DefaultVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultVpc {
		return vs[0].([]*DefaultVpc)[vs[1].(int)]
	}).(DefaultVpcOutput)
}

type DefaultVpcMapOutput struct{ *pulumi.OutputState }

func (DefaultVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultVpc)(nil)).Elem()
}

func (o DefaultVpcMapOutput) ToDefaultVpcMapOutput() DefaultVpcMapOutput {
	return o
}

func (o DefaultVpcMapOutput) ToDefaultVpcMapOutputWithContext(ctx context.Context) DefaultVpcMapOutput {
	return o
}

func (o DefaultVpcMapOutput) MapIndex(k pulumi.StringInput) DefaultVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultVpc {
		return vs[0].(map[string]*DefaultVpc)[vs[1].(string)]
	}).(DefaultVpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcInput)(nil)).Elem(), &DefaultVpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcArrayInput)(nil)).Elem(), DefaultVpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcMapInput)(nil)).Elem(), DefaultVpcMap{})
	pulumi.RegisterOutputType(DefaultVpcOutput{})
	pulumi.RegisterOutputType(DefaultVpcArrayOutput{})
	pulumi.RegisterOutputType(DefaultVpcMapOutput{})
}
