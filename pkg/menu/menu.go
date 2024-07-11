package menu

type BIOSOptionType int

const (
	OptionTypeString BIOSOptionType = iota
	OptionTypeInteger
	OptionTypeEnum
)

type BIOSOption struct {
	Name         string
	Type         BIOSOptionType
	StringValue  *string
	IntValue     *int
	EnumValues   []string
	HelpText     string
	DefaultValue interface{}
}

type BIOSMenu struct {
	Name     string
	HelpText string
	Options  []BIOSOption
	SubMenus []BIOSMenu
}

func New() *BIOSMenu {
	return &BIOSMenu{}
}

func (b *BIOSMenu) Categories() map[string]BIOSMenu {
	categories := make(map[string]BIOSMenu)
	for _, submenu := range b.SubMenus {
		categories[submenu.Name] = submenu
	}
	return categories
}
