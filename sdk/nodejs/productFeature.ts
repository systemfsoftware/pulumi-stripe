// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ProductFeature extends pulumi.CustomResource {
    /**
     * Get an existing ProductFeature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductFeatureState, opts?: pulumi.CustomResourceOptions): ProductFeature {
        return new ProductFeature(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/productFeature:ProductFeature';

    /**
     * Returns true if the given object is an instance of ProductFeature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductFeature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductFeature.__pulumiType;
    }

    /**
     * The ID of the Entitlements Feature the product will be attached to
     */
    public readonly entitlementsFeature!: pulumi.Output<string>;
    /**
     * Has the value true if the object exists in live mode or the value false if the object exists in test mode
     */
    public /*out*/ readonly livemode!: pulumi.Output<boolean>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    public /*out*/ readonly object!: pulumi.Output<string>;
    /**
     * The ID of the product that this Entitlements Feature will be attached to.
     */
    public readonly product!: pulumi.Output<string>;

    /**
     * Create a ProductFeature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductFeatureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductFeatureArgs | ProductFeatureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProductFeatureState | undefined;
            resourceInputs["entitlementsFeature"] = state ? state.entitlementsFeature : undefined;
            resourceInputs["livemode"] = state ? state.livemode : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["product"] = state ? state.product : undefined;
        } else {
            const args = argsOrState as ProductFeatureArgs | undefined;
            if ((!args || args.entitlementsFeature === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entitlementsFeature'");
            }
            if ((!args || args.product === undefined) && !opts.urn) {
                throw new Error("Missing required property 'product'");
            }
            resourceInputs["entitlementsFeature"] = args ? args.entitlementsFeature : undefined;
            resourceInputs["product"] = args ? args.product : undefined;
            resourceInputs["livemode"] = undefined /*out*/;
            resourceInputs["object"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProductFeature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductFeature resources.
 */
export interface ProductFeatureState {
    /**
     * The ID of the Entitlements Feature the product will be attached to
     */
    entitlementsFeature?: pulumi.Input<string>;
    /**
     * Has the value true if the object exists in live mode or the value false if the object exists in test mode
     */
    livemode?: pulumi.Input<boolean>;
    /**
     * String representing the object’s type. Objects of the same type share the same value.
     */
    object?: pulumi.Input<string>;
    /**
     * The ID of the product that this Entitlements Feature will be attached to.
     */
    product?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductFeature resource.
 */
export interface ProductFeatureArgs {
    /**
     * The ID of the Entitlements Feature the product will be attached to
     */
    entitlementsFeature: pulumi.Input<string>;
    /**
     * The ID of the product that this Entitlements Feature will be attached to.
     */
    product: pulumi.Input<string>;
}
