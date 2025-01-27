// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Solana extends pulumi.CustomResource {
    /**
     * Get an existing Solana resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Solana {
        return new Solana(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'svmkit:genesis:Solana';

    /**
     * Returns true if the given object is an instance of Solana.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Solana {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Solana.__pulumiType;
    }

    public readonly connection!: pulumi.Output<outputs.ssh.Connection>;
    public readonly flags!: pulumi.Output<outputs.genesis.GenesisFlags>;
    public /*out*/ readonly genesisHash!: pulumi.Output<string>;
    public readonly primordial!: pulumi.Output<outputs.genesis.PrimorialEntry[]>;
    public readonly runnerConfig!: pulumi.Output<outputs.runner.Config | undefined>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Solana resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SolanaArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.flags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flags'");
            }
            if ((!args || args.primordial === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primordial'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(inputs.ssh.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["primordial"] = args ? args.primordial : undefined;
            resourceInputs["runnerConfig"] = args ? args.runnerConfig : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["genesisHash"] = undefined /*out*/;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["genesisHash"] = undefined /*out*/;
            resourceInputs["primordial"] = undefined /*out*/;
            resourceInputs["runnerConfig"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Solana.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Solana resource.
 */
export interface SolanaArgs {
    connection: pulumi.Input<inputs.ssh.ConnectionArgs>;
    flags: pulumi.Input<inputs.genesis.GenesisFlagsArgs>;
    primordial: pulumi.Input<pulumi.Input<inputs.genesis.PrimorialEntryArgs>[]>;
    runnerConfig?: pulumi.Input<inputs.runner.ConfigArgs>;
    version?: pulumi.Input<string>;
}
