// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Explorer extends pulumi.CustomResource {
    /**
     * Get an existing Explorer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Explorer {
        return new Explorer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'svmkit:explorer:Explorer';

    /**
     * Returns true if the given object is an instance of Explorer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Explorer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Explorer.__pulumiType;
    }

    public readonly RPCURL!: pulumi.Output<string | undefined>;
    public readonly clusterName!: pulumi.Output<string | undefined>;
    public readonly connection!: pulumi.Output<outputs.ssh.Connection>;
    public readonly environment!: pulumi.Output<outputs.solana.Environment>;
    public readonly flags!: pulumi.Output<outputs.explorer.ExplorerFlags>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly runnerConfig!: pulumi.Output<outputs.runner.Config | undefined>;
    public readonly symbol!: pulumi.Output<string | undefined>;
    public readonly triggers!: pulumi.Output<any[] | undefined>;
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Explorer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExplorerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.flags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flags'");
            }
            resourceInputs["RPCURL"] = args ? args.RPCURL : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(inputs.ssh.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["runnerConfig"] = args ? args.runnerConfig : undefined;
            resourceInputs["symbol"] = args ? args.symbol : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        } else {
            resourceInputs["RPCURL"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["runnerConfig"] = undefined /*out*/;
            resourceInputs["symbol"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Explorer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Explorer resource.
 */
export interface ExplorerArgs {
    RPCURL?: pulumi.Input<string>;
    clusterName?: pulumi.Input<string>;
    connection: pulumi.Input<inputs.ssh.ConnectionArgs>;
    environment: pulumi.Input<inputs.solana.EnvironmentArgs>;
    flags: pulumi.Input<inputs.explorer.ExplorerFlagsArgs>;
    name?: pulumi.Input<string>;
    runnerConfig?: pulumi.Input<inputs.runner.ConfigArgs>;
    symbol?: pulumi.Input<string>;
    triggers?: pulumi.Input<any[]>;
    version?: pulumi.Input<string>;
}
