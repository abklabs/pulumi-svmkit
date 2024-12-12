// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Faucet extends pulumi.CustomResource {
    /**
     * Get an existing Faucet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Faucet {
        return new Faucet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'svmkit:faucet:Faucet';

    /**
     * Returns true if the given object is an instance of Faucet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Faucet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Faucet.__pulumiType;
    }

    public readonly connection!: pulumi.Output<outputs.ssh.Connection>;
    public readonly flags!: pulumi.Output<outputs.solana.FaucetFlags>;
    public readonly keypair!: pulumi.Output<string>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Faucet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FaucetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.flags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flags'");
            }
            if ((!args || args.keypair === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keypair'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(inputs.ssh.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["keypair"] = args?.keypair ? pulumi.secret(args.keypair) : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["keypair"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keypair"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Faucet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Faucet resource.
 */
export interface FaucetArgs {
    connection: pulumi.Input<inputs.ssh.ConnectionArgs>;
    flags: pulumi.Input<inputs.solana.FaucetFlagsArgs>;
    keypair: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}