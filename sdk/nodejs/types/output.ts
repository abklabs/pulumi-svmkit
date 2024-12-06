// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace agave {
    export interface Flags {
        accountIndex?: string[];
        accountIndexExcludeKey?: string[];
        accountIndexIncludeKey?: string[];
        accountShrinkPath?: string[];
        accountsDbCacheLimitMb?: number;
        accountsDbTestHashCalculation?: boolean;
        accountsHashCachePath?: string;
        accountsIndexBins?: number;
        accountsIndexPath?: string[];
        accountsIndexScanResultsLimitMb?: number;
        accountsShrinkOptimizeTotalSpace?: boolean;
        accountsShrinkRatio?: string;
        allowPrivateAddr?: boolean;
        authorizedVoter?: string[];
        bindAddress?: string;
        blockProductionMethod?: string;
        blockVerificationMethod?: string;
        checkVoteAccount?: string;
        contactDebugInterval?: number;
        cuda?: boolean;
        debugKey?: string[];
        devHaltAtSlot?: number;
        disableBankingTrace?: boolean;
        dynamicPortRange?: string;
        enableBankingTrace?: number;
        enableBigtableLedgerUpload?: boolean;
        enableExtendedTxMetadataStorage?: boolean;
        enableRpcBigtableLedgerStorage?: boolean;
        enableRpcTransactionHistory?: boolean;
        entryPoint?: string[];
        etcdCacertFile?: string;
        etcdCertFile?: string;
        etcdDomainName?: string;
        etcdEndpoint?: string[];
        etcdKeyFile?: string;
        expectedBankHash?: string;
        expectedGenesisHash?: string;
        expectedShredVersion?: number;
        extraFlags?: string[];
        fullRpcAPI?: boolean;
        fullSnapshotArchivePath?: string;
        fullSnapshotIntervalSlots?: number;
        geyserPluginAlwaysEnabled?: boolean;
        geyserPluginConfig?: string[];
        gossipHost?: string;
        gossipPort?: number;
        gossipValidator?: string[];
        hardFork?: number[];
        healthCheckSlotDistance?: number;
        incrementalSnapshotArchivePath?: string;
        initCompleteFile?: string;
        knownValidator?: string[];
        limitLedgerSize?: number;
        logMessagesBytesLimit?: number;
        maxGenesisArchiveUnpackedSize?: number;
        maximumFullSnapshotsToRetain?: number;
        maximumIncrementalSnapshotsToRetain?: number;
        maximumLocalSnapshotAge?: number;
        maximumSnapshotDownloadAbort?: number;
        minimalSnapshotDownloadSpeed?: number;
        noGenesisFetch?: boolean;
        noIncrementalSnapshots?: boolean;
        noSnapshotFetch?: boolean;
        noVoting?: boolean;
        noWaitForVoteToStartLeader: boolean;
        onlyKnownRPC?: boolean;
        privateRPC?: boolean;
        publicRpcAddress?: string;
        publicTpuAddress?: string;
        publicTpuForwardsAddress?: string;
        repairValidator?: string[];
        requireTower?: boolean;
        restrictedRepairOnlyMode?: boolean;
        rocksdbFifoShredStorageSize?: number;
        rocksdbShredCompaction?: string;
        rpcBigtableAppProfileId?: string;
        rpcBigtableInstanceName?: string;
        rpcBigtableMaxMessageSize?: number;
        rpcBigtableTimeout?: number;
        rpcBindAddress: string;
        rpcFaucetAddress?: string;
        rpcMaxMultipleAccounts?: number;
        rpcMaxRequestBodySize?: number;
        rpcNicenessAdjustment?: number;
        rpcPort: number;
        rpcPubsubEnableBlockSubscription?: boolean;
        rpcPubsubEnableVoteSubscription?: boolean;
        rpcPubsubMaxActiveSubscriptions?: number;
        rpcPubsubNotificationThreads?: number;
        rpcPubsubQueueCapacityBytes?: number;
        rpcPubsubQueueCapacityItems?: number;
        rpcPubsubWorkerThreads?: number;
        rpcScanAndFixRoots?: boolean;
        rpcSendLeaderCount?: number;
        rpcSendRetryMs?: number;
        rpcSendServiceMaxRetries?: number;
        rpcSendTransactionAlsoLeader?: boolean;
        rpcSendTransactionRetryPoolMaxSize?: number;
        rpcSendTransactionTpuPeer?: string[];
        rpcThreads?: number;
        skipPreflightHealthCheck?: boolean;
        skipSeedPhraseValidation?: boolean;
        skipStartupLedgerVerification?: boolean;
        snapshotArchiveFormat?: string;
        snapshotIntervalSlots?: number;
        snapshotPackagerNicenessAdjustment?: number;
        snapshotVersion?: string;
        stakedNodesOverrides?: string;
        towerStorage?: string;
        tpuCoalesceMs?: number;
        tpuConnectionPoolSize?: number;
        tpuDisableQuic?: boolean;
        tpuEnableUdp?: boolean;
        tvuReceiveThreads?: number;
        unifiedSchedulerHandlerThreads?: number;
        useSnapshotArchivesAtStartup?: string;
        waitForSupermajority?: number;
        walRecoveryMode: string;
    }

    export interface KeyPairs {
        identity: string;
        voteAccount: string;
    }

    export interface Metrics {
        database: string;
        password: string;
        url: string;
        user: string;
    }

    export interface ShutdownPolicy {
        force?: boolean;
        maxDelinquentStake?: number;
        minIdleTime?: number;
        skipHealthCheck?: boolean;
        skipNewSnapshotCheck?: boolean;
    }

    export interface StartupPolicy {
        waitForRPCHealth?: boolean;
    }

    export interface TimeoutConfig {
        rpcServiceTimeout?: number;
    }

}

export namespace firedancer {
    export interface Config {
        consensus?: outputs.firedancer.ConfigConsensus;
        dynamicPortRange?: string;
        extraConfig?: string[];
        gossip?: outputs.firedancer.ConfigGossip;
        hugetlbfs?: outputs.firedancer.ConfigHugeTLBFS;
        layout?: outputs.firedancer.ConfigLayout;
        ledger?: outputs.firedancer.ConfigLedger;
        log?: outputs.firedancer.ConfigLog;
        name?: string;
        reporting?: outputs.firedancer.ConfigReporting;
        rpc?: outputs.firedancer.ConfigRPC;
        scratchDirectory?: string;
        snapshots?: outputs.firedancer.ConfigSnapshots;
        user?: string;
    }

    export interface ConfigConsensus {
        authorizedVoterPaths?: string[];
        expectedBankHash?: string;
        expectedGenesisHash?: string;
        expectedShredVersion?: number;
        genesisFetch?: boolean;
        hardForkAtSlots?: string[];
        identityPath?: string;
        knownValidators?: string[];
        osNetworkLimitsTest?: boolean;
        pohSpeedTest?: boolean;
        snapshotFetch?: boolean;
        voteAccountPath?: string;
        waitForSupermajorityAtSlot?: number;
        waitForVoteToStartLeader?: boolean;
    }

    export interface ConfigGossip {
        entrypoints?: string[];
        host?: string;
        port?: number;
        portCheck?: boolean;
    }

    export interface ConfigHugeTLBFS {
        mountPath?: string;
    }

    export interface ConfigLayout {
        affinity?: string;
        agaveAffinity?: string;
        bankTileCount?: number;
        netTileCount?: number;
        quicTileCount?: number;
        resolvTileCount?: number;
        shredTileCount?: number;
        verifyTileCount?: number;
    }

    export interface ConfigLedger {
        accountIndexExcludeKeys?: string[];
        accountIndexIncludeKeys?: string[];
        accountIndexes?: string[];
        accountsPath?: string;
        limitSize?: number;
        path?: string;
        requireTower?: boolean;
        snapshotArchiveFormat?: string;
    }

    export interface ConfigLog {
        colorize?: string;
        levelFlush?: string;
        levelLogfile?: string;
        levelStderr?: string;
        path?: string;
    }

    export interface ConfigRPC {
        bigtableLedgerStorage?: boolean;
        extendedTxMetadataStorage?: boolean;
        fullApi?: boolean;
        onlyKnown?: boolean;
        port?: number;
        private?: boolean;
        pubsubEnableBlockSubscription?: boolean;
        pubsubEnableVoteSubscription?: boolean;
        transactionHistory?: boolean;
    }

    export interface ConfigReporting {
        solanaMetricsConfig?: string;
    }

    export interface ConfigSnapshots {
        fullSnapshotIntervalSlots?: number;
        incrementalPath?: string;
        incrementalSnapshotIntervalSlots?: number;
        incrementalSnapshots?: boolean;
        maximumFullSnapshotsToRetain?: number;
        maximumIncrementalSnapshotsToRetain?: number;
        minimumSnapshotDownloadSpeed?: number;
        path?: string;
    }

    export interface KeyPairs {
        identity: string;
        voteAccount: string;
    }

}

export namespace solana {
    export interface Environment {
        rpcURL: string;
    }

    export interface GenesisFlags {
        clusterType?: string;
        extraFlags?: string[];
        faucetLamports?: string;
        faucetPubkey: string;
        identityPubkey: string;
        inflation?: string;
        lamportsPerByteYear?: string;
        ledgerPath: string;
        slotPerEpoch?: string;
        stakePubkey: string;
        targetLamportsPerSignature?: string;
        votePubkey: string;
    }

    export interface PrimorialEntry {
        lamports: string;
        pubkey: string;
    }

    export interface StakeAccountKeyPairs {
        stakeAccount: string;
        voteAccount: string;
    }

    export interface TxnOptions {
        blockHash?: string;
        commitment?: string;
        feePayer?: string;
        from?: string;
        keyPair?: string;
        nonce?: string;
        nonceAuthority?: string;
        signer?: string[];
        url?: string;
        withComputeUnitPrice?: number;
        withMemo?: string;
        ws?: string;
    }

    export interface ValidatorInfo {
        details?: string;
        iconURL?: string;
        name: string;
        website?: string;
    }

    export interface VoteAccountKeyPairs {
        authWithdrawer: string;
        identity: string;
        voteAccount: string;
    }

}

export namespace ssh {
    /**
     * Instructions for how to connect to a remote endpoint.
     */
    export interface Connection {
        /**
         * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
         */
        agentSocketPath?: string;
        /**
         * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
         */
        dialErrorLimit?: number;
        /**
         * The address of the resource to connect to.
         */
        host: string;
        /**
         * The password we should use for the connection.
         */
        password?: string;
        /**
         * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
         */
        perDialTimeout?: number;
        /**
         * The port to connect to. Defaults to 22.
         */
        port?: number;
        /**
         * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
         */
        privateKey?: string;
        /**
         * The password to use in case the private key is encrypted.
         */
        privateKeyPassword?: string;
        /**
         * The connection settings for the bastion/proxy host.
         */
        proxy?: outputs.ssh.ProxyConnection;
        /**
         * The user that we should use for the connection.
         */
        user?: string;
    }
    /**
     * connectionProvideDefaults sets the appropriate defaults for Connection
     */
    export function connectionProvideDefaults(val: Connection): Connection {
        return {
            ...val,
            dialErrorLimit: (val.dialErrorLimit) ?? 10,
            perDialTimeout: (val.perDialTimeout) ?? 15,
            port: (val.port) ?? 22,
            proxy: (val.proxy ? outputs.ssh.proxyConnectionProvideDefaults(val.proxy) : undefined),
            user: (val.user) ?? "root",
        };
    }

    /**
     * Instructions for how to connect to a remote endpoint via a bastion host.
     */
    export interface ProxyConnection {
        /**
         * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
         */
        agentSocketPath?: string;
        /**
         * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
         */
        dialErrorLimit?: number;
        /**
         * The address of the bastion host to connect to.
         */
        host: string;
        /**
         * The password we should use for the connection to the bastion host.
         */
        password?: string;
        /**
         * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
         */
        perDialTimeout?: number;
        /**
         * The port of the bastion host to connect to.
         */
        port?: number;
        /**
         * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
         */
        privateKey?: string;
        /**
         * The password to use in case the private key is encrypted.
         */
        privateKeyPassword?: string;
        /**
         * The user that we should use for the connection to the bastion host.
         */
        user?: string;
    }
    /**
     * proxyConnectionProvideDefaults sets the appropriate defaults for ProxyConnection
     */
    export function proxyConnectionProvideDefaults(val: ProxyConnection): ProxyConnection {
        return {
            ...val,
            dialErrorLimit: (val.dialErrorLimit) ?? 10,
            perDialTimeout: (val.perDialTimeout) ?? 15,
            port: (val.port) ?? 22,
            user: (val.user) ?? "root",
        };
    }

}
