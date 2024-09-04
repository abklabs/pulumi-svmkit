import { BackendValidatorConfig } from "./backend";

export function generateAgaveValidatorFlags(config: BackendValidatorConfig) {
    const l: string[] = [];

    const b = (k: string, v: any) => {
        if (v) {
            l.push(`--${k}`);
        }
    };

    const s = (k: string, v: string | number | undefined) => {
        if (v !== undefined) {
            l.push(`--${k}`, `${v}`);
        }
    };

    // Note: These locations are hard coded inside asset-builder.
    s("identity", "/home/sol/validator-keypair.json");
    s("vote-account", "/home/sol/vote-account-keypair.json");

    if (config.entryPoint) {
        config.entryPoint.forEach((v) => s("entrypoint", v));
    }

    if (config.knownValidator) {
        config.knownValidator.forEach((v) => s("known-validator", v));
    }

    s("use-snapshot-archives-at-startup", config.useSnapshotArchivesAtStartup);
    s("rpc-port", config.rpcPort);
    s("dynamic-port-range", config.dynamicPortRange);
    s("gossip-port", config.gossipPort);
    s("rpc-bind-address", config.rpcBindAddress);
    s("wal-recovery-mode", config.walRecoveryMode);
    s("log", config.paths.log);
    s("accounts", config.paths.accounts);
    s("ledger", config.paths.ledger);
    s("limit-ledger-size", config.limitLedgerSize);
    s("block-production-method", config.blockProductionMethod);
    s("tvu-receive-threads", config.tvuReceiveThreads);
    s("full-snapshot-interval-slots", config.fullSnapshotIntervalSlots);

    b("no-wait-for-vote-to-start-leader", config.noWaitForVoteToStartLeader);
    b("only-known-rpc", config.onlyKnownRPC);
    b("private-rpc", config.privateRPC);

    return l;
}
