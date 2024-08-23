package objectgraph

type Schema struct {
	Query map[string]*SchemaField
	Types map[string]map[string]*SchemaField
}

type SchemaField struct {
	SchemaType   string
	FieldType    *SchemaFieldType
	EngineConfig *SchemaFieldEngine
}

type SchemaFieldType struct {
	NonNull        bool
	IsArray        bool
	IsArrayNonNull bool
}

type SchemaFieldEngine struct {
	Retriever *SchemaFieldEngineRetrieve
	Object    *SchemaFieldEngineObject
}

type SchemaFieldEngineRetrieve struct {
	Key  *SchemaObjectKey
	Args map[string]string
}

type SchemaFieldEngineObject struct {
	Key *SchemaObjectKey
}

type SchemaObjectKey struct {
	TypeName string
	FieldKey string
}
