// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetDefaultTunerParamsArgs, GetDefaultTunerParamsResult, GetDefaultTunerParamsOutputArgs } from "./getDefaultTunerParams";
export const getDefaultTunerParams: typeof import("./getDefaultTunerParams").getDefaultTunerParams = null as any;
export const getDefaultTunerParamsOutput: typeof import("./getDefaultTunerParams").getDefaultTunerParamsOutput = null as any;
utilities.lazyLoad(exports, ["getDefaultTunerParams","getDefaultTunerParamsOutput"], () => require("./getDefaultTunerParams"));

export { TunerArgs } from "./tuner";
export type Tuner = import("./tuner").Tuner;
export const Tuner: typeof import("./tuner").Tuner = null as any;
utilities.lazyLoad(exports, ["Tuner"], () => require("./tuner"));


// Export enums:
export * from "../types/enums/tuner";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "svmkit:tuner:Tuner":
                return new Tuner(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("svmkit", "tuner", _module)
