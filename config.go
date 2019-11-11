package hades

//Config allows easy management of map config data
type Config struct {
	data map[string]interface{}
}

//Interface returns an interface value from config
func (config Config) Interface(name string) interface{} {
	return config.data[name]
}

//Str returns a string from config
func (config Config) Str(name string) string {
	return config.data[name].(string)
}

//Bool returns a bool from config
func (config Config) Bool(name string) bool {
	return config.data[name].(bool)
}

//Int returns an int from config
func (config Config) Int(name string) int {
	return config.data[name].(int)
}

//Float returns an float64 from config
func (config Config) Float(name string) float64 {
	return config.data[name].(float64)
}

//List returns an list from config
func (config Config) List(name string) []interface{} {
	return config.data[name].([]interface{})
}

//Map returns a sub map from config
func (config Config) Map(name string) Config {
	return Config{
		data: config.data[name].(map[string]interface{}),
	}
}
