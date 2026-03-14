package internal

import (

	"golang.org/x/tools/go/analysis"
)

func linterZAP(pass *analysis.Pass, log *LogCall) {
	if len(log.Call.Args) <= log.MsgIndex {
		return
	}

	msgArg := log.Call.Args[log.MsgIndex]
	checkMessage(pass, msgArg)

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		arg := log.Call.Args[i]
		offset := i - (log.MsgIndex + 1)

		if offset%2 == 0 {
			checkSlogKey(pass, arg)
		}
	}

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		checkSensitiveArg(pass, log.Call.Args[i])
	}
}

