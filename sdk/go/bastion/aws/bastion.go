// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-tailscale-bastion/sdk/go/bastion/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Bastion struct {
	pulumi.ResourceState

	// The name of the ASG that managed the bastion instances
	AsgName pulumi.StringOutput `pulumi:"asgName"`
	// The SSH private key to access your bastion
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
}

// NewBastion registers a new resource with the given unique name, arguments, and options.
func NewBastion(ctx *pulumi.Context,
	name string, args *BastionArgs, opts ...pulumi.ResourceOption) (*Bastion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Route == nil {
		return nil, errors.New("invalid value for required argument 'Route'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bastion
	err := ctx.RegisterRemoteComponentResource("tailscale-bastion:aws:Bastion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type bastionArgs struct {
	// The EC2 instance type to use for the bastion.
	InstanceType *string `pulumi:"instanceType"`
	// The AWS region you're using.
	Region string `pulumi:"region"`
	// The route you'd like to advertise via tailscale.
	Route string `pulumi:"route"`
	// The subnet Ids to launch instances in.
	SubnetIds []string `pulumi:"subnetIds"`
	// The VPC the Bastion should be created in.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Bastion resource.
type BastionArgs struct {
	// The EC2 instance type to use for the bastion.
	InstanceType pulumi.StringPtrInput
	// The AWS region you're using.
	Region pulumi.StringInput
	// The route you'd like to advertise via tailscale.
	Route pulumi.StringInput
	// The subnet Ids to launch instances in.
	SubnetIds pulumi.StringArrayInput
	// The VPC the Bastion should be created in.
	VpcId pulumi.StringInput
}

func (BastionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionArgs)(nil)).Elem()
}

type BastionInput interface {
	pulumi.Input

	ToBastionOutput() BastionOutput
	ToBastionOutputWithContext(ctx context.Context) BastionOutput
}

func (*Bastion) ElementType() reflect.Type {
	return reflect.TypeOf((**Bastion)(nil)).Elem()
}

func (i *Bastion) ToBastionOutput() BastionOutput {
	return i.ToBastionOutputWithContext(context.Background())
}

func (i *Bastion) ToBastionOutputWithContext(ctx context.Context) BastionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionOutput)
}

func (i *Bastion) ToOutput(ctx context.Context) pulumix.Output[*Bastion] {
	return pulumix.Output[*Bastion]{
		OutputState: i.ToBastionOutputWithContext(ctx).OutputState,
	}
}

// BastionArrayInput is an input type that accepts BastionArray and BastionArrayOutput values.
// You can construct a concrete instance of `BastionArrayInput` via:
//
//	BastionArray{ BastionArgs{...} }
type BastionArrayInput interface {
	pulumi.Input

	ToBastionArrayOutput() BastionArrayOutput
	ToBastionArrayOutputWithContext(context.Context) BastionArrayOutput
}

type BastionArray []BastionInput

func (BastionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bastion)(nil)).Elem()
}

func (i BastionArray) ToBastionArrayOutput() BastionArrayOutput {
	return i.ToBastionArrayOutputWithContext(context.Background())
}

func (i BastionArray) ToBastionArrayOutputWithContext(ctx context.Context) BastionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionArrayOutput)
}

func (i BastionArray) ToOutput(ctx context.Context) pulumix.Output[[]*Bastion] {
	return pulumix.Output[[]*Bastion]{
		OutputState: i.ToBastionArrayOutputWithContext(ctx).OutputState,
	}
}

// BastionMapInput is an input type that accepts BastionMap and BastionMapOutput values.
// You can construct a concrete instance of `BastionMapInput` via:
//
//	BastionMap{ "key": BastionArgs{...} }
type BastionMapInput interface {
	pulumi.Input

	ToBastionMapOutput() BastionMapOutput
	ToBastionMapOutputWithContext(context.Context) BastionMapOutput
}

type BastionMap map[string]BastionInput

func (BastionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bastion)(nil)).Elem()
}

func (i BastionMap) ToBastionMapOutput() BastionMapOutput {
	return i.ToBastionMapOutputWithContext(context.Background())
}

func (i BastionMap) ToBastionMapOutputWithContext(ctx context.Context) BastionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionMapOutput)
}

func (i BastionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bastion] {
	return pulumix.Output[map[string]*Bastion]{
		OutputState: i.ToBastionMapOutputWithContext(ctx).OutputState,
	}
}

type BastionOutput struct{ *pulumi.OutputState }

func (BastionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bastion)(nil)).Elem()
}

func (o BastionOutput) ToBastionOutput() BastionOutput {
	return o
}

func (o BastionOutput) ToBastionOutputWithContext(ctx context.Context) BastionOutput {
	return o
}

func (o BastionOutput) ToOutput(ctx context.Context) pulumix.Output[*Bastion] {
	return pulumix.Output[*Bastion]{
		OutputState: o.OutputState,
	}
}

// The name of the ASG that managed the bastion instances
func (o BastionOutput) AsgName() pulumi.StringOutput {
	return o.ApplyT(func(v *Bastion) pulumi.StringOutput { return v.AsgName }).(pulumi.StringOutput)
}

// The SSH private key to access your bastion
func (o BastionOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Bastion) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

type BastionArrayOutput struct{ *pulumi.OutputState }

func (BastionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bastion)(nil)).Elem()
}

func (o BastionArrayOutput) ToBastionArrayOutput() BastionArrayOutput {
	return o
}

func (o BastionArrayOutput) ToBastionArrayOutputWithContext(ctx context.Context) BastionArrayOutput {
	return o
}

func (o BastionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Bastion] {
	return pulumix.Output[[]*Bastion]{
		OutputState: o.OutputState,
	}
}

func (o BastionArrayOutput) Index(i pulumi.IntInput) BastionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bastion {
		return vs[0].([]*Bastion)[vs[1].(int)]
	}).(BastionOutput)
}

type BastionMapOutput struct{ *pulumi.OutputState }

func (BastionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bastion)(nil)).Elem()
}

func (o BastionMapOutput) ToBastionMapOutput() BastionMapOutput {
	return o
}

func (o BastionMapOutput) ToBastionMapOutputWithContext(ctx context.Context) BastionMapOutput {
	return o
}

func (o BastionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bastion] {
	return pulumix.Output[map[string]*Bastion]{
		OutputState: o.OutputState,
	}
}

func (o BastionMapOutput) MapIndex(k pulumi.StringInput) BastionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bastion {
		return vs[0].(map[string]*Bastion)[vs[1].(string)]
	}).(BastionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BastionInput)(nil)).Elem(), &Bastion{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionArrayInput)(nil)).Elem(), BastionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionMapInput)(nil)).Elem(), BastionMap{})
	pulumi.RegisterOutputType(BastionOutput{})
	pulumi.RegisterOutputType(BastionArrayOutput{})
	pulumi.RegisterOutputType(BastionMapOutput{})
}
