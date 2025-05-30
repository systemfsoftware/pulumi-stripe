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
    public sealed class FileLinkData
    {
        /// <summary>
        /// Set this to true to create a file link for the newly created file. Creating a link is only possible when the file’s purpose is one of the following: business_icon, business_logo, customer_signature, dispute_evidence, pci_document, tax_document_user_upload, or terminal_reader_splashscreen.
        /// </summary>
        public readonly bool Create;
        /// <summary>
        /// The link isn’t available after this future timestamp.
        /// </summary>
        public readonly int? ExpiresAt;
        /// <summary>
        /// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Metadata;

        [OutputConstructor]
        private FileLinkData(
            bool create,

            int? expiresAt,

            ImmutableDictionary<string, string>? metadata)
        {
            Create = create;
            ExpiresAt = expiresAt;
            Metadata = metadata;
        }
    }
}
