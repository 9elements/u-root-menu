package cfr

// Define CFR Tags as enum
var (
	CFR_TAG_OPTION_FORM         = uint32(1)
	CFR_TAG_ENUM_VALUE          = uint32(2)
	CFR_TAG_OPTION_ENUM         = uint32(3)
	CFR_TAG_OPTION_NUMBER       = uint32(4)
	CFR_TAG_OPTION_BOOL         = uint32(5)
	CFR_TAG_OPTION_VARCHAR      = uint32(6)
	CFR_TAG_VARCHAR_OPT_NAME    = uint32(7)
	CFR_TAG_VARCHAR_UI_NAME     = uint32(8)
	CFR_TAG_VARCHAR_UI_HELPTEXT = uint32(9)
	CFR_TAG_VARCHAR_DEF_VALUE   = uint32(10)
	CFR_TAG_OPTION_COMMENT      = uint32(11)
)

var (
	CFR_OPTFLAG_READONLY = uint32(1 << 0)
	CFR_OPTFLAG_GRAYOUT  = uint32(1 << 1)
	CFR_OPTFLAG_SUPPRESS = uint32(1 << 2)
	CFR_OPTFLAG_VOLATILE = uint32(1 << 3)
	CFR_OPTFLAG_RUNTIME  = uint32(1 << 4)
)

type LB_CFR_VARBINARY struct {
	Tag        uint32
	Size       uint32
	DataLength uint32
}

type LB_CFR_ENUM_VALUE struct {
	Tag   uint32
	Size  uint32
	Value uint32
}

type LB_CFR_NUMERIC_OPTION struct {
	Tag          uint32
	Size         uint32
	ObjectID     uint64
	DependencyID uint64
	Flags        uint32
	DefaultValue uint32
}

type LB_CFR_VARCHAR_OPTION struct {
	Tag          uint32
	Size         uint32
	ObjectID     uint64
	DependencyID uint64
	Flags        uint32
}

type LB_CFR_OPTION_COMMENT struct {
	Tag          uint32
	Size         uint32
	ObjectID     uint64
	DependencyID uint64
	Flags        uint32
}

type LB_CFR_OPTION_FORM struct {
	Tag          uint32
	Size         uint32
	ObjectID     uint64
	DependencyID uint64
	Flags        uint32
}

type LB_CFR_HEADER struct {
	Tag  uint32
	Size uint32
}

type CFRItem struct {
	Type    uint32
	Comment string
}
type CFR struct {
	Items []CFRItem
	Forms []CFR
}
