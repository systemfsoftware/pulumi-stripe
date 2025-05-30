// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe
{
    [StripeResourceType("stripe:index/meter:Meter")]
    public partial class Meter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Fields that specify how to map a meter event to a customer.
        /// </summary>
        [Output("customerMapping")]
        public Output<Outputs.MeterCustomerMapping?> CustomerMapping { get; private set; } = null!;

        /// <summary>
        /// The default settings to aggregate a meter’s events with
        /// </summary>
        [Output("defaultAggregation")]
        public Output<Outputs.MeterDefaultAggregation> DefaultAggregation { get; private set; } = null!;

        /// <summary>
        /// The meter’s name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the meter event to record usage for. Corresponds with the event_name field on meter events
        /// </summary>
        [Output("eventName")]
        public Output<string> EventName { get; private set; } = null!;

        /// <summary>
        /// The time window to pre-aggregate meter events for, if any.
        /// </summary>
        [Output("eventTimeWindow")]
        public Output<string?> EventTimeWindow { get; private set; } = null!;

        /// <summary>
        /// Fields that specify how to calculate a meter event’s value.
        /// </summary>
        [Output("valueSettings")]
        public Output<Outputs.MeterValueSettings?> ValueSettings { get; private set; } = null!;


        /// <summary>
        /// Create a Meter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Meter(string name, MeterArgs args, CustomResourceOptions? options = null)
            : base("stripe:index/meter:Meter", name, args ?? new MeterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Meter(string name, Input<string> id, MeterState? state = null, CustomResourceOptions? options = null)
            : base("stripe:index/meter:Meter", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Meter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Meter Get(string name, Input<string> id, MeterState? state = null, CustomResourceOptions? options = null)
        {
            return new Meter(name, id, state, options);
        }
    }

    public sealed class MeterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fields that specify how to map a meter event to a customer.
        /// </summary>
        [Input("customerMapping")]
        public Input<Inputs.MeterCustomerMappingArgs>? CustomerMapping { get; set; }

        /// <summary>
        /// The default settings to aggregate a meter’s events with
        /// </summary>
        [Input("defaultAggregation", required: true)]
        public Input<Inputs.MeterDefaultAggregationArgs> DefaultAggregation { get; set; } = null!;

        /// <summary>
        /// The meter’s name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The name of the meter event to record usage for. Corresponds with the event_name field on meter events
        /// </summary>
        [Input("eventName", required: true)]
        public Input<string> EventName { get; set; } = null!;

        /// <summary>
        /// The time window to pre-aggregate meter events for, if any.
        /// </summary>
        [Input("eventTimeWindow")]
        public Input<string>? EventTimeWindow { get; set; }

        /// <summary>
        /// Fields that specify how to calculate a meter event’s value.
        /// </summary>
        [Input("valueSettings")]
        public Input<Inputs.MeterValueSettingsArgs>? ValueSettings { get; set; }

        public MeterArgs()
        {
        }
        public static new MeterArgs Empty => new MeterArgs();
    }

    public sealed class MeterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fields that specify how to map a meter event to a customer.
        /// </summary>
        [Input("customerMapping")]
        public Input<Inputs.MeterCustomerMappingGetArgs>? CustomerMapping { get; set; }

        /// <summary>
        /// The default settings to aggregate a meter’s events with
        /// </summary>
        [Input("defaultAggregation")]
        public Input<Inputs.MeterDefaultAggregationGetArgs>? DefaultAggregation { get; set; }

        /// <summary>
        /// The meter’s name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the meter event to record usage for. Corresponds with the event_name field on meter events
        /// </summary>
        [Input("eventName")]
        public Input<string>? EventName { get; set; }

        /// <summary>
        /// The time window to pre-aggregate meter events for, if any.
        /// </summary>
        [Input("eventTimeWindow")]
        public Input<string>? EventTimeWindow { get; set; }

        /// <summary>
        /// Fields that specify how to calculate a meter event’s value.
        /// </summary>
        [Input("valueSettings")]
        public Input<Inputs.MeterValueSettingsGetArgs>? ValueSettings { get; set; }

        public MeterState()
        {
        }
        public static new MeterState Empty => new MeterState();
    }
}
