// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";

import { constructValidator } from "./validator";
import { constructKeyPair } from "./keypair";
import { constructGenesis } from "./genesis";
import { constructStake } from "./stake";
import { constructReady } from "./ready";
import { constructVote } from "./vote";

export { RPC_PORT, GOSSIP_PORT } from "./agave";

export class Provider implements provider.Provider {
    constructor(
        readonly version: string,
        readonly schema: string,
    ) {}

    async construct(
        name: string,
        type: string,
        inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions,
    ): Promise<provider.ConstructResult> {
        switch (type) {
            case "svmkit:index:Validator":
                return await constructValidator(name, inputs, options);
            case "svmkit:index:KeyPair":
                return await constructKeyPair(name, inputs, options);
            case "svmkit:index:Genesis":
                return await constructGenesis(name, inputs, options);
            case "svmkit:index:Stake":
                return await constructStake(name, inputs, options);
            case "svmkit:index:Ready":
                return await constructReady(name, inputs, options);
            case "svmkit:index:Vote":
                return await constructVote(name, inputs, options);
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    }
}
