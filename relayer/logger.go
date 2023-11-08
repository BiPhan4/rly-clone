package relayer

import (
	"log"
	"os"
	"github.com/cosmos/relayer/v2/relayer/provider"
	"fmt"

)

func CreateLogger() (log.Logger, *os.File) {
	f, _ := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(f)

	return *log.Default(), f
}

// LogIBCHeader formats and logs an IBCHeader
func LogIBCHeader(logger log.Logger, header provider.IBCHeader) {
	// Format the output
	logOutput := fmt.Sprintf(
		"Height: %d\n, ConsensusState: %v\n, NextValidatorsHash: %x\n",
		header.Height(),
		header.ConsensusState(),
		header.NextValidatorsHash(),
	)

	// Log the formatted output
	logger.Println(logOutput)
}

// Example usage, assuming you have an instance of IBCHeader
// LogIBCHeader(logger, ibcHeaderInstance)