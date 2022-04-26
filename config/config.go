package config

type Configuration struct {
	Module        string
	OutputDir     string
	Verbose       bool
	IncludeTables []string
	ExcludeTables []string
	Global        *GlobalConfiguration
	Entity        *EntityConfiguration
	MapperEnable  bool
	Config        *CfgConfiguration
	Controller    *ControllerConfiguration
	Service       *ServiceConfiguration
}

type GlobalConfiguration struct {
	Author           string
	Date             bool
	DateLayout       string
	Copyright        bool
	CopyrightContent string
	Website          bool
	WebsiteContent   string
}

type EntityConfiguration struct {
	PKG                   string
	TableToEntityStrategy StrategyType
	ColumnToFieldStrategy StrategyType
	FileNameStrategy      StrategyType
	JSONTag               bool
	JSONTagKeyStrategy    StrategyType
	FieldIdUpper          bool
	Comment               bool
	FieldComment          bool
	NamePrefix            string
	NameSuffix            string
	Orm                   bool
}

type CfgConfiguration struct {
	PKG  string
	Name string
}

type ControllerConfiguration struct {
	PKG              string
	NameStrategy     StrategyType
	VarNameStrategy  StrategyType
	RouteStrategy    StrategyType
	FileNameStrategy StrategyType
	NamePrefix       string
	NameSuffix       string
	RoutePrefix      string
	RouteSuffix      string
	VarNamePrefix    string
	VarNameSuffix    string
	Comment          bool
}

type ServiceConfiguration struct {
	PKG              string
	NameStrategy     StrategyType
	VarNameStrategy  StrategyType
	FileNameStrategy StrategyType
	NamePrefix       string
	NameSuffix       string
	VarNamePrefix    string
	VarNameSuffix    string
	Comment          bool
}
