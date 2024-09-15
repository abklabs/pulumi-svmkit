export const RPC_PORT = 8899;
export const GOSSIP_PORT = 8001;

export type ValidatorPaths = {
    accounts: string;
    ledger: string;
    log: string;
};

export type ValidatorFlags = {
    entryPoint?: string[];
    knownValidator?: string[];
    useSnapshotArchivesAtStartup: string;
    rpcPort: number;
    privateRPC: boolean;
    onlyKnownRPC: boolean;
    dynamicPortRange: string;
    gossipPort: number;
    rpcBindAddress: string;
    walRecoveryMode: string;
    limitLedgerSize: number;
    blockProductionMethod: string;
    tvuReceiveThreads?: number;
    noWaitForVoteToStartLeader: boolean;
    fullSnapshotIntervalSlots: number;
    expectedGenesisHash?: string;
    fullRpcAPI?: boolean;
    noVoting?: boolean;
    paths: ValidatorPaths;
};

export type ValidatorKeyPairs = {
    identity: string;
    voteAccount: string;
    authorizedWithdrawer: string;
    stakeAccount?: string;
};

export function flags(flags: ValidatorFlags) {
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

    (flags.entryPoint || []).forEach((entrypoint) => {
        s("entrypoint", entrypoint);
    });

    (flags.knownValidator || []).forEach((knownValidator) => {
        s("known-validator", knownValidator);
    });

    s("expected-genesis-hash", flags.expectedGenesisHash);
    s("use-snapshot-archives-at-startup", flags.useSnapshotArchivesAtStartup);
    s("rpc-port", flags.rpcPort);
    s("dynamic-port-range", flags.dynamicPortRange);
    s("gossip-port", flags.gossipPort);
    s("rpc-bind-address", flags.rpcBindAddress);
    s("wal-recovery-mode", flags.walRecoveryMode);
    s("log", flags.paths.log);
    s("accounts", flags.paths.accounts);
    s("ledger", flags.paths.ledger);
    s("limit-ledger-size", flags.limitLedgerSize);
    s("block-production-method", flags.blockProductionMethod);
    s("tvu-receive-threads", flags.tvuReceiveThreads);
    s("full-snapshot-interval-slots", flags.fullSnapshotIntervalSlots);

    b("no-wait-for-vote-to-start-leader", flags.noWaitForVoteToStartLeader);
    b("only-known-rpc", flags.onlyKnownRPC);
    b("private-rpc", flags.privateRPC);
    b("full-rpc-api", flags.fullRpcAPI);
    b("no-voting", flags.noVoting);

    return l;
}
