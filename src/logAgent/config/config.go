package config

type AppConfig struct{
	KafkaConfig		`ini:"kafka"`
	TailLogConfig	`ini:"tail"`
}

type KafkaConfig struct{
	Address		string	`ini:"address"`
	Topic		string	`ini:"topic"`
}

type TailLogConfig struct{
	FilePath	string	`ini:"filepath"`
}