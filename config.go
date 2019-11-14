package hades

//Config allows easy management of map config data
type Config struct {
	Data map[string]interface{}
}

//Interface returns an interface value from config
func (config Config) Interface(name string) interface{} {
	return config.Data[name]
}

//Str returns a string from config
func (config Config) Str(name string) string {
	return config.Data[name].(string)
}

//Bool returns a bool from config
func (config Config) Bool(name string) bool {
	return config.Data[name].(bool)
}

//Int returns an int from config
func (config Config) Int(name string) int {
	return config.Data[name].(int)
}

//Float returns an float64 from config
func (config Config) Float(name string) float64 {
	return config.Data[name].(float64)
}

//List returns an list from config
func (config Config) List(name string) []interface{} {
	return config.Data[name].([]interface{})
}

//Map returns a sub map from config
func (config Config) Map(name string) Config {
	return Config{
		Data: config.Data[name].(map[string]interface{}),
	}
}
