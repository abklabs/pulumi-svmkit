// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Validator extends pulumi.CustomResource {
    /**
     * Get an existing Validator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Validator {
        return new Validator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'svm:svm:Validator';

    /**
     * Returns true if the given object is an instance of Validator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Validator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Validator.__pulumiType;
    }

    public readonly flags!: pulumi.Output<outputs.agave.ValidatorFlags>;
    public readonly keyPairs!: pulumi.Output<outputs.agave.ValidatorKeyPairs>;
    public readonly variant!: pulumi.Output<string | undefined>;

    /**
     * Create a Validator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ValidatorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.flags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flags'");
            }
            if ((!args || args.keyPairs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyPairs'");
            }
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["keyPairs"] = args?.keyPairs ? pulumi.secret(args.keyPairs) : undefined;
            resourceInputs["variant"] = args ? args.variant : undefined;
        } else {
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["keyPairs"] = undefined /*out*/;
            resourceInputs["variant"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keyPairs"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Validator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Validator resource.
 */
export interface ValidatorArgs {
    flags: pulumi.Input<inputs.agave.ValidatorFlagsArgs>;
    keyPairs: pulumi.Input<inputs.agave.ValidatorKeyPairsArgs>;
    variant?: pulumi.Input<string>;
}
