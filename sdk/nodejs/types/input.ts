// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace agave {
    export interface FlagsArgs {
        accountIndex?: pulumi.Input<pulumi.Input<string>[]>;
        accountIndexExcludeKey?: pulumi.Input<pulumi.Input<string>[]>;
        accountIndexIncludeKey?: pulumi.Input<pulumi.Input<string>[]>;
        accountShrinkPath?: pulumi.Input<pulumi.Input<string>[]>;
        accountsDbCacheLimitMb?: pulumi.Input<number>;
        accountsDbTestHashCalculation?: pulumi.Input<boolean>;
        accountsHashCachePath?: pulumi.Input<string>;
        accountsIndexBins?: pulumi.Input<number>;
        accountsIndexPath?: pulumi.Input<pulumi.Input<string>[]>;
        accountsIndexScanResultsLimitMb?: pulumi.Input<number>;
        accountsShrinkOptimizeTotalSpace?: pulumi.Input<boolean>;
        accountsShrinkRatio?: pulumi.Input<string>;
        allowPrivateAddr?: pulumi.Input<boolean>;
        authorizedVoter?: pulumi.Input<pulumi.Input<string>[]>;
        bindAddress?: pulumi.Input<string>;
        blockProductionMethod?: pulumi.Input<string>;
        blockVerificationMethod?: pulumi.Input<string>;
        checkVoteAccount?: pulumi.Input<string>;
        contactDebugInterval?: pulumi.Input<number>;
        cuda?: pulumi.Input<boolean>;
        debugKey?: pulumi.Input<pulumi.Input<string>[]>;
        devHaltAtSlot?: pulumi.Input<number>;
        disableBankingTrace?: pulumi.Input<boolean>;
        dynamicPortRange?: pulumi.Input<string>;
        enableBankingTrace?: pulumi.Input<number>;
        enableBigtableLedgerUpload?: pulumi.Input<boolean>;
        enableExtendedTxMetadataStorage?: pulumi.Input<boolean>;
        enableRpcBigtableLedgerStorage?: pulumi.Input<boolean>;
        enableRpcTransactionHistory?: pulumi.Input<boolean>;
        entryPoint?: pulumi.Input<pulumi.Input<string>[]>;
        etcdCacertFile?: pulumi.Input<string>;
        etcdCertFile?: pulumi.Input<string>;
        etcdDomainName?: pulumi.Input<string>;
        etcdEndpoint?: pulumi.Input<pulumi.Input<string>[]>;
        etcdKeyFile?: pulumi.Input<string>;
        expectedBankHash?: pulumi.Input<string>;
        expectedGenesisHash?: pulumi.Input<string>;
        expectedShredVersion?: pulumi.Input<number>;
        extraFlags?: pulumi.Input<pulumi.Input<string>[]>;
        fullRpcAPI?: pulumi.Input<boolean>;
        fullSnapshotArchivePath?: pulumi.Input<string>;
        fullSnapshotIntervalSlots?: pulumi.Input<number>;
        geyserPluginAlwaysEnabled?: pulumi.Input<boolean>;
        geyserPluginConfig?: pulumi.Input<pulumi.Input<string>[]>;
        gossipHost?: pulumi.Input<string>;
        gossipPort?: pulumi.Input<number>;
        gossipValidator?: pulumi.Input<pulumi.Input<string>[]>;
        hardFork?: pulumi.Input<pulumi.Input<number>[]>;
        healthCheckSlotDistance?: pulumi.Input<number>;
        incrementalSnapshotArchivePath?: pulumi.Input<string>;
        initCompleteFile?: pulumi.Input<string>;
        knownValidator?: pulumi.Input<pulumi.Input<string>[]>;
        limitLedgerSize?: pulumi.Input<number>;
        log?: pulumi.Input<string>;
        logMessagesBytesLimit?: pulumi.Input<number>;
        maxGenesisArchiveUnpackedSize?: pulumi.Input<number>;
        maximumFullSnapshotsToRetain?: pulumi.Input<number>;
        maximumIncrementalSnapshotsToRetain?: pulumi.Input<number>;
        maximumLocalSnapshotAge?: pulumi.Input<number>;
        maximumSnapshotDownloadAbort?: pulumi.Input<number>;
        minimalSnapshotDownloadSpeed?: pulumi.Input<number>;
        noGenesisFetch?: pulumi.Input<boolean>;
        noIncrementalSnapshots?: pulumi.Input<boolean>;
        noSnapshotFetch?: pulumi.Input<boolean>;
        noVoting?: pulumi.Input<boolean>;
        noWaitForVoteToStartLeader: pulumi.Input<boolean>;
        onlyKnownRPC?: pulumi.Input<boolean>;
        privateRPC?: pulumi.Input<boolean>;
        publicRpcAddress?: pulumi.Input<string>;
        publicTpuAddress?: pulumi.Input<string>;
        publicTpuForwardsAddress?: pulumi.Input<string>;
        repairValidator?: pulumi.Input<pulumi.Input<string>[]>;
        requireTower?: pulumi.Input<boolean>;
        restrictedRepairOnlyMode?: pulumi.Input<boolean>;
        rocksdbFifoShredStorageSize?: pulumi.Input<number>;
        rocksdbShredCompaction?: pulumi.Input<string>;
        rpcBigtableAppProfileId?: pulumi.Input<string>;
        rpcBigtableInstanceName?: pulumi.Input<string>;
        rpcBigtableMaxMessageSize?: pulumi.Input<number>;
        rpcBigtableTimeout?: pulumi.Input<number>;
        rpcBindAddress: pulumi.Input<string>;
        rpcFaucetAddress?: pulumi.Input<string>;
        rpcMaxMultipleAccounts?: pulumi.Input<number>;
        rpcMaxRequestBodySize?: pulumi.Input<number>;
        rpcNicenessAdjustment?: pulumi.Input<number>;
        rpcPort: pulumi.Input<number>;
        rpcPubsubEnableBlockSubscription?: pulumi.Input<boolean>;
        rpcPubsubEnableVoteSubscription?: pulumi.Input<boolean>;
        rpcPubsubMaxActiveSubscriptions?: pulumi.Input<number>;
        rpcPubsubNotificationThreads?: pulumi.Input<number>;
        rpcPubsubQueueCapacityBytes?: pulumi.Input<number>;
        rpcPubsubQueueCapacityItems?: pulumi.Input<number>;
        rpcPubsubWorkerThreads?: pulumi.Input<number>;
        rpcScanAndFixRoots?: pulumi.Input<boolean>;
        rpcSendLeaderCount?: pulumi.Input<number>;
        rpcSendRetryMs?: pulumi.Input<number>;
        rpcSendServiceMaxRetries?: pulumi.Input<number>;
        rpcSendTransactionAlsoLeader?: pulumi.Input<boolean>;
        rpcSendTransactionRetryPoolMaxSize?: pulumi.Input<number>;
        rpcSendTransactionTpuPeer?: pulumi.Input<pulumi.Input<string>[]>;
        rpcThreads?: pulumi.Input<number>;
        skipPreflightHealthCheck?: pulumi.Input<boolean>;
        skipSeedPhraseValidation?: pulumi.Input<boolean>;
        skipStartupLedgerVerification?: pulumi.Input<boolean>;
        snapshotArchiveFormat?: pulumi.Input<string>;
        snapshotIntervalSlots?: pulumi.Input<number>;
        snapshotPackagerNicenessAdjustment?: pulumi.Input<number>;
        snapshotVersion?: pulumi.Input<string>;
        stakedNodesOverrides?: pulumi.Input<string>;
        towerStorage?: pulumi.Input<string>;
        tpuCoalesceMs?: pulumi.Input<number>;
        tpuConnectionPoolSize?: pulumi.Input<number>;
        tpuDisableQuic?: pulumi.Input<boolean>;
        tpuEnableUdp?: pulumi.Input<boolean>;
        tvuReceiveThreads?: pulumi.Input<number>;
        unifiedSchedulerHandlerThreads?: pulumi.Input<number>;
        useSnapshotArchivesAtStartup?: pulumi.Input<string>;
        waitForSupermajority?: pulumi.Input<number>;
        walRecoveryMode: pulumi.Input<string>;
    }

    export interface KeyPairsArgs {
        identity: pulumi.Input<string>;
        voteAccount: pulumi.Input<string>;
    }

    export interface MetricsArgs {
        database: pulumi.Input<string>;
        password: pulumi.Input<string>;
        url: pulumi.Input<string>;
        user: pulumi.Input<string>;
    }

    export interface ShutdownPolicyArgs {
        force?: pulumi.Input<boolean>;
        maxDelinquentStake?: pulumi.Input<number>;
        minIdleTime?: pulumi.Input<number>;
        skipHealthCheck?: pulumi.Input<boolean>;
        skipNewSnapshotCheck?: pulumi.Input<boolean>;
    }

    export interface StartupPolicyArgs {
        waitForRPCHealth?: pulumi.Input<boolean>;
    }

    export interface TimeoutConfigArgs {
        rpcServiceTimeout?: pulumi.Input<number>;
    }
}

export namespace apt {
    export interface ConfigArgs {
        excludeDefaultSources?: pulumi.Input<boolean>;
        sources?: pulumi.Input<pulumi.Input<inputs.apt.SourceArgs>[]>;
    }

    export interface SignedByArgs {
        paths?: pulumi.Input<pulumi.Input<string>[]>;
        publicKey?: pulumi.Input<string>;
    }

    export interface SourceArgs {
        URIs: pulumi.Input<pulumi.Input<string>[]>;
        allowDowngradeToInsecure?: pulumi.Input<boolean>;
        allowInsecure?: pulumi.Input<boolean>;
        allowWeak?: pulumi.Input<boolean>;
        architectures?: pulumi.Input<pulumi.Input<string>[]>;
        checkDate?: pulumi.Input<boolean>;
        checkValidUntil?: pulumi.Input<boolean>;
        components: pulumi.Input<pulumi.Input<string>[]>;
        dateMaxFuture?: pulumi.Input<number>;
        extraLines?: pulumi.Input<pulumi.Input<string>[]>;
        inReleasePath?: pulumi.Input<string>;
        signedBy?: pulumi.Input<inputs.apt.SignedByArgs>;
        snapshot?: pulumi.Input<string>;
        suites: pulumi.Input<pulumi.Input<string>[]>;
        trusted?: pulumi.Input<boolean>;
        types: pulumi.Input<pulumi.Input<string>[]>;
        validUntilMax?: pulumi.Input<number>;
        validUntilMin?: pulumi.Input<number>;
    }
}

export namespace deb {
    export interface PackageArgs {
        name: pulumi.Input<string>;
        path?: pulumi.Input<string>;
        targetRelease?: pulumi.Input<string>;
        version?: pulumi.Input<string>;
    }

    export interface PackageConfigArgs {
        additional?: pulumi.Input<pulumi.Input<string>[]>;
        override?: pulumi.Input<pulumi.Input<inputs.deb.PackageArgs>[]>;
    }
}

export namespace explorer {
    export interface ExplorerFlagsArgs {
        hostname?: pulumi.Input<string>;
        keepAliveTimeout?: pulumi.Input<number>;
        port?: pulumi.Input<number>;
    }
}

export namespace faucet {
    export interface FaucetFlagsArgs {
        allowIPs?: pulumi.Input<pulumi.Input<string>[]>;
        perRequestCap?: pulumi.Input<number>;
        perTimeCap?: pulumi.Input<number>;
        sliceSeconds?: pulumi.Input<number>;
    }
}

export namespace firedancer {
    export interface ConfigArgs {
        consensus?: pulumi.Input<inputs.firedancer.ConfigConsensusArgs>;
        dynamicPortRange?: pulumi.Input<string>;
        extraConfig?: pulumi.Input<pulumi.Input<string>[]>;
        gossip?: pulumi.Input<inputs.firedancer.ConfigGossipArgs>;
        hugetlbfs?: pulumi.Input<inputs.firedancer.ConfigHugeTLBFSArgs>;
        layout?: pulumi.Input<inputs.firedancer.ConfigLayoutArgs>;
        ledger?: pulumi.Input<inputs.firedancer.ConfigLedgerArgs>;
        log?: pulumi.Input<inputs.firedancer.ConfigLogArgs>;
        name?: pulumi.Input<string>;
        reporting?: pulumi.Input<inputs.firedancer.ConfigReportingArgs>;
        rpc?: pulumi.Input<inputs.firedancer.ConfigRPCArgs>;
        scratchDirectory?: pulumi.Input<string>;
        snapshots?: pulumi.Input<inputs.firedancer.ConfigSnapshotsArgs>;
        user?: pulumi.Input<string>;
    }

    export interface ConfigConsensusArgs {
        authorizedVoterPaths?: pulumi.Input<pulumi.Input<string>[]>;
        expectedBankHash?: pulumi.Input<string>;
        expectedGenesisHash?: pulumi.Input<string>;
        expectedShredVersion?: pulumi.Input<number>;
        genesisFetch?: pulumi.Input<boolean>;
        hardForkAtSlots?: pulumi.Input<pulumi.Input<string>[]>;
        identityPath?: pulumi.Input<string>;
        knownValidators?: pulumi.Input<pulumi.Input<string>[]>;
        osNetworkLimitsTest?: pulumi.Input<boolean>;
        pohSpeedTest?: pulumi.Input<boolean>;
        snapshotFetch?: pulumi.Input<boolean>;
        voteAccountPath?: pulumi.Input<string>;
        waitForSupermajorityAtSlot?: pulumi.Input<number>;
        waitForVoteToStartLeader?: pulumi.Input<boolean>;
    }

    export interface ConfigGossipArgs {
        entrypoints?: pulumi.Input<pulumi.Input<string>[]>;
        host?: pulumi.Input<string>;
        port?: pulumi.Input<number>;
        portCheck?: pulumi.Input<boolean>;
    }

    export interface ConfigHugeTLBFSArgs {
        mountPath?: pulumi.Input<string>;
    }

    export interface ConfigLayoutArgs {
        affinity?: pulumi.Input<string>;
        agaveAffinity?: pulumi.Input<string>;
        bankTileCount?: pulumi.Input<number>;
        netTileCount?: pulumi.Input<number>;
        quicTileCount?: pulumi.Input<number>;
        resolvTileCount?: pulumi.Input<number>;
        shredTileCount?: pulumi.Input<number>;
        verifyTileCount?: pulumi.Input<number>;
    }

    export interface ConfigLedgerArgs {
        accountIndexExcludeKeys?: pulumi.Input<pulumi.Input<string>[]>;
        accountIndexIncludeKeys?: pulumi.Input<pulumi.Input<string>[]>;
        accountIndexes?: pulumi.Input<pulumi.Input<string>[]>;
        accountsPath?: pulumi.Input<string>;
        limitSize?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
        requireTower?: pulumi.Input<boolean>;
        snapshotArchiveFormat?: pulumi.Input<string>;
    }

    export interface ConfigLogArgs {
        colorize?: pulumi.Input<string>;
        levelFlush?: pulumi.Input<string>;
        levelLogfile?: pulumi.Input<string>;
        levelStderr?: pulumi.Input<string>;
        path?: pulumi.Input<string>;
    }

    export interface ConfigRPCArgs {
        bigtableLedgerStorage?: pulumi.Input<boolean>;
        extendedTxMetadataStorage?: pulumi.Input<boolean>;
        fullApi?: pulumi.Input<boolean>;
        onlyKnown?: pulumi.Input<boolean>;
        port?: pulumi.Input<number>;
        private?: pulumi.Input<boolean>;
        pubsubEnableBlockSubscription?: pulumi.Input<boolean>;
        pubsubEnableVoteSubscription?: pulumi.Input<boolean>;
        transactionHistory?: pulumi.Input<boolean>;
    }

    export interface ConfigReportingArgs {
        solanaMetricsConfig?: pulumi.Input<string>;
    }

    export interface ConfigSnapshotsArgs {
        fullSnapshotIntervalSlots?: pulumi.Input<number>;
        incrementalPath?: pulumi.Input<string>;
        incrementalSnapshotIntervalSlots?: pulumi.Input<number>;
        incrementalSnapshots?: pulumi.Input<boolean>;
        maximumFullSnapshotsToRetain?: pulumi.Input<number>;
        maximumIncrementalSnapshotsToRetain?: pulumi.Input<number>;
        minimumSnapshotDownloadSpeed?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
    }

    export interface KeyPairsArgs {
        identity: pulumi.Input<string>;
        voteAccount: pulumi.Input<string>;
    }
}

export namespace firewall {
    export interface FirewallParamsArgs {
        allowPorts: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace genesis {
    export interface BootstrapAccountArgs {
        balanceLamports?: pulumi.Input<number>;
        identityPubkey: pulumi.Input<string>;
        stakeLamports?: pulumi.Input<number>;
        stakePubkey: pulumi.Input<string>;
        votePubkey: pulumi.Input<string>;
    }

    export interface BootstrapValidatorArgs {
        identityPubkey: pulumi.Input<string>;
        stakePubkey: pulumi.Input<string>;
        votePubkey: pulumi.Input<string>;
    }

    export interface GenesisFlagsArgs {
        bootstrapStakeAuthorizedPubkey?: pulumi.Input<string>;
        bootstrapValidatorLamports?: pulumi.Input<number>;
        bootstrapValidatorStakeLamports?: pulumi.Input<number>;
        bootstrapValidators: pulumi.Input<pulumi.Input<inputs.genesis.BootstrapValidatorArgs>[]>;
        clusterType?: pulumi.Input<string>;
        creationTime?: pulumi.Input<string>;
        deactivateFeatures?: pulumi.Input<pulumi.Input<string>[]>;
        enableWarmupEpochs?: pulumi.Input<boolean>;
        extraFlags?: pulumi.Input<pulumi.Input<string>[]>;
        faucetLamports?: pulumi.Input<number>;
        faucetPubkey?: pulumi.Input<string>;
        feeBurnPercentage?: pulumi.Input<number>;
        hashesPerTick?: pulumi.Input<string>;
        inflation?: pulumi.Input<string>;
        lamportsPerByteYear?: pulumi.Input<number>;
        ledgerPath: pulumi.Input<string>;
        maxGenesisArchiveUnpackedSize?: pulumi.Input<number>;
        rentBurnPercentage?: pulumi.Input<number>;
        rentExemptionThreshold?: pulumi.Input<number>;
        slotsPerEpoch?: pulumi.Input<number>;
        targetLamportsPerSignature?: pulumi.Input<number>;
        targetSignaturesPerSlot?: pulumi.Input<number>;
        targetTickDuration?: pulumi.Input<number>;
        ticksPerSlot?: pulumi.Input<number>;
        url?: pulumi.Input<string>;
        voteCommissionPercentage?: pulumi.Input<number>;
    }

    export interface PrimordialAccountArgs {
        data?: pulumi.Input<string>;
        executable?: pulumi.Input<boolean>;
        lamports: pulumi.Input<number>;
        owner?: pulumi.Input<string>;
        pubkey: pulumi.Input<string>;
    }
}

export namespace geyser {
    export interface ConfigArgs {
        debugClientsHttp?: pulumi.Input<boolean>;
        grpc: pulumi.Input<inputs.geyser.GrpcConfigGrpcArgs>;
        log?: pulumi.Input<inputs.geyser.GrpcConfigLogArgs>;
        prometheus?: pulumi.Input<inputs.geyser.GrpcConfigPrometheusArgs>;
        tokio?: pulumi.Input<inputs.geyser.GrpcConfigTokioArgs>;
    }

    export interface GeyserPluginArgs {
        genericPluginConfig?: pulumi.Input<string>;
        yellowstoneGRPC?: pulumi.Input<inputs.geyser.YellowstoneGRPCArgs>;
    }

    export interface GrpcConfigFilterLimitsArgs {
        accounts?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsAccountsArgs>;
        blocks?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsBlocksArgs>;
        blocksMeta?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsBlocksMetaArgs>;
        entries?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsEntriesArgs>;
        slots?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsSlotsArgs>;
        transactions?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsTransactionsArgs>;
        transactionsStatus?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsTransactionsArgs>;
    }

    export interface GrpcConfigFilterLimitsAccountsArgs {
        accountMax?: pulumi.Input<number>;
        accountReject?: pulumi.Input<pulumi.Input<string>[]>;
        any?: pulumi.Input<boolean>;
        dataSliceMax?: pulumi.Input<number>;
        max?: pulumi.Input<number>;
        ownerMax?: pulumi.Input<number>;
        ownerReject?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface GrpcConfigFilterLimitsBlocksArgs {
        accountIncludeAny?: pulumi.Input<boolean>;
        accountIncludeMax?: pulumi.Input<number>;
        accountIncludeReject?: pulumi.Input<pulumi.Input<string>[]>;
        includeAccounts?: pulumi.Input<boolean>;
        includeEntries?: pulumi.Input<boolean>;
        includeTransactions?: pulumi.Input<boolean>;
        max?: pulumi.Input<number>;
    }

    export interface GrpcConfigFilterLimitsBlocksMetaArgs {
        max?: pulumi.Input<number>;
    }

    export interface GrpcConfigFilterLimitsEntriesArgs {
        max?: pulumi.Input<number>;
    }

    export interface GrpcConfigFilterLimitsSlotsArgs {
        max?: pulumi.Input<number>;
    }

    export interface GrpcConfigFilterLimitsTransactionsArgs {
        accountExcludeMax?: pulumi.Input<number>;
        accountIncludeMax?: pulumi.Input<number>;
        accountIncludeReject?: pulumi.Input<pulumi.Input<string>[]>;
        accountRequiredMax?: pulumi.Input<number>;
        any?: pulumi.Input<boolean>;
        max?: pulumi.Input<number>;
    }

    export interface GrpcConfigGrpcArgs {
        address: pulumi.Input<string>;
        channelCapacity?: pulumi.Input<number>;
        compression?: pulumi.Input<inputs.geyser.GrpcConfigGrpcCompressionArgs>;
        filterLimits?: pulumi.Input<inputs.geyser.GrpcConfigFilterLimitsArgs>;
        filterNameSizeLimit?: pulumi.Input<number>;
        filterNamesCleanupInterval?: pulumi.Input<string>;
        filterNamesSizeLimit?: pulumi.Input<number>;
        maxDecodingMessageSize?: pulumi.Input<number>;
        replayStoredSlots?: pulumi.Input<number>;
        serverHttp2AdaptiveWindow?: pulumi.Input<boolean>;
        serverHttp2KeepaliveInterval?: pulumi.Input<string>;
        serverHttp2KeepaliveTimeout?: pulumi.Input<string>;
        serverInitialConnectionWindowSize?: pulumi.Input<number>;
        serverInitialStreamWindowSize?: pulumi.Input<number>;
        snapshotClientChannelCapacity?: pulumi.Input<number>;
        snapshotPluginChannelCapacity?: pulumi.Input<number>;
        tlsConfig?: pulumi.Input<inputs.geyser.GrpcConfigGrpcServerTLSArgs>;
        unaryConcurrencyLimit?: pulumi.Input<number>;
        unaryDisabled?: pulumi.Input<boolean>;
        xToken?: pulumi.Input<string>;
    }

    export interface GrpcConfigGrpcCompressionArgs {
        accept?: pulumi.Input<pulumi.Input<string>[]>;
        send?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface GrpcConfigGrpcServerTLSArgs {
        certPath: pulumi.Input<string>;
        keyPath: pulumi.Input<string>;
    }

    export interface GrpcConfigLogArgs {
        level?: pulumi.Input<string>;
    }

    export interface GrpcConfigPrometheusArgs {
        address: pulumi.Input<string>;
    }

    export interface GrpcConfigTokioArgs {
        affinity?: pulumi.Input<pulumi.Input<number>[]>;
        workerThreads?: pulumi.Input<number>;
    }

    export interface YellowstoneGRPCArgs {
        config?: pulumi.Input<inputs.geyser.ConfigArgs>;
        json?: pulumi.Input<string>;
        version: pulumi.Input<string>;
    }
}

export namespace runner {
    export interface ConfigArgs {
        aptLockTimeout?: pulumi.Input<number>;
        keepPayload?: pulumi.Input<boolean>;
        packageConfig?: pulumi.Input<inputs.deb.PackageConfigArgs>;
    }
}

export namespace solana {
    export interface EnvironmentArgs {
        rpcURL: pulumi.Input<string>;
    }

    export interface StakeAccountKeyPairsArgs {
        stakeAccount: pulumi.Input<string>;
        voteAccount: pulumi.Input<string>;
    }

    export interface TxnOptionsArgs {
        blockHash?: pulumi.Input<string>;
        commitment?: pulumi.Input<string>;
        feePayer?: pulumi.Input<string>;
        from?: pulumi.Input<string>;
        keyPair?: pulumi.Input<string>;
        nonce?: pulumi.Input<string>;
        nonceAuthority?: pulumi.Input<string>;
        signer?: pulumi.Input<pulumi.Input<string>[]>;
        url?: pulumi.Input<string>;
        withComputeUnitPrice?: pulumi.Input<number>;
        withMemo?: pulumi.Input<string>;
        ws?: pulumi.Input<string>;
    }

    export interface ValidatorInfoArgs {
        details?: pulumi.Input<string>;
        iconURL?: pulumi.Input<string>;
        name: pulumi.Input<string>;
        website?: pulumi.Input<string>;
    }

    export interface VoteAccountKeyPairsArgs {
        authWithdrawer: pulumi.Input<string>;
        identity: pulumi.Input<string>;
        voteAccount: pulumi.Input<string>;
    }
}

export namespace ssh {
    /**
     * Instructions for how to connect to a remote endpoint.
     */
    export interface ConnectionArgs {
        /**
         * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
         */
        agentSocketPath?: pulumi.Input<string>;
        /**
         * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
         */
        dialErrorLimit?: pulumi.Input<number>;
        /**
         * The address of the resource to connect to.
         */
        host: pulumi.Input<string>;
        /**
         * The password we should use for the connection.
         */
        password?: pulumi.Input<string>;
        /**
         * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
         */
        perDialTimeout?: pulumi.Input<number>;
        /**
         * The port to connect to. Defaults to 22.
         */
        port?: pulumi.Input<number>;
        /**
         * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
         */
        privateKey?: pulumi.Input<string>;
        /**
         * The password to use in case the private key is encrypted.
         */
        privateKeyPassword?: pulumi.Input<string>;
        /**
         * The connection settings for the bastion/proxy host.
         */
        proxy?: pulumi.Input<inputs.ssh.ProxyConnectionArgs>;
        /**
         * The user that we should use for the connection.
         */
        user?: pulumi.Input<string>;
    }
    /**
     * connectionArgsProvideDefaults sets the appropriate defaults for ConnectionArgs
     */
    export function connectionArgsProvideDefaults(val: ConnectionArgs): ConnectionArgs {
        return {
            ...val,
            dialErrorLimit: (val.dialErrorLimit) ?? 10,
            perDialTimeout: (val.perDialTimeout) ?? 15,
            port: (val.port) ?? 22,
            proxy: (val.proxy ? pulumi.output(val.proxy).apply(inputs.ssh.proxyConnectionArgsProvideDefaults) : undefined),
            user: (val.user) ?? "root",
        };
    }

    /**
     * Instructions for how to connect to a remote endpoint via a bastion host.
     */
    export interface ProxyConnectionArgs {
        /**
         * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
         */
        agentSocketPath?: pulumi.Input<string>;
        /**
         * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
         */
        dialErrorLimit?: pulumi.Input<number>;
        /**
         * The address of the bastion host to connect to.
         */
        host: pulumi.Input<string>;
        /**
         * The password we should use for the connection to the bastion host.
         */
        password?: pulumi.Input<string>;
        /**
         * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
         */
        perDialTimeout?: pulumi.Input<number>;
        /**
         * The port of the bastion host to connect to.
         */
        port?: pulumi.Input<number>;
        /**
         * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
         */
        privateKey?: pulumi.Input<string>;
        /**
         * The password to use in case the private key is encrypted.
         */
        privateKeyPassword?: pulumi.Input<string>;
        /**
         * The user that we should use for the connection to the bastion host.
         */
        user?: pulumi.Input<string>;
    }
    /**
     * proxyConnectionArgsProvideDefaults sets the appropriate defaults for ProxyConnectionArgs
     */
    export function proxyConnectionArgsProvideDefaults(val: ProxyConnectionArgs): ProxyConnectionArgs {
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
    export interface TunerFsParamsArgs {
        fsNrOpen?: pulumi.Input<number>;
    }

    export interface TunerKernelParamsArgs {
        kernelHungTaskTimeoutSecs?: pulumi.Input<number>;
        kernelNmiWatchdog?: pulumi.Input<number>;
        kernelPidMax?: pulumi.Input<number>;
        kernelSchedMinGranularityNs?: pulumi.Input<number>;
        kernelSchedWakeupGranularityNs?: pulumi.Input<number>;
        kernelTimerMigration?: pulumi.Input<number>;
    }

    export interface TunerNetParamsArgs {
        netCoreRmemDefault?: pulumi.Input<number>;
        netCoreRmemMax?: pulumi.Input<number>;
        netCoreWmemDefault?: pulumi.Input<number>;
        netCoreWmemMax?: pulumi.Input<number>;
        netIpv4TcpCongestionControl?: pulumi.Input<string>;
        netIpv4TcpFastopen?: pulumi.Input<number>;
        netIpv4TcpLowLatency?: pulumi.Input<number>;
        netIpv4TcpModerateRcvbuf?: pulumi.Input<number>;
        netIpv4TcpNoMetricsSave?: pulumi.Input<number>;
        netIpv4TcpRmem?: pulumi.Input<string>;
        netIpv4TcpSack?: pulumi.Input<number>;
        netIpv4TcpTimestamps?: pulumi.Input<number>;
        netIpv4TcpTwReuse?: pulumi.Input<number>;
        netIpv4TcpWmem?: pulumi.Input<string>;
    }

    export interface TunerParamsArgs {
        cpuGovernor?: pulumi.Input<enums.tuner.CpuGovernor>;
        fs?: pulumi.Input<inputs.tuner.TunerFsParamsArgs>;
        kernel?: pulumi.Input<inputs.tuner.TunerKernelParamsArgs>;
        net?: pulumi.Input<inputs.tuner.TunerNetParamsArgs>;
        vm?: pulumi.Input<inputs.tuner.TunerVmParamsArgs>;
    }

    export interface TunerVmParamsArgs {
        vmDirtyBackgroundRatio?: pulumi.Input<number>;
        vmDirtyExpireCentisecs?: pulumi.Input<number>;
        vmDirtyRatio?: pulumi.Input<number>;
        vmDirtyWritebackCentisecs?: pulumi.Input<number>;
        vmDirtytimeExpireSeconds?: pulumi.Input<number>;
        vmMaxMapCount?: pulumi.Input<number>;
        vmMinFreeKbytes?: pulumi.Input<number>;
        vmStatInterval?: pulumi.Input<number>;
        vmSwappiness?: pulumi.Input<number>;
    }
}

export namespace watchtower {
    export interface DiscordConfigArgs {
        webhookUrl: pulumi.Input<string>;
    }

    export interface NotificationConfigArgs {
        discord?: pulumi.Input<inputs.watchtower.DiscordConfigArgs>;
        pagerDuty?: pulumi.Input<inputs.watchtower.PagerDutyConfigArgs>;
        slack?: pulumi.Input<inputs.watchtower.SlackConfigArgs>;
        telegram?: pulumi.Input<inputs.watchtower.TelegramConfigArgs>;
        twilio?: pulumi.Input<inputs.watchtower.TwilioConfigArgs>;
    }

    export interface PagerDutyConfigArgs {
        integrationKey: pulumi.Input<string>;
    }

    export interface SlackConfigArgs {
        webhookUrl: pulumi.Input<string>;
    }

    export interface TelegramConfigArgs {
        botToken: pulumi.Input<string>;
        chatId: pulumi.Input<string>;
    }

    export interface TwilioConfigArgs {
        accountSid: pulumi.Input<string>;
        authToken: pulumi.Input<string>;
        fromNumber: pulumi.Input<string>;
        toNumber: pulumi.Input<string>;
    }

    export interface WatchtowerFlagsArgs {
        activeStakeAlertThreshold?: pulumi.Input<number>;
        ignoreHttpBadGateway?: pulumi.Input<boolean>;
        interval?: pulumi.Input<number>;
        minimumValidatorIdentityBalance?: pulumi.Input<number>;
        monitorActiveStake?: pulumi.Input<boolean>;
        nameSuffix?: pulumi.Input<string>;
        rpcTimeout?: pulumi.Input<number>;
        unhealthyThreshold?: pulumi.Input<number>;
        validatorIdentity: pulumi.Input<pulumi.Input<string>[]>;
    }
}
