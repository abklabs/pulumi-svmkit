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
        log?: string;
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

export namespace apt {
    export interface Config {
        excludeDefaultSources?: boolean;
        sources?: outputs.apt.Source[];
    }

    export interface SignedBy {
        paths?: string[];
        publicKey?: string;
    }

    export interface Source {
        URIs: string[];
        allowDowngradeToInsecure?: boolean;
        allowInsecure?: boolean;
        allowWeak?: boolean;
        architectures?: string[];
        checkDate?: boolean;
        checkValidUntil?: boolean;
        components: string[];
        dateMaxFuture?: number;
        extraLines?: string[];
        inReleasePath?: string;
        signedBy?: outputs.apt.SignedBy;
        snapshot?: string;
        suites: string[];
        trusted?: boolean;
        types: string[];
        validUntilMax?: number;
        validUntilMin?: number;
    }

}

export namespace deb {
    export interface Package {
        name: string;
        path?: string;
        targetRelease?: string;
        version?: string;
    }

    export interface PackageConfig {
        additional?: string[];
        override?: outputs.deb.Package[];
    }

}

export namespace explorer {
    export interface ExplorerFlags {
        hostname?: string;
        keepAliveTimeout?: number;
        port?: number;
    }

}

export namespace faucet {
    export interface FaucetFlags {
        allowIPs?: string[];
        perRequestCap?: number;
        perTimeCap?: number;
        sliceSeconds?: number;
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

export namespace firewall {
    export interface FirewallParams {
        allowPorts: string[];
    }

}

export namespace genesis {
    export interface BootstrapAccount {
        balanceLamports?: number;
        identityPubkey: string;
        stakeLamports?: number;
        stakePubkey: string;
        votePubkey: string;
    }

    export interface BootstrapValidator {
        identityPubkey: string;
        stakePubkey: string;
        votePubkey: string;
    }

    export interface GenesisFlags {
        bootstrapStakeAuthorizedPubkey?: string;
        bootstrapValidatorLamports?: number;
        bootstrapValidatorStakeLamports?: number;
        bootstrapValidators: outputs.genesis.BootstrapValidator[];
        clusterType?: string;
        creationTime?: string;
        deactivateFeatures?: string[];
        enableWarmupEpochs?: boolean;
        extraFlags?: string[];
        faucetLamports?: number;
        faucetPubkey?: string;
        feeBurnPercentage?: number;
        hashesPerTick?: string;
        inflation?: string;
        lamportsPerByteYear?: number;
        ledgerPath: string;
        maxGenesisArchiveUnpackedSize?: number;
        rentBurnPercentage?: number;
        rentExemptionThreshold?: number;
        slotsPerEpoch?: number;
        targetLamportsPerSignature?: number;
        targetSignaturesPerSlot?: number;
        targetTickDuration?: number;
        ticksPerSlot?: number;
        url?: string;
        voteCommissionPercentage?: number;
    }

    export interface PrimordialAccount {
        data?: string;
        executable?: boolean;
        lamports: number;
        owner?: string;
        pubkey: string;
    }

}

export namespace geyser {
    export interface Config {
        debugClientsHttp?: boolean;
        grpc: outputs.geyser.GrpcConfigGrpc;
        log?: outputs.geyser.GrpcConfigLog;
        prometheus?: outputs.geyser.GrpcConfigPrometheus;
        tokio?: outputs.geyser.GrpcConfigTokio;
    }

    export interface GeyserPlugin {
        genericPluginConfig?: string;
        yellowstoneGRPC?: outputs.geyser.YellowstoneGRPC;
    }

    export interface GrpcConfigFilterLimits {
        accounts?: outputs.geyser.GrpcConfigFilterLimitsAccounts;
        blocks?: outputs.geyser.GrpcConfigFilterLimitsBlocks;
        blocksMeta?: outputs.geyser.GrpcConfigFilterLimitsBlocksMeta;
        entries?: outputs.geyser.GrpcConfigFilterLimitsEntries;
        slots?: outputs.geyser.GrpcConfigFilterLimitsSlots;
        transactions?: outputs.geyser.GrpcConfigFilterLimitsTransactions;
        transactionsStatus?: outputs.geyser.GrpcConfigFilterLimitsTransactions;
    }

    export interface GrpcConfigFilterLimitsAccounts {
        accountMax?: number;
        accountReject?: string[];
        any?: boolean;
        dataSliceMax?: number;
        max?: number;
        ownerMax?: number;
        ownerReject?: string[];
    }

    export interface GrpcConfigFilterLimitsBlocks {
        accountIncludeAny?: boolean;
        accountIncludeMax?: number;
        accountIncludeReject?: string[];
        includeAccounts?: boolean;
        includeEntries?: boolean;
        includeTransactions?: boolean;
        max?: number;
    }

    export interface GrpcConfigFilterLimitsBlocksMeta {
        max?: number;
    }

    export interface GrpcConfigFilterLimitsEntries {
        max?: number;
    }

    export interface GrpcConfigFilterLimitsSlots {
        max?: number;
    }

    export interface GrpcConfigFilterLimitsTransactions {
        accountExcludeMax?: number;
        accountIncludeMax?: number;
        accountIncludeReject?: string[];
        accountRequiredMax?: number;
        any?: boolean;
        max?: number;
    }

    export interface GrpcConfigGrpc {
        address: string;
        channelCapacity?: number;
        compression?: outputs.geyser.GrpcConfigGrpcCompression;
        filterLimits?: outputs.geyser.GrpcConfigFilterLimits;
        filterNameSizeLimit?: number;
        filterNamesCleanupInterval?: string;
        filterNamesSizeLimit?: number;
        maxDecodingMessageSize?: number;
        replayStoredSlots?: number;
        serverHttp2AdaptiveWindow?: boolean;
        serverHttp2KeepaliveInterval?: string;
        serverHttp2KeepaliveTimeout?: string;
        serverInitialConnectionWindowSize?: number;
        serverInitialStreamWindowSize?: number;
        snapshotClientChannelCapacity?: number;
        snapshotPluginChannelCapacity?: number;
        tlsConfig?: outputs.geyser.GrpcConfigGrpcServerTLS;
        unaryConcurrencyLimit?: number;
        unaryDisabled?: boolean;
        xToken?: string;
    }

    export interface GrpcConfigGrpcCompression {
        accept?: string[];
        send?: string[];
    }

    export interface GrpcConfigGrpcServerTLS {
        certPath: string;
        keyPath: string;
    }

    export interface GrpcConfigLog {
        level?: string;
    }

    export interface GrpcConfigPrometheus {
        address: string;
    }

    export interface GrpcConfigTokio {
        affinity?: number[];
        workerThreads?: number;
    }

    export interface YellowstoneGRPC {
        config?: outputs.geyser.Config;
        json?: string;
        version: string;
    }

}

export namespace runner {
    export interface Config {
        aptLockTimeout?: number;
        keepPayload?: boolean;
        packageConfig?: outputs.deb.PackageConfig;
    }

}

export namespace solana {
    export interface Environment {
        rpcURL: string;
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

export namespace tuner {
    export interface TunerFsParams {
        fsNrOpen?: number;
    }

    export interface TunerKernelParams {
        kernelHungTaskTimeoutSecs?: number;
        kernelNmiWatchdog?: number;
        kernelPidMax?: number;
        kernelSchedMinGranularityNs?: number;
        kernelSchedWakeupGranularityNs?: number;
        kernelTimerMigration?: number;
    }

    export interface TunerNetParams {
        netCoreRmemDefault?: number;
        netCoreRmemMax?: number;
        netCoreWmemDefault?: number;
        netCoreWmemMax?: number;
        netIpv4TcpCongestionControl?: string;
        netIpv4TcpFastopen?: number;
        netIpv4TcpLowLatency?: number;
        netIpv4TcpModerateRcvbuf?: number;
        netIpv4TcpNoMetricsSave?: number;
        netIpv4TcpRmem?: string;
        netIpv4TcpSack?: number;
        netIpv4TcpTimestamps?: number;
        netIpv4TcpTwReuse?: number;
        netIpv4TcpWmem?: string;
    }

    export interface TunerParams {
        cpuGovernor?: enums.tuner.CpuGovernor;
        fs?: outputs.tuner.TunerFsParams;
        kernel?: outputs.tuner.TunerKernelParams;
        net?: outputs.tuner.TunerNetParams;
        vm?: outputs.tuner.TunerVmParams;
    }

    export interface TunerVmParams {
        vmDirtyBackgroundRatio?: number;
        vmDirtyExpireCentisecs?: number;
        vmDirtyRatio?: number;
        vmDirtyWritebackCentisecs?: number;
        vmDirtytimeExpireSeconds?: number;
        vmMaxMapCount?: number;
        vmMinFreeKbytes?: number;
        vmStatInterval?: number;
        vmSwappiness?: number;
    }

}

export namespace watchtower {
    export interface DiscordConfig {
        webhookUrl: string;
    }

    export interface NotificationConfig {
        discord?: outputs.watchtower.DiscordConfig;
        pagerDuty?: outputs.watchtower.PagerDutyConfig;
        slack?: outputs.watchtower.SlackConfig;
        telegram?: outputs.watchtower.TelegramConfig;
        twilio?: outputs.watchtower.TwilioConfig;
    }

    export interface PagerDutyConfig {
        integrationKey: string;
    }

    export interface SlackConfig {
        webhookUrl: string;
    }

    export interface TelegramConfig {
        botToken: string;
        chatId: string;
    }

    export interface TwilioConfig {
        accountSid: string;
        authToken: string;
        fromNumber: string;
        toNumber: string;
    }

    export interface WatchtowerFlags {
        activeStakeAlertThreshold?: number;
        ignoreHttpBadGateway?: boolean;
        interval?: number;
        minimumValidatorIdentityBalance?: number;
        monitorActiveStake?: boolean;
        nameSuffix?: string;
        rpcTimeout?: number;
        unhealthyThreshold?: number;
        validatorIdentity: string[];
    }

}
