// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysProductVariation is the golang structure of table sys_product_variation for DAO operations like Where/Data.
type SysProductVariation struct {
	g.Meta      `orm:"table:sys_product_variation, do:true"`
	VariationId interface{} //
	ProductId   interface{} //
	NameEn      interface{} // 默认名称
	NameCn      interface{} // 默认名称
	Sku         interface{} // SKU
	Status      interface{} // 状态 0.下线 1.上线
}
