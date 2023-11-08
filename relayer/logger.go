package relayer

import (
	"log"
	"os"
	"github.com/cosmos/relayer/v2/relayer/provider"

)

func CreateLogger() (log.Logger, *os.File) {
	f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	return *log.Default(), f
}

// Your IBCHeader interface definition
type IBCHeader interface {
	Height() uint64
	ConsensusState() ibcexported.ConsensusState
	NextValidatorsHash() []byte
}

// LogIBCHeader formats and logs an IBCHeader
func LogIBCHeader(logger log.Logger, header provider.IBCHeader) {
	// Format the output
	logOutput := fmt.Sprintf(
		"Height: %d, ConsensusState: %v, NextValidatorsHash: %x",
		header.Height(),
		header.ConsensusState(),
		header.NextValidatorsHash(),
	)

	// Log the formatted output
	logger.Println(logOutput)
}

	// Example usage, assuming you have an instance of IBCHeader
	// LogIBCHeader(logger, ibcHeaderInstance)