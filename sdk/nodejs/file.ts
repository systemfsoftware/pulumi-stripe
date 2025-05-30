// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class File extends pulumi.CustomResource {
    /**
     * Get an existing File resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileState, opts?: pulumi.CustomResourceOptions): File {
        return new File(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/file:File';

    /**
     * Returns true if the given object is an instance of File.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is File {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === File.__pulumiType;
    }

    /**
     * A content file to upload encoded by base64.
     */
    public readonly base64content!: pulumi.Output<string>;
    /**
     * Time at which the object was created. Measured in seconds since the Unix epoch.
     */
    public /*out*/ readonly created!: pulumi.Output<number>;
    /**
     * The file expires and isn’t available at this time in epoch seconds.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<number>;
    /**
     * The suitable name for saving the file to a filesystem.
     */
    public readonly filename!: pulumi.Output<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    public readonly linkData!: pulumi.Output<outputs.FileLinkData | undefined>;
    public /*out*/ readonly links!: pulumi.Output<outputs.FileLink[]>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    public /*out*/ readonly object!: pulumi.Output<string>;
    /**
     * The purpose of the uploaded file.
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * The size of the file object in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The returned file type (for example, csv, pdf, jpg, or png).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Use your live secret API key to download the file from this URL.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a File resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileArgs | FileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileState | undefined;
            resourceInputs["base64content"] = state ? state.base64content : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["linkData"] = state ? state.linkData : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["purpose"] = state ? state.purpose : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if ((!args || args.base64content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'base64content'");
            }
            if ((!args || args.filename === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filename'");
            }
            if ((!args || args.purpose === undefined) && !opts.urn) {
                throw new Error("Missing required property 'purpose'");
            }
            resourceInputs["base64content"] = args ? args.base64content : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["linkData"] = args ? args.linkData : undefined;
            resourceInputs["purpose"] = args ? args.purpose : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["object"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(File.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering File resources.
 */
export interface FileState {
    /**
     * A content file to upload encoded by base64.
     */
    base64content?: pulumi.Input<string>;
    /**
     * Time at which the object was created. Measured in seconds since the Unix epoch.
     */
    created?: pulumi.Input<number>;
    /**
     * The file expires and isn’t available at this time in epoch seconds.
     */
    expiresAt?: pulumi.Input<number>;
    /**
     * The suitable name for saving the file to a filesystem.
     */
    filename?: pulumi.Input<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    linkData?: pulumi.Input<inputs.FileLinkData>;
    links?: pulumi.Input<pulumi.Input<inputs.FileLink>[]>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    object?: pulumi.Input<string>;
    /**
     * The purpose of the uploaded file.
     */
    purpose?: pulumi.Input<string>;
    /**
     * The size of the file object in bytes.
     */
    size?: pulumi.Input<number>;
    /**
     * The returned file type (for example, csv, pdf, jpg, or png).
     */
    type?: pulumi.Input<string>;
    /**
     * Use your live secret API key to download the file from this URL.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a File resource.
 */
export interface FileArgs {
    /**
     * A content file to upload encoded by base64.
     */
    base64content: pulumi.Input<string>;
    /**
     * The suitable name for saving the file to a filesystem.
     */
    filename: pulumi.Input<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    linkData?: pulumi.Input<inputs.FileLinkData>;
    /**
     * The purpose of the uploaded file.
     */
    purpose: pulumi.Input<string>;
}
