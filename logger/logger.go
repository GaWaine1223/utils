package logger

//Logger :interface for log
type Logger interface {
	//Trace :Trace interface for log
	Trace(args ...interface{})
	//Tracef :Tracef interface for log
	Tracef(format string, args ...interface{})
	//Debug :Debug interface for log
	Debug(args ...interface{})
	//Debugf :Debugf interface for log
	Debugf(format string, args ...interface{})

	//Info :Info interface for log
	Info(args ...interface{})
	//Infof :Infof interface for log
	Infof(format string, args ...interface{})
	//Warn :Warn interface for log
	Warn(args ...interface{})
	//Warnf :Warnf interface for log
	Warnf(format string, args ...interface{})
	//Error :Error interface for log
	Error(args ...interface{})
	//Errorf :Errorf interface for log
	Errorf(format string, args ...interface{})
	//Fatal :Fatal interface for log
	Fatal(args ...interface{})
	//Fatalf :Fatalf interface for log
	Fatalf(format string, args ...interface{})
	//Close :Close interface for log
	Close()
}