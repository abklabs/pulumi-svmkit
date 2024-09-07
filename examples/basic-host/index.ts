import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as command from "@pulumi/command";
import * as fs from "fs";

const config = new pulumi.Config();

const host = config.require("host");
const user = config.require("user");
const sshKeyFile = config.requireSecret("ssh-key-file");

const sshPrivateKey = sshKeyFile.apply(filePath => fs.readFileSync(filePath, 'utf8'));

const validatorKey = new svmkit.KeyPair("validatorKey");
const voteAccountKey = new svmkit.KeyPair("voteAccount");

const validatorConfig = {
    identityKeyPair: validatorKey.json,
    voteAccountKeyPair: voteAccountKey.json,
    entryPoint: [
        "entrypoint.testnet.solana.com:8001",
        "entrypoint2.testnet.solana.com:8001",
        "entrypoint3.testnet.solana.com:8001",
    ],
    knownValidator: [
        "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
        "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
        "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
        "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
    ],
    expectedGenesisHash: "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
    useSnapshotArchivesAtStartup: "when-newest",
    rpcPort: 8899,
    privateRPC: true,
    onlyKnownRPC: true,
    dynamicPortRange: "8002-8020",
    gossipPort: 8001,
    rpcBindAddress: "0.0.0.0",
    walRecoveryMode: "skip_any_corrupted_record",
    limitLedgerSize: 50000000,
    blockProductionMethod: "central-scheduler",
    tvuReceiveThreads: 2,
    paths: {
        accounts: "/home/sol/accounts",
        ledger: "/home/sol/ledger",
        log: "/home/sol/log",
    },
};

const sshConnInfo: command.types.input.remote.ConnectionArgs = {
    host: host,
    user: user,
    privateKey: sshPrivateKey,
};

const backendSetup = new svmkit.Backend("backendSetup", {
    connection: sshConnInfo,
    validatorConfig,
});

export const HOST = host;
export const USER = user;
