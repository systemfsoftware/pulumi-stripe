// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EntitlementsFeature extends pulumi.CustomResource {
    /**
     * Get an existing EntitlementsFeature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntitlementsFeatureState, opts?: pulumi.CustomResourceOptions): EntitlementsFeature {
        return new EntitlementsFeature(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/entitlementsFeature:EntitlementsFeature';

    /**
     * Returns true if the given object is an instance of EntitlementsFeature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntitlementsFeature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntitlementsFeature.__pulumiType;
    }

    /**
     * Inactive features cannot be attached to new products and will not be returned from the features list endpoint.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Has the value true if the object exists in live mode or the value false if the object exists in test mode
     */
    public /*out*/ readonly livemode!: pulumi.Output<boolean>;
    /**
     * A unique key you provide as your own system identifier. This may be up to 80 characters.
     */
    public readonly lookupKey!: pulumi.Output<string>;
    /**
     * Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
     * object in a structured format.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The feature’s name, for your own purpose, not meant to be displayable to the customer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    public /*out*/ readonly object!: pulumi.Output<string>;

    /**
     * Create a EntitlementsFeature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntitlementsFeatureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntitlementsFeatureArgs | EntitlementsFeatureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntitlementsFeatureState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["livemode"] = state ? state.livemode : undefined;
            resourceInputs["lookupKey"] = state ? state.lookupKey : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
        } else {
            const args = argsOrState as EntitlementsFeatureArgs | undefined;
            if ((!args || args.lookupKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lookupKey'");
            }
            resourceInputs["lookupKey"] = args ? args.lookupKey : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["livemode"] = undefined /*out*/;
            resourceInputs["object"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntitlementsFeature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EntitlementsFeature resources.
 */
export interface EntitlementsFeatureState {
    /**
     * Inactive features cannot be attached to new products and will not be returned from the features list endpoint.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Has the value true if the object exists in live mode or the value false if the object exists in test mode
     */
    livemode?: pulumi.Input<boolean>;
    /**
     * A unique key you provide as your own system identifier. This may be up to 80 characters.
     */
    lookupKey?: pulumi.Input<string>;
    /**
     * Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
     * object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The feature’s name, for your own purpose, not meant to be displayable to the customer.
     */
    name?: pulumi.Input<string>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    object?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EntitlementsFeature resource.
 */
export interface EntitlementsFeatureArgs {
    /**
     * A unique key you provide as your own system identifier. This may be up to 80 characters.
     */
    lookupKey: pulumi.Input<string>;
    /**
     * Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the
     * object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The feature’s name, for your own purpose, not meant to be displayable to the customer.
     */
    name?: pulumi.Input<string>;
}
