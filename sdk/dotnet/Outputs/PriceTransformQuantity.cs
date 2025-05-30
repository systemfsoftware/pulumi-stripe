// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Outputs
{

    [OutputType]
    public sealed class PriceTransformQuantity
    {
        /// <summary>
        /// Divide usage by this number.
        /// </summary>
        public readonly int DivideBy;
        /// <summary>
        /// After division, either round the result up or down
        /// </summary>
        public readonly string Round;

        [OutputConstructor]
        private PriceTransformQuantity(
            int divideBy,

            string round)
        {
            DivideBy = divideBy;
            Round = round;
        }
    }
}
