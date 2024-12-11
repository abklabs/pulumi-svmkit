// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { KeyPairArgs } from "./keyPair";
export type KeyPair = import("./keyPair").KeyPair;
export const KeyPair: typeof import("./keyPair").KeyPair = null as any;
utilities.lazyLoad(exports, ["KeyPair"], () => require("./keyPair"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as account from "./account";
import * as agave from "./agave";
import * as faucet from "./faucet";
import * as genesis from "./genesis";
import * as types from "./types";
import * as validator from "./validator";

export {
    account,
    agave,
    faucet,
    genesis,
    types,
    validator,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "svmkit:index:KeyPair":
                return new KeyPair(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("svmkit", "index", _module)
pulumi.runtime.registerResourcePackage("svmkit", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:svmkit") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
