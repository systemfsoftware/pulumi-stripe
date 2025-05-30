// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stripe

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/systemfsoftware/pulumi-stripe/sdk/go/stripe/internal"
)

type Customer struct {
	pulumi.CustomResourceState

	// Address map with fields related to the address: line1, line2, city, state, postalCode and country
	Address pulumi.StringMapOutput `pulumi:"address"`
	// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
	// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
	// the amount due on an invoice.
	Balance pulumi.IntPtrOutput `pulumi:"balance"`
	// The default (auto-generated) prefix for the customer used to generate unique invoice numbers.
	DefaultInvoicePrefix pulumi.StringOutput `pulumi:"defaultInvoicePrefix"`
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
	// and tracking. This may be up to 512 characters.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix pulumi.StringPtrOutput `pulumi:"invoicePrefix"`
	// Default invoice settings for this customer.
	InvoiceSettings pulumi.StringMapOutput `pulumi:"invoiceSettings"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
	// object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The customer’s full name or business name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sequence to be used on the customer’s next invoice. Defaults to 1.
	NextInvoiceSequence pulumi.IntPtrOutput `pulumi:"nextInvoiceSequence"`
	// The customer’s phone number.
	Phone pulumi.StringPtrOutput `pulumi:"phone"`
	// Customer’s preferred languages, ordered by preference.
	PreferredLocales pulumi.StringArrayOutput `pulumi:"preferredLocales"`
	// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
	// country.
	Shipping pulumi.StringMapOutput `pulumi:"shipping"`
}

// NewCustomer registers a new resource with the given unique name, arguments, and options.
func NewCustomer(ctx *pulumi.Context,
	name string, args *CustomerArgs, opts ...pulumi.ResourceOption) (*Customer, error) {
	if args == nil {
		args = &CustomerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Customer
	err := ctx.RegisterResource("stripe:index/customer:Customer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomer gets an existing Customer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerState, opts ...pulumi.ResourceOption) (*Customer, error) {
	var resource Customer
	err := ctx.ReadResource("stripe:index/customer:Customer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Customer resources.
type customerState struct {
	// Address map with fields related to the address: line1, line2, city, state, postalCode and country
	Address map[string]string `pulumi:"address"`
	// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
	// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
	// the amount due on an invoice.
	Balance *int `pulumi:"balance"`
	// The default (auto-generated) prefix for the customer used to generate unique invoice numbers.
	DefaultInvoicePrefix *string `pulumi:"defaultInvoicePrefix"`
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description *string `pulumi:"description"`
	// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
	// and tracking. This may be up to 512 characters.
	Email *string `pulumi:"email"`
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix *string `pulumi:"invoicePrefix"`
	// Default invoice settings for this customer.
	InvoiceSettings map[string]string `pulumi:"invoiceSettings"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
	// object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// The customer’s full name or business name.
	Name *string `pulumi:"name"`
	// The sequence to be used on the customer’s next invoice. Defaults to 1.
	NextInvoiceSequence *int `pulumi:"nextInvoiceSequence"`
	// The customer’s phone number.
	Phone *string `pulumi:"phone"`
	// Customer’s preferred languages, ordered by preference.
	PreferredLocales []string `pulumi:"preferredLocales"`
	// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
	// country.
	Shipping map[string]string `pulumi:"shipping"`
}

type CustomerState struct {
	// Address map with fields related to the address: line1, line2, city, state, postalCode and country
	Address pulumi.StringMapInput
	// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
	// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
	// the amount due on an invoice.
	Balance pulumi.IntPtrInput
	// The default (auto-generated) prefix for the customer used to generate unique invoice numbers.
	DefaultInvoicePrefix pulumi.StringPtrInput
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description pulumi.StringPtrInput
	// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
	// and tracking. This may be up to 512 characters.
	Email pulumi.StringPtrInput
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix pulumi.StringPtrInput
	// Default invoice settings for this customer.
	InvoiceSettings pulumi.StringMapInput
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
	// object in a structured format.
	Metadata pulumi.StringMapInput
	// The customer’s full name or business name.
	Name pulumi.StringPtrInput
	// The sequence to be used on the customer’s next invoice. Defaults to 1.
	NextInvoiceSequence pulumi.IntPtrInput
	// The customer’s phone number.
	Phone pulumi.StringPtrInput
	// Customer’s preferred languages, ordered by preference.
	PreferredLocales pulumi.StringArrayInput
	// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
	// country.
	Shipping pulumi.StringMapInput
}

func (CustomerState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerState)(nil)).Elem()
}

type customerArgs struct {
	// Address map with fields related to the address: line1, line2, city, state, postalCode and country
	Address map[string]string `pulumi:"address"`
	// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
	// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
	// the amount due on an invoice.
	Balance *int `pulumi:"balance"`
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description *string `pulumi:"description"`
	// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
	// and tracking. This may be up to 512 characters.
	Email *string `pulumi:"email"`
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix *string `pulumi:"invoicePrefix"`
	// Default invoice settings for this customer.
	InvoiceSettings map[string]string `pulumi:"invoiceSettings"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
	// object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// The customer’s full name or business name.
	Name *string `pulumi:"name"`
	// The sequence to be used on the customer’s next invoice. Defaults to 1.
	NextInvoiceSequence *int `pulumi:"nextInvoiceSequence"`
	// The customer’s phone number.
	Phone *string `pulumi:"phone"`
	// Customer’s preferred languages, ordered by preference.
	PreferredLocales []string `pulumi:"preferredLocales"`
	// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
	// country.
	Shipping map[string]string `pulumi:"shipping"`
}

// The set of arguments for constructing a Customer resource.
type CustomerArgs struct {
	// Address map with fields related to the address: line1, line2, city, state, postalCode and country
	Address pulumi.StringMapInput
	// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
	// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
	// the amount due on an invoice.
	Balance pulumi.IntPtrInput
	// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
	Description pulumi.StringPtrInput
	// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
	// and tracking. This may be up to 512 characters.
	Email pulumi.StringPtrInput
	// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
	InvoicePrefix pulumi.StringPtrInput
	// Default invoice settings for this customer.
	InvoiceSettings pulumi.StringMapInput
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
	// object in a structured format.
	Metadata pulumi.StringMapInput
	// The customer’s full name or business name.
	Name pulumi.StringPtrInput
	// The sequence to be used on the customer’s next invoice. Defaults to 1.
	NextInvoiceSequence pulumi.IntPtrInput
	// The customer’s phone number.
	Phone pulumi.StringPtrInput
	// Customer’s preferred languages, ordered by preference.
	PreferredLocales pulumi.StringArrayInput
	// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
	// country.
	Shipping pulumi.StringMapInput
}

func (CustomerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerArgs)(nil)).Elem()
}

type CustomerInput interface {
	pulumi.Input

	ToCustomerOutput() CustomerOutput
	ToCustomerOutputWithContext(ctx context.Context) CustomerOutput
}

func (*Customer) ElementType() reflect.Type {
	return reflect.TypeOf((**Customer)(nil)).Elem()
}

func (i *Customer) ToCustomerOutput() CustomerOutput {
	return i.ToCustomerOutputWithContext(context.Background())
}

func (i *Customer) ToCustomerOutputWithContext(ctx context.Context) CustomerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerOutput)
}

// CustomerArrayInput is an input type that accepts CustomerArray and CustomerArrayOutput values.
// You can construct a concrete instance of `CustomerArrayInput` via:
//
//	CustomerArray{ CustomerArgs{...} }
type CustomerArrayInput interface {
	pulumi.Input

	ToCustomerArrayOutput() CustomerArrayOutput
	ToCustomerArrayOutputWithContext(context.Context) CustomerArrayOutput
}

type CustomerArray []CustomerInput

func (CustomerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Customer)(nil)).Elem()
}

func (i CustomerArray) ToCustomerArrayOutput() CustomerArrayOutput {
	return i.ToCustomerArrayOutputWithContext(context.Background())
}

func (i CustomerArray) ToCustomerArrayOutputWithContext(ctx context.Context) CustomerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerArrayOutput)
}

// CustomerMapInput is an input type that accepts CustomerMap and CustomerMapOutput values.
// You can construct a concrete instance of `CustomerMapInput` via:
//
//	CustomerMap{ "key": CustomerArgs{...} }
type CustomerMapInput interface {
	pulumi.Input

	ToCustomerMapOutput() CustomerMapOutput
	ToCustomerMapOutputWithContext(context.Context) CustomerMapOutput
}

type CustomerMap map[string]CustomerInput

func (CustomerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Customer)(nil)).Elem()
}

func (i CustomerMap) ToCustomerMapOutput() CustomerMapOutput {
	return i.ToCustomerMapOutputWithContext(context.Background())
}

func (i CustomerMap) ToCustomerMapOutputWithContext(ctx context.Context) CustomerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerMapOutput)
}

type CustomerOutput struct{ *pulumi.OutputState }

func (CustomerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Customer)(nil)).Elem()
}

func (o CustomerOutput) ToCustomerOutput() CustomerOutput {
	return o
}

func (o CustomerOutput) ToCustomerOutputWithContext(ctx context.Context) CustomerOutput {
	return o
}

// Address map with fields related to the address: line1, line2, city, state, postalCode and country
func (o CustomerOutput) Address() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringMapOutput { return v.Address }).(pulumi.StringMapOutput)
}

// An integer amount in cents that represents the customer’s current balance, which affect the customer’s future
// invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases
// the amount due on an invoice.
func (o CustomerOutput) Balance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.IntPtrOutput { return v.Balance }).(pulumi.IntPtrOutput)
}

// The default (auto-generated) prefix for the customer used to generate unique invoice numbers.
func (o CustomerOutput) DefaultInvoicePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringOutput { return v.DefaultInvoicePrefix }).(pulumi.StringOutput)
}

// An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
func (o CustomerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching
// and tracking. This may be up to 512 characters.
func (o CustomerOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
func (o CustomerOutput) InvoicePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringPtrOutput { return v.InvoicePrefix }).(pulumi.StringPtrOutput)
}

// Default invoice settings for this customer.
func (o CustomerOutput) InvoiceSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringMapOutput { return v.InvoiceSettings }).(pulumi.StringMapOutput)
}

// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
// object in a structured format.
func (o CustomerOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The customer’s full name or business name.
func (o CustomerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sequence to be used on the customer’s next invoice. Defaults to 1.
func (o CustomerOutput) NextInvoiceSequence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.IntPtrOutput { return v.NextInvoiceSequence }).(pulumi.IntPtrOutput)
}

// The customer’s phone number.
func (o CustomerOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringPtrOutput { return v.Phone }).(pulumi.StringPtrOutput)
}

// Customer’s preferred languages, ordered by preference.
func (o CustomerOutput) PreferredLocales() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringArrayOutput { return v.PreferredLocales }).(pulumi.StringArrayOutput)
}

// Shipping map with fields like name, phone and fields related to the address: line1, line2, city, state, postalCode and
// country.
func (o CustomerOutput) Shipping() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Customer) pulumi.StringMapOutput { return v.Shipping }).(pulumi.StringMapOutput)
}

type CustomerArrayOutput struct{ *pulumi.OutputState }

func (CustomerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Customer)(nil)).Elem()
}

func (o CustomerArrayOutput) ToCustomerArrayOutput() CustomerArrayOutput {
	return o
}

func (o CustomerArrayOutput) ToCustomerArrayOutputWithContext(ctx context.Context) CustomerArrayOutput {
	return o
}

func (o CustomerArrayOutput) Index(i pulumi.IntInput) CustomerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Customer {
		return vs[0].([]*Customer)[vs[1].(int)]
	}).(CustomerOutput)
}

type CustomerMapOutput struct{ *pulumi.OutputState }

func (CustomerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Customer)(nil)).Elem()
}

func (o CustomerMapOutput) ToCustomerMapOutput() CustomerMapOutput {
	return o
}

func (o CustomerMapOutput) ToCustomerMapOutputWithContext(ctx context.Context) CustomerMapOutput {
	return o
}

func (o CustomerMapOutput) MapIndex(k pulumi.StringInput) CustomerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Customer {
		return vs[0].(map[string]*Customer)[vs[1].(string)]
	}).(CustomerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerInput)(nil)).Elem(), &Customer{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerArrayInput)(nil)).Elem(), CustomerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerMapInput)(nil)).Elem(), CustomerMap{})
	pulumi.RegisterOutputType(CustomerOutput{})
	pulumi.RegisterOutputType(CustomerArrayOutput{})
	pulumi.RegisterOutputType(CustomerMapOutput{})
}
