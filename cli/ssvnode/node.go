package ssvnode

import (
	"encoding/hex"
	"fmt"
	"github.com/bloxapp/ssv/cli/config"
	"github.com/bloxapp/ssv/storage"
	"github.com/bloxapp/ssv/storage/basedb"
	"github.com/bloxapp/ssv/utils/logex"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"

	"github.com/bloxapp/ssv/cli/flags"
)

type ssvNodeConfig struct {
	config.GlobalConfig `yaml:"global"`
	DBOptions basedb.Options `yaml:"db"`
}

var cfg ssvNodeConfig

var globalArgs config.Args


// startNodeCmd is the command to start SSV node
var StartNodeCmd = &cobra.Command{
	Use:   "start-node",
	Short: "Starts an instance of SSV node",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cleanenv.ReadConfig(globalArgs.ConfigPath, &cfg); err != nil {
			log.Fatal(err)
		}

		loggerLevel, err := logex.GetLoggerLevelValue(cfg.LogLevel)
		Logger := logex.Build(cmd.Parent().Short, loggerLevel)

		if err != nil {
			Logger.Warn(fmt.Sprintf("Default log level set to %s", loggerLevel),zap.Error(err))
		}

		cfg.DBOptions.Logger = Logger
		db, err := storage.GetStorageFactory(cfg.DBOptions)


		//
		//nodeID, err := flags.GetNodeIDKeyFlagValue(cmd)
		//if err != nil {
		//	Logger.Fatal("failed to get node ID flag value", zap.Error(err))
		//}
		//logger := Logger.With(zap.Uint64("node_id", nodeID))
		//
		//eth2Network, err := flags.GetNetworkFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get eth2Network flag value", zap.Error(err))
		//}
		//
		//discoveryType, err := flags.GetDiscoveryFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get val flag value", zap.Error(err))
		//}
		//
		//consensusType, err := flags.GetConsensusFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get val flag value", zap.Error(err))
		//}
		//
		//beaconAddr, err := flags.GetBeaconAddrFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get beacon node address flag value", zap.Error(err))
		//}
		//
		//privKey, err := flags.GetPrivKeyFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get private key flag value", zap.Error(err))
		//}
		//
		//validatorKey, err := flags.GetValidatorKeyFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get validator public key flag value", zap.Error(err))
		//}
		//
		//sigCollectionTimeout, err := flags.GetSignatureCollectionTimeValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get signature timeout key flag value", zap.Error(err))
		//}
		//
		//hostDNS, err := flags.GetHostDNSFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get hostDNS key flag value", zap.Error(err))
		//}
		//
		//hostAddress, err := flags.GetHostAddressFlagValue(cmd)
		//if err != nil {
		//	logger.Fatal("failed to get hostAddress key flag value", zap.Error(err))
		//}
		//
		//tcpPort, err := flags.GetTCPPortFlagValue(cmd)
		//if err != nil {
		//	Logger.Fatal("failed to get tcp port flag value", zap.Error(err))
		//}
		//udpPort, err := flags.GetUDPPortFlagValue(cmd)
		//if err != nil {
		//	Logger.Fatal("failed to get udp port flag value", zap.Error(err))
		//}
		//
		//genesisEpoch, err := flags.GetGenesisEpochValue(cmd)
		//if err != nil {
		//	Logger.Fatal("failed to get genesis epoch flag value", zap.Error(err))
		//}
		//
		//storagePath, err := flags.GetStoragePathValue(cmd)
		//if err != nil {
		//	Logger.Fatal("failed to get storage path flag value", zap.Error(err))
		//}
		//
		//validatorPk := &bls.PublicKey{}
		//if err := validatorPk.DeserializeHexStr(validatorKey); err != nil {
		//	logger.Fatal("failed to decode validator key", zap.Error(err))
		//}
		//
		//shareKey := &bls.SecretKey{}
		//if err := shareKey.SetHexString(privKey); err != nil {
		//	logger.Fatal("failed to set hex private key", zap.Error(err))
		//}
		//
		//// init storage
		//validatorStorage, ibftStorage := configureStorage(storagePath, logger, validatorPk, shareKey, nodeID)
		//
		//beaconClient, err := prysmgrpc.New(cmd.Context(), logger, eth2Network, []byte("BloxStaking"), beaconAddr)
		//if err != nil {
		//	logger.Fatal("failed to create beacon client", zap.Error(err))
		//}
		//
		//Logger.Info("Running node with ports", zap.Int("tcp", tcpPort), zap.Int("udp", udpPort))
		//Logger.Info("Running node with genesis epoch", zap.Uint64("epoch", genesisEpoch))
		//logger.Debug("Node params",
		//	zap.String("eth2Network", string(eth2Network)),
		//	zap.String("discovery-type", discoveryType),
		//	zap.String("val", consensusType),
		//	zap.String("beacon-addr", beaconAddr),
		//	zap.String("validator", "0x"+validatorKey[:12]+"..."))
		//
		//cfg := p2p.Config{
		//	DiscoveryType: discoveryType,
		//	BootstrapNodeAddr: []string{
		//		// deployemnt
		//		// internal ip
		//		//"enr:-LK4QDAmZK-69qRU5q-cxW6BqLwIlWoYH-BoRlX2N7D9rXBlM7OJ9tWRRtryqvCW04geHC_ab8QmWT9QULnT0Tc5S1cBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhArqAsGJc2VjcDI1NmsxoQO8KQz5L1UEXzEr-CXFFq1th0eG6gopbdul2OQVMuxfMoN0Y3CCE4iDdWRwgg-g",
		//		//external ip
		//		"enr:-LK4QHVq6HEA2KVnAw593SRMqUOvMGlkP8Jb-qHn4yPLHx--cStvWc38Or2xLcWgDPynVxXPT9NWIEXRzrBUsLmcFkUBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhDbUHcyJc2VjcDI1NmsxoQO8KQz5L1UEXzEr-CXFFq1th0eG6gopbdul2OQVMuxfMoN0Y3CCE4iDdWRwgg-g",
		//		// ssh
		//		//"enr:-LK4QAkFwcROm9CByx3aabpd9Muqxwj8oQeqnr7vm8PAA8l1ZbDWVZTF_bosINKhN4QVRu5eLPtyGCccRPb3yKG2xjcBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhArqAOOJc2VjcDI1NmsxoQMCphx1UQ1PkBsdOb-4FRiSWM4JE7HoDarAzOp82SO4s4N0Y3CCE4iDdWRwgg-g",
		//	},
		//	UDPPort:     udpPort,
		//	TCPPort:     tcpPort,
		//	TopicName:   validatorKey,
		//	HostDNS:     hostDNS,
		//	HostAddress: hostAddress,
		//}
		//network, err := p2p.New(cmd.Context(), logger, &cfg)
		//if err != nil {
		//	logger.Fatal("failed to create network", zap.Error(err))
		//}
		//
		//msgQ := msgqueue.New()
		//
		//ssvNode := node.New(node.Options{
		//	ValidatorStorage: validatorStorage,
		//	Beacon:           beaconClient,
		//	ETHNetwork:       eth2Network,
		//	Network:          network,
		//	Queue:            msgQ,
		//	Consensus:        consensusType,
		//	IBFT: ibft.New(
		//		ibftStorage,
		//		network,
		//		msgQ,
		//		&proto.InstanceParams{
		//			ConsensusParams: proto.DefaultConsensusParams(),
		//		},
		//	),
		//	Logger:                     logger,
		//	SignatureCollectionTimeout: sigCollectionTimeout,
		//	Phase1TestGenesis:          genesisEpoch,
		//})
		//
		//if err := ssvNode.Start(cmd.Context()); err != nil {
		//	logger.Fatal("failed to start SSV node", zap.Error(err))
		//}
	},
}

//func configureStorage(storagePath string, logger *zap.Logger, validatorPk *bls.PublicKey, shareKey *bls.SecretKey, nodeID uint64) (collections.ValidatorStorage, collections.IbftStorage) {
//
//	db, err := storage.GetStorageFactory(basedb.Options{
//		Type:   "db",
//		Path:   storagePath,
//		Logger: logger,
//	})
//
//	//db, err := kv.New(storagePath, *logger)
//	if err != nil {
//		logger.Fatal("failed to create db!", zap.Error(err))
//	}
//
//	validatorStorage := collections.NewValidator(db, logger)
//	// saves .env validator to storage
//	ibftCommittee := map[uint64]*proto.Node{
//		1: {
//			IbftId: 1,
//			Pk:     _getBytesFromHex(os.Getenv("PUBKEY_NODE_1")),
//		},
//		2: {
//			IbftId: 2,
//			Pk:     _getBytesFromHex(os.Getenv("PUBKEY_NODE_2")),
//		},
//		3: {
//			IbftId: 3,
//			Pk:     _getBytesFromHex(os.Getenv("PUBKEY_NODE_3")),
//		},
//		4: {
//			IbftId: 4,
//			Pk:     _getBytesFromHex(os.Getenv("PUBKEY_NODE_4")),
//		},
//	}
//
//	if err := validatorStorage.LoadFromConfig(nodeID, validatorPk, shareKey, ibftCommittee); err != nil {
//		logger.Error("Failed to load validator share data from config", zap.Error(err))
//	}
//	ibftStorage := collections.NewIbft(db, logger, "attestation")
//	return validatorStorage, ibftStorage
//}

func _getBytesFromHex(str string) []byte {
	val, _ := hex.DecodeString(str)
	return val
}

func init() {
	flags.AddPrivKeyFlag(StartNodeCmd)
	flags.AddValidatorKeyFlag(StartNodeCmd)
	flags.AddBeaconAddrFlag(StartNodeCmd)
	flags.AddNetworkFlag(StartNodeCmd)
	flags.AddDiscoveryFlag(StartNodeCmd)
	flags.AddConsensusFlag(StartNodeCmd)
	flags.AddNodeIDKeyFlag(StartNodeCmd)
	flags.AddSignatureCollectionTimeFlag(StartNodeCmd)
	flags.AddHostDNSFlag(StartNodeCmd)
	flags.AddHostAddressFlag(StartNodeCmd)
	flags.AddTCPPortFlag(StartNodeCmd)
	flags.AddUDPPortFlag(StartNodeCmd)
	flags.AddGenesisEpochFlag(StartNodeCmd)
	flags.AddStoragePathFlag(StartNodeCmd)

	//RootCmd.AddCommand(startNodeCmd)
}
