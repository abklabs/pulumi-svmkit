// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ABKLabs.Svmkit.Agave.Outputs
{

    [OutputType]
    public sealed class Flags
    {
        public readonly ImmutableArray<string> AccountIndex;
        public readonly ImmutableArray<string> AccountIndexExcludeKey;
        public readonly ImmutableArray<string> AccountIndexIncludeKey;
        public readonly ImmutableArray<string> AccountShrinkPath;
        public readonly int? AccountsDbCacheLimitMb;
        public readonly bool? AccountsDbTestHashCalculation;
        public readonly string? AccountsHashCachePath;
        public readonly int? AccountsIndexBins;
        public readonly ImmutableArray<string> AccountsIndexPath;
        public readonly int? AccountsIndexScanResultsLimitMb;
        public readonly bool? AccountsShrinkOptimizeTotalSpace;
        public readonly string? AccountsShrinkRatio;
        public readonly bool? AllowPrivateAddr;
        public readonly ImmutableArray<string> AuthorizedVoter;
        public readonly string? BindAddress;
        public readonly string? BlockProductionMethod;
        public readonly string? BlockVerificationMethod;
        public readonly string? CheckVoteAccount;
        public readonly int? ContactDebugInterval;
        public readonly bool? Cuda;
        public readonly ImmutableArray<string> DebugKey;
        public readonly int? DevHaltAtSlot;
        public readonly bool? DisableBankingTrace;
        public readonly string? DynamicPortRange;
        public readonly int? EnableBankingTrace;
        public readonly bool? EnableBigtableLedgerUpload;
        public readonly bool? EnableExtendedTxMetadataStorage;
        public readonly bool? EnableRpcBigtableLedgerStorage;
        public readonly bool? EnableRpcTransactionHistory;
        public readonly ImmutableArray<string> EntryPoint;
        public readonly string? EtcdCacertFile;
        public readonly string? EtcdCertFile;
        public readonly string? EtcdDomainName;
        public readonly ImmutableArray<string> EtcdEndpoint;
        public readonly string? EtcdKeyFile;
        public readonly string? ExpectedBankHash;
        public readonly string? ExpectedGenesisHash;
        public readonly int? ExpectedShredVersion;
        public readonly ImmutableArray<string> ExtraFlags;
        public readonly bool? FullRpcAPI;
        public readonly string? FullSnapshotArchivePath;
        public readonly int? FullSnapshotIntervalSlots;
        public readonly bool? GeyserPluginAlwaysEnabled;
        public readonly ImmutableArray<string> GeyserPluginConfig;
        public readonly string? GossipHost;
        public readonly int? GossipPort;
        public readonly ImmutableArray<string> GossipValidator;
        public readonly ImmutableArray<int> HardFork;
        public readonly int? HealthCheckSlotDistance;
        public readonly string? IncrementalSnapshotArchivePath;
        public readonly string? InitCompleteFile;
        public readonly ImmutableArray<string> KnownValidator;
        public readonly int? LimitLedgerSize;
        public readonly string? Log;
        public readonly int? LogMessagesBytesLimit;
        public readonly int? MaxGenesisArchiveUnpackedSize;
        public readonly int? MaximumFullSnapshotsToRetain;
        public readonly int? MaximumIncrementalSnapshotsToRetain;
        public readonly int? MaximumLocalSnapshotAge;
        public readonly int? MaximumSnapshotDownloadAbort;
        public readonly int? MinimalSnapshotDownloadSpeed;
        public readonly bool? NoGenesisFetch;
        public readonly bool? NoIncrementalSnapshots;
        public readonly bool? NoSnapshotFetch;
        public readonly bool? NoVoting;
        public readonly bool NoWaitForVoteToStartLeader;
        public readonly bool? OnlyKnownRPC;
        public readonly bool? PrivateRPC;
        public readonly string? PublicRpcAddress;
        public readonly string? PublicTpuAddress;
        public readonly string? PublicTpuForwardsAddress;
        public readonly ImmutableArray<string> RepairValidator;
        public readonly bool? RequireTower;
        public readonly bool? RestrictedRepairOnlyMode;
        public readonly int? RocksdbFifoShredStorageSize;
        public readonly string? RocksdbShredCompaction;
        public readonly string? RpcBigtableAppProfileId;
        public readonly string? RpcBigtableInstanceName;
        public readonly int? RpcBigtableMaxMessageSize;
        public readonly int? RpcBigtableTimeout;
        public readonly string RpcBindAddress;
        public readonly string? RpcFaucetAddress;
        public readonly int? RpcMaxMultipleAccounts;
        public readonly int? RpcMaxRequestBodySize;
        public readonly int? RpcNicenessAdjustment;
        public readonly int RpcPort;
        public readonly bool? RpcPubsubEnableBlockSubscription;
        public readonly bool? RpcPubsubEnableVoteSubscription;
        public readonly int? RpcPubsubMaxActiveSubscriptions;
        public readonly int? RpcPubsubNotificationThreads;
        public readonly int? RpcPubsubQueueCapacityBytes;
        public readonly int? RpcPubsubQueueCapacityItems;
        public readonly int? RpcPubsubWorkerThreads;
        public readonly bool? RpcScanAndFixRoots;
        public readonly int? RpcSendLeaderCount;
        public readonly int? RpcSendRetryMs;
        public readonly int? RpcSendServiceMaxRetries;
        public readonly bool? RpcSendTransactionAlsoLeader;
        public readonly int? RpcSendTransactionRetryPoolMaxSize;
        public readonly ImmutableArray<string> RpcSendTransactionTpuPeer;
        public readonly int? RpcThreads;
        public readonly bool? SkipPreflightHealthCheck;
        public readonly bool? SkipSeedPhraseValidation;
        public readonly bool? SkipStartupLedgerVerification;
        public readonly string? SnapshotArchiveFormat;
        public readonly int? SnapshotIntervalSlots;
        public readonly int? SnapshotPackagerNicenessAdjustment;
        public readonly string? SnapshotVersion;
        public readonly string? StakedNodesOverrides;
        public readonly string? TowerStorage;
        public readonly int? TpuCoalesceMs;
        public readonly int? TpuConnectionPoolSize;
        public readonly bool? TpuDisableQuic;
        public readonly bool? TpuEnableUdp;
        public readonly int? TvuReceiveThreads;
        public readonly int? UnifiedSchedulerHandlerThreads;
        public readonly string? UseSnapshotArchivesAtStartup;
        public readonly int? WaitForSupermajority;
        public readonly string WalRecoveryMode;

        [OutputConstructor]
        private Flags(
            ImmutableArray<string> accountIndex,

            ImmutableArray<string> accountIndexExcludeKey,

            ImmutableArray<string> accountIndexIncludeKey,

            ImmutableArray<string> accountShrinkPath,

            int? accountsDbCacheLimitMb,

            bool? accountsDbTestHashCalculation,

            string? accountsHashCachePath,

            int? accountsIndexBins,

            ImmutableArray<string> accountsIndexPath,

            int? accountsIndexScanResultsLimitMb,

            bool? accountsShrinkOptimizeTotalSpace,

            string? accountsShrinkRatio,

            bool? allowPrivateAddr,

            ImmutableArray<string> authorizedVoter,

            string? bindAddress,

            string? blockProductionMethod,

            string? blockVerificationMethod,

            string? checkVoteAccount,

            int? contactDebugInterval,

            bool? cuda,

            ImmutableArray<string> debugKey,

            int? devHaltAtSlot,

            bool? disableBankingTrace,

            string? dynamicPortRange,

            int? enableBankingTrace,

            bool? enableBigtableLedgerUpload,

            bool? enableExtendedTxMetadataStorage,

            bool? enableRpcBigtableLedgerStorage,

            bool? enableRpcTransactionHistory,

            ImmutableArray<string> entryPoint,

            string? etcdCacertFile,

            string? etcdCertFile,

            string? etcdDomainName,

            ImmutableArray<string> etcdEndpoint,

            string? etcdKeyFile,

            string? expectedBankHash,

            string? expectedGenesisHash,

            int? expectedShredVersion,

            ImmutableArray<string> extraFlags,

            bool? fullRpcAPI,

            string? fullSnapshotArchivePath,

            int? fullSnapshotIntervalSlots,

            bool? geyserPluginAlwaysEnabled,

            ImmutableArray<string> geyserPluginConfig,

            string? gossipHost,

            int? gossipPort,

            ImmutableArray<string> gossipValidator,

            ImmutableArray<int> hardFork,

            int? healthCheckSlotDistance,

            string? incrementalSnapshotArchivePath,

            string? initCompleteFile,

            ImmutableArray<string> knownValidator,

            int? limitLedgerSize,

            string? log,

            int? logMessagesBytesLimit,

            int? maxGenesisArchiveUnpackedSize,

            int? maximumFullSnapshotsToRetain,

            int? maximumIncrementalSnapshotsToRetain,

            int? maximumLocalSnapshotAge,

            int? maximumSnapshotDownloadAbort,

            int? minimalSnapshotDownloadSpeed,

            bool? noGenesisFetch,

            bool? noIncrementalSnapshots,

            bool? noSnapshotFetch,

            bool? noVoting,

            bool noWaitForVoteToStartLeader,

            bool? onlyKnownRPC,

            bool? privateRPC,

            string? publicRpcAddress,

            string? publicTpuAddress,

            string? publicTpuForwardsAddress,

            ImmutableArray<string> repairValidator,

            bool? requireTower,

            bool? restrictedRepairOnlyMode,

            int? rocksdbFifoShredStorageSize,

            string? rocksdbShredCompaction,

            string? rpcBigtableAppProfileId,

            string? rpcBigtableInstanceName,

            int? rpcBigtableMaxMessageSize,

            int? rpcBigtableTimeout,

            string rpcBindAddress,

            string? rpcFaucetAddress,

            int? rpcMaxMultipleAccounts,

            int? rpcMaxRequestBodySize,

            int? rpcNicenessAdjustment,

            int rpcPort,

            bool? rpcPubsubEnableBlockSubscription,

            bool? rpcPubsubEnableVoteSubscription,

            int? rpcPubsubMaxActiveSubscriptions,

            int? rpcPubsubNotificationThreads,

            int? rpcPubsubQueueCapacityBytes,

            int? rpcPubsubQueueCapacityItems,

            int? rpcPubsubWorkerThreads,

            bool? rpcScanAndFixRoots,

            int? rpcSendLeaderCount,

            int? rpcSendRetryMs,

            int? rpcSendServiceMaxRetries,

            bool? rpcSendTransactionAlsoLeader,

            int? rpcSendTransactionRetryPoolMaxSize,

            ImmutableArray<string> rpcSendTransactionTpuPeer,

            int? rpcThreads,

            bool? skipPreflightHealthCheck,

            bool? skipSeedPhraseValidation,

            bool? skipStartupLedgerVerification,

            string? snapshotArchiveFormat,

            int? snapshotIntervalSlots,

            int? snapshotPackagerNicenessAdjustment,

            string? snapshotVersion,

            string? stakedNodesOverrides,

            string? towerStorage,

            int? tpuCoalesceMs,

            int? tpuConnectionPoolSize,

            bool? tpuDisableQuic,

            bool? tpuEnableUdp,

            int? tvuReceiveThreads,

            int? unifiedSchedulerHandlerThreads,

            string? useSnapshotArchivesAtStartup,

            int? waitForSupermajority,

            string walRecoveryMode)
        {
            AccountIndex = accountIndex;
            AccountIndexExcludeKey = accountIndexExcludeKey;
            AccountIndexIncludeKey = accountIndexIncludeKey;
            AccountShrinkPath = accountShrinkPath;
            AccountsDbCacheLimitMb = accountsDbCacheLimitMb;
            AccountsDbTestHashCalculation = accountsDbTestHashCalculation;
            AccountsHashCachePath = accountsHashCachePath;
            AccountsIndexBins = accountsIndexBins;
            AccountsIndexPath = accountsIndexPath;
            AccountsIndexScanResultsLimitMb = accountsIndexScanResultsLimitMb;
            AccountsShrinkOptimizeTotalSpace = accountsShrinkOptimizeTotalSpace;
            AccountsShrinkRatio = accountsShrinkRatio;
            AllowPrivateAddr = allowPrivateAddr;
            AuthorizedVoter = authorizedVoter;
            BindAddress = bindAddress;
            BlockProductionMethod = blockProductionMethod;
            BlockVerificationMethod = blockVerificationMethod;
            CheckVoteAccount = checkVoteAccount;
            ContactDebugInterval = contactDebugInterval;
            Cuda = cuda;
            DebugKey = debugKey;
            DevHaltAtSlot = devHaltAtSlot;
            DisableBankingTrace = disableBankingTrace;
            DynamicPortRange = dynamicPortRange;
            EnableBankingTrace = enableBankingTrace;
            EnableBigtableLedgerUpload = enableBigtableLedgerUpload;
            EnableExtendedTxMetadataStorage = enableExtendedTxMetadataStorage;
            EnableRpcBigtableLedgerStorage = enableRpcBigtableLedgerStorage;
            EnableRpcTransactionHistory = enableRpcTransactionHistory;
            EntryPoint = entryPoint;
            EtcdCacertFile = etcdCacertFile;
            EtcdCertFile = etcdCertFile;
            EtcdDomainName = etcdDomainName;
            EtcdEndpoint = etcdEndpoint;
            EtcdKeyFile = etcdKeyFile;
            ExpectedBankHash = expectedBankHash;
            ExpectedGenesisHash = expectedGenesisHash;
            ExpectedShredVersion = expectedShredVersion;
            ExtraFlags = extraFlags;
            FullRpcAPI = fullRpcAPI;
            FullSnapshotArchivePath = fullSnapshotArchivePath;
            FullSnapshotIntervalSlots = fullSnapshotIntervalSlots;
            GeyserPluginAlwaysEnabled = geyserPluginAlwaysEnabled;
            GeyserPluginConfig = geyserPluginConfig;
            GossipHost = gossipHost;
            GossipPort = gossipPort;
            GossipValidator = gossipValidator;
            HardFork = hardFork;
            HealthCheckSlotDistance = healthCheckSlotDistance;
            IncrementalSnapshotArchivePath = incrementalSnapshotArchivePath;
            InitCompleteFile = initCompleteFile;
            KnownValidator = knownValidator;
            LimitLedgerSize = limitLedgerSize;
            Log = log;
            LogMessagesBytesLimit = logMessagesBytesLimit;
            MaxGenesisArchiveUnpackedSize = maxGenesisArchiveUnpackedSize;
            MaximumFullSnapshotsToRetain = maximumFullSnapshotsToRetain;
            MaximumIncrementalSnapshotsToRetain = maximumIncrementalSnapshotsToRetain;
            MaximumLocalSnapshotAge = maximumLocalSnapshotAge;
            MaximumSnapshotDownloadAbort = maximumSnapshotDownloadAbort;
            MinimalSnapshotDownloadSpeed = minimalSnapshotDownloadSpeed;
            NoGenesisFetch = noGenesisFetch;
            NoIncrementalSnapshots = noIncrementalSnapshots;
            NoSnapshotFetch = noSnapshotFetch;
            NoVoting = noVoting;
            NoWaitForVoteToStartLeader = noWaitForVoteToStartLeader;
            OnlyKnownRPC = onlyKnownRPC;
            PrivateRPC = privateRPC;
            PublicRpcAddress = publicRpcAddress;
            PublicTpuAddress = publicTpuAddress;
            PublicTpuForwardsAddress = publicTpuForwardsAddress;
            RepairValidator = repairValidator;
            RequireTower = requireTower;
            RestrictedRepairOnlyMode = restrictedRepairOnlyMode;
            RocksdbFifoShredStorageSize = rocksdbFifoShredStorageSize;
            RocksdbShredCompaction = rocksdbShredCompaction;
            RpcBigtableAppProfileId = rpcBigtableAppProfileId;
            RpcBigtableInstanceName = rpcBigtableInstanceName;
            RpcBigtableMaxMessageSize = rpcBigtableMaxMessageSize;
            RpcBigtableTimeout = rpcBigtableTimeout;
            RpcBindAddress = rpcBindAddress;
            RpcFaucetAddress = rpcFaucetAddress;
            RpcMaxMultipleAccounts = rpcMaxMultipleAccounts;
            RpcMaxRequestBodySize = rpcMaxRequestBodySize;
            RpcNicenessAdjustment = rpcNicenessAdjustment;
            RpcPort = rpcPort;
            RpcPubsubEnableBlockSubscription = rpcPubsubEnableBlockSubscription;
            RpcPubsubEnableVoteSubscription = rpcPubsubEnableVoteSubscription;
            RpcPubsubMaxActiveSubscriptions = rpcPubsubMaxActiveSubscriptions;
            RpcPubsubNotificationThreads = rpcPubsubNotificationThreads;
            RpcPubsubQueueCapacityBytes = rpcPubsubQueueCapacityBytes;
            RpcPubsubQueueCapacityItems = rpcPubsubQueueCapacityItems;
            RpcPubsubWorkerThreads = rpcPubsubWorkerThreads;
            RpcScanAndFixRoots = rpcScanAndFixRoots;
            RpcSendLeaderCount = rpcSendLeaderCount;
            RpcSendRetryMs = rpcSendRetryMs;
            RpcSendServiceMaxRetries = rpcSendServiceMaxRetries;
            RpcSendTransactionAlsoLeader = rpcSendTransactionAlsoLeader;
            RpcSendTransactionRetryPoolMaxSize = rpcSendTransactionRetryPoolMaxSize;
            RpcSendTransactionTpuPeer = rpcSendTransactionTpuPeer;
            RpcThreads = rpcThreads;
            SkipPreflightHealthCheck = skipPreflightHealthCheck;
            SkipSeedPhraseValidation = skipSeedPhraseValidation;
            SkipStartupLedgerVerification = skipStartupLedgerVerification;
            SnapshotArchiveFormat = snapshotArchiveFormat;
            SnapshotIntervalSlots = snapshotIntervalSlots;
            SnapshotPackagerNicenessAdjustment = snapshotPackagerNicenessAdjustment;
            SnapshotVersion = snapshotVersion;
            StakedNodesOverrides = stakedNodesOverrides;
            TowerStorage = towerStorage;
            TpuCoalesceMs = tpuCoalesceMs;
            TpuConnectionPoolSize = tpuConnectionPoolSize;
            TpuDisableQuic = tpuDisableQuic;
            TpuEnableUdp = tpuEnableUdp;
            TvuReceiveThreads = tvuReceiveThreads;
            UnifiedSchedulerHandlerThreads = unifiedSchedulerHandlerThreads;
            UseSnapshotArchivesAtStartup = useSnapshotArchivesAtStartup;
            WaitForSupermajority = waitForSupermajority;
            WalRecoveryMode = walRecoveryMode;
        }
    }
}
