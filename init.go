package gorma

import "github.com/raphael/goa/design"

// GormaConstructs is a map of Gorma types (currently only a StorageGroupDefinition
var GormaConstructs design.Construct

const (
	// Gorma is the constant string used as the index in the
	// goa DesignConstructs map
	Gorma = "gorma"
	// StorageGroup is the constant string used as the index in the
	// GormaConstructs map
	StorageGroup                            = "storagegroup"
	MySQL             RelationalStorageType = "mysql"
	Postgres          RelationalStorageType = "postgres"
	Boolean           FieldType             = "bool"
	Integer           FieldType             = "integer"
	BigInteger        FieldType             = "biginteger"
	AutoInteger       FieldType             = "auto_integer"
	AutoBigInteger    FieldType             = "auto_biginteger"
	Decimal           FieldType             = "decimal"
	BigDecimal        FieldType             = "bigdecimal"
	String            FieldType             = "string"
	Text              FieldType             = "text"
	UUID              FieldType             = "uuid"
	PKInteger         FieldType             = "pkinteger"
	PKBigInteger      FieldType             = "pkbiginteger"
	PKUUID            FieldType             = "pkuuid"
	Timestamp         FieldType             = "timestamp"
	NullableTimestamp FieldType             = "nulltimestamp"
)

// Init creates the necessary data structures for parsing a DSL
func Init() {
	// 	GormaConstructs = design.Design.NewConstructsSet("gorma") // later
	GormaConstructs = design.NewConstruct(Gorma)
	sg := &StorageGroupDefinition{}
	GormaConstructs[StorageGroup] = sg

}
