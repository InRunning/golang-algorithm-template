package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"gorm.io/plugin/soft_delete"
)

/*
---------- DDL ----------
CREATE TABLE `employee` (
  `id` int(11)  NOT NULL COMMENT '主键，员工虚拟帐户ID',
  `euid` bigint(11) NOT NULL COMMENT '虚拟帐户ID',
  `company_id` int(11)  NOT NULL COMMENT '外键-公司ID',
  `group_id` int(11) NOT NULL DEFAULT '-1' COMMENT '分组id',
  `uid` varchar(2058) NOT NULL DEFAULT '' COMMENT '用户牛牛号码',
  `uid_type` enum('MANUAL','AUTO') NOT NULL DEFAULT 'MANUAL' COMMENT '员工牛牛号的来源,MANUAL: 员工自己注册的牛牛号；AUTO：系统自动注册的牛牛号',
  `user_id` varchar(2058) NOT NULL DEFAULT '' COMMENT '用户公司内部工号',
  `name` varchar(2058) NOT NULL COMMENT '用户姓名',
  `phone_region` varchar(50) NOT NULL DEFAULT '' COMMENT '手机区号',
  `phone` varchar(2058) NOT NULL DEFAULT '' COMMENT '员工手机号',
  `email` varchar(2058) NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `email2` varchar(2058) NOT NULL DEFAULT '' COMMENT '用户子邮箱',
  `id_code_type` enum('ID_CARD','PASSPORT','HKM_PASS','TW_PASS','OTHER') NOT NULL DEFAULT 'ID_CARD' COMMENT '证件类型  ID_CARD:居民身份证 PASSPORT:护照 HKM_PASS:港澳居民来往内地通行证 TW_PASS:台湾居民来往大陆通行证 OTHER:其它',
  `id_code` varchar(2058) NOT NULL DEFAULT '' COMMENT '证件号码',
  `company_tax_rule_id` int(10)  NOT NULL DEFAULT '1' COMMENT '外键-自定义计税规则ID',
  `bind_status` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否完成与牛牛号绑定',
  `enable_entrance` enum('Y','N') NOT NULL DEFAULT 'Y' COMMENT 'Y:打开激励入口 N：关闭激励入口',
  `receive_msg` enum('Y','N') NOT NULL DEFAULT 'Y' COMMENT '是否接收消息通知',
  `bind_notice_status` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否发送通知绑定',
  `affiliated_person` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否关联人士',
  `has_incentive` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否有激励, 此字段并不是实时更新的，可能会有几秒钟的延时，所以对实时性要求非常高的判断，请不要用本字段',
  `job_status` enum('ON','OFF','RESIGNING') NOT NULL DEFAULT 'ON' COMMENT '工作状态：在职，离职，离职流程中',
  `last_day` date DEFAULT NULL COMMENT '离职日期',
  `data_source` varchar(20) NOT NULL DEFAULT 'ESOP' COMMENT '数据来源： ESOP 来自esop系统内  EHR 来自第三方ehr系统',
  `create_time` int(11)  NOT NULL DEFAULT '0' COMMENT '记录创建时间',
  `delete_time` int(11)  NOT NULL DEFAULT '0' COMMENT '记录删除时间',
  `update_time` int(11)  NOT NULL DEFAULT '0' COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `euid` (`euid`),
  UNIQUE KEY `company_id` (`company_id`,`id_code`(136),`delete_time`) USING BTREE,
  KEY `company_id_name` (`company_id`,`name`(136),`delete_time`) USING BTREE,
  KEY `company_id_user_id` (`company_id`,`user_id`(136),`delete_time`) USING BTREE,
  KEY `company_id_affiliated_person` (`company_id`),
  KEY `company_id_uid` (`company_id`,`uid`(136)),
  KEY `affiliated_person` (`affiliated_person`)
) ENGINE=InnoDB =227208 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='公司员工'
-------------------------
*/

type Employee struct {
	// 主键，员工虚拟帐户ID
	Id uint32
	// 虚拟帐户ID
	Euid int64
	// 外键-公司ID
	CompanyId uint32
	// 分组id
	GroupId int32
	// 用户牛牛号码
	Uid *string
	// 员工牛牛号的来源,MANUAL: 员工自己注册的牛牛号；AUTO：系统自动注册的牛牛号
	UidType string
	// 用户公司内部工号
	UserId *string
	// 用户姓名
	Name *string
	// 手机区号
	PhoneRegion *string
	// 员工手机号
	Phone *string
	// 用户邮箱
	Email *string
	// 用户子邮箱
	Email2 *string
	// 证件类型  ID_CARD:居民身份证 PASSPORT:护照 HKM_PASS:港澳居民来往内地通行证 TW_PASS:台湾居民来往大陆通行证 OTHER:其它
	IdCodeType string
	// 证件号码
	IdCode *string
	// 外键-自定义计税规则ID
	CompanyTaxRuleId uint32
	// 是否完成与牛牛号绑定
	BindStatus string
	// Y:打开激励入口 N：关闭激励入口
	EnableEntrance string
	// 是否接收消息通知
	ReceiveMsg string
	// 是否发送通知绑定
	BindNoticeStatus string
	// 是否关联人士
	AffiliatedPerson string
	// 是否有激励, 此字段并不是实时更新的，可能会有几秒钟的延时，所以对实时性要求非常高的判断，请不要用本字段
	HasIncentive string
	// 工作状态：在职，离职，离职流程中
	JobStatus string
	// 离职日期
	LastDay *sql.NullTime
	// 数据来源： ESOP 来自esop系统内  EHR 来自第三方ehr系统
	DataSource string
	// 记录创建时间
	CreateTime uint32
	// 记录删除时间
	DeleteTime soft_delete.DeletedAt
	// 记录更新时间
	UpdateTime uint32
}

func GenerateCreateTableEmployeeSQL(tableName string) (string, error) {
	s := "KAogIGBpZGAgaW50KDExKSB1bnNpZ25lZCBOT1QgTlVMTCBBVVRPX0lOQ1JFTUVOVCBDT01NRU5UICfkuLvplK7vvIzlkZjlt6XomZrmi5/luJDmiLdJRCcsCiAgYGV1aWRgIGJpZ2ludCgxMSkgTk9UIE5VTEwgQ09NTUVOVCAn6Jma5ouf5biQ5oi3SUQnLAogIGBjb21wYW55X2lkYCBpbnQoMTEpIHVuc2lnbmVkIE5PVCBOVUxMIENPTU1FTlQgJ+WklumUri3lhazlj7hJRCcsCiAgYGdyb3VwX2lkYCBpbnQoMTEpIE5PVCBOVUxMIERFRkFVTFQgJy0xJyBDT01NRU5UICfliIbnu4RpZCcsCiAgYHVpZGAgdmFyY2hhcigyMDU4KSBOT1QgTlVMTCBERUZBVUxUICcnIENPTU1FTlQgJ+eUqOaIt+eJm+eJm+WPt+eggScsCiAgYHVpZF90eXBlYCBlbnVtKCdNQU5VQUwnLCdBVVRPJykgTk9UIE5VTEwgREVGQVVMVCAnTUFOVUFMJyBDT01NRU5UICflkZjlt6XniZvniZvlj7fnmoTmnaXmupAsTUFOVUFMOiDlkZjlt6Xoh6rlt7Hms6jlhoznmoTniZvniZvlj7fvvJtBVVRP77ya57O757uf6Ieq5Yqo5rOo5YaM55qE54mb54mb5Y+3JywKICBgdXNlcl9pZGAgdmFyY2hhcigyMDU4KSBOT1QgTlVMTCBERUZBVUxUICcnIENPTU1FTlQgJ+eUqOaIt+WFrOWPuOWGhemDqOW3peWPtycsCiAgYG5hbWVgIHZhcmNoYXIoMjA1OCkgTk9UIE5VTEwgQ09NTUVOVCAn55So5oi35aeT5ZCNJywKICBgcGhvbmVfcmVnaW9uYCB2YXJjaGFyKDUwKSBOT1QgTlVMTCBERUZBVUxUICcnIENPTU1FTlQgJ+aJi+acuuWMuuWPtycsCiAgYHBob25lYCB2YXJjaGFyKDIwNTgpIE5PVCBOVUxMIERFRkFVTFQgJycgQ09NTUVOVCAn5ZGY5bel5omL5py65Y+3JywKICBgZW1haWxgIHZhcmNoYXIoMjA1OCkgTk9UIE5VTEwgREVGQVVMVCAnJyBDT01NRU5UICfnlKjmiLfpgq7nrrEnLAogIGBlbWFpbDJgIHZhcmNoYXIoMjA1OCkgTk9UIE5VTEwgREVGQVVMVCAnJyBDT01NRU5UICfnlKjmiLflrZDpgq7nrrEnLAogIGBpZF9jb2RlX3R5cGVgIGVudW0oJ0lEX0NBUkQnLCdQQVNTUE9SVCcsJ0hLTV9QQVNTJywnVFdfUEFTUycsJ09USEVSJykgTk9UIE5VTEwgREVGQVVMVCAnSURfQ0FSRCcgQ09NTUVOVCAn6K+B5Lu257G75Z6LICBJRF9DQVJEOuWxheawkei6q+S7veivgSBQQVNTUE9SVDrmiqTnhacgSEtNX1BBU1M65riv5r6z5bGF5rCR5p2l5b6A5YaF5Zyw6YCa6KGM6K+BIFRXX1BBU1M65Y+w5rm+5bGF5rCR5p2l5b6A5aSn6ZmG6YCa6KGM6K+BIE9USEVSOuWFtuWugycsCiAgYGlkX2NvZGVgIHZhcmNoYXIoMjA1OCkgTk9UIE5VTEwgREVGQVVMVCAnJyBDT01NRU5UICfor4Hku7blj7fnoIEnLAogIGBjb21wYW55X3RheF9ydWxlX2lkYCBpbnQoMTApIHVuc2lnbmVkIE5PVCBOVUxMIERFRkFVTFQgJzEnIENPTU1FTlQgJ+WklumUri3oh6rlrprkuYnorqHnqI7op4TliJlJRCcsCiAgYGJpbmRfc3RhdHVzYCBlbnVtKCdZJywnTicpIE5PVCBOVUxMIERFRkFVTFQgJ04nIENPTU1FTlQgJ+aYr+WQpuWujOaIkOS4jueJm+eJm+WPt+e7keWumicsCiAgYGVuYWJsZV9lbnRyYW5jZWAgZW51bSgnWScsJ04nKSBOT1QgTlVMTCBERUZBVUxUICdZJyBDT01NRU5UICdZOuaJk+W8gOa/gOWKseWFpeWPoyBO77ya5YWz6Zet5r+A5Yqx5YWl5Y+jJywKICBgcmVjZWl2ZV9tc2dgIGVudW0oJ1knLCdOJykgTk9UIE5VTEwgREVGQVVMVCAnWScgQ09NTUVOVCAn5piv5ZCm5o6l5pS25raI5oGv6YCa55+lJywKICBgYmluZF9ub3RpY2Vfc3RhdHVzYCBlbnVtKCdZJywnTicpIE5PVCBOVUxMIERFRkFVTFQgJ04nIENPTU1FTlQgJ+aYr+WQpuWPkemAgemAmuefpee7keWumicsCiAgYGFmZmlsaWF0ZWRfcGVyc29uYCBlbnVtKCdZJywnTicpIE5PVCBOVUxMIERFRkFVTFQgJ04nIENPTU1FTlQgJ+aYr+WQpuWFs+iBlOS6uuWjqycsCiAgYGhhc19pbmNlbnRpdmVgIGVudW0oJ1knLCdOJykgTk9UIE5VTEwgREVGQVVMVCAnTicgQ09NTUVOVCAn5piv5ZCm5pyJ5r+A5YqxLCDmraTlrZfmrrXlubbkuI3mmK/lrp7ml7bmm7TmlrDnmoTvvIzlj6/og73kvJrmnInlh6Dnp5Lpkp/nmoTlu7bml7bvvIzmiYDku6Xlr7nlrp7ml7bmgKfopoHmsYLpnZ7luLjpq5jnmoTliKTmlq3vvIzor7fkuI3opoHnlKjmnKzlrZfmrrUnLAogIGBqb2Jfc3RhdHVzYCBlbnVtKCdPTicsJ09GRicsJ1JFU0lHTklORycpIE5PVCBOVUxMIERFRkFVTFQgJ09OJyBDT01NRU5UICflt6XkvZznirbmgIHvvJrlnKjogYzvvIznprvogYzvvIznprvogYzmtYHnqIvkuK0nLAogIGBsYXN0X2RheWAgZGF0ZSBERUZBVUxUIE5VTEwgQ09NTUVOVCAn56a76IGM5pel5pyfJywKICBgZGF0YV9zb3VyY2VgIHZhcmNoYXIoMjApIE5PVCBOVUxMIERFRkFVTFQgJ0VTT1AnIENPTU1FTlQgJ+aVsOaNruadpea6kO+8miBFU09QIOadpeiHqmVzb3Dns7vnu5/lhoUgIEVIUiDmnaXoh6rnrKzkuInmlrllaHLns7vnu58nLAogIGBjcmVhdGVfdGltZWAgaW50KDExKSB1bnNpZ25lZCBOT1QgTlVMTCBERUZBVUxUICcwJyBDT01NRU5UICforrDlvZXliJvlu7rml7bpl7QnLAogIGBkZWxldGVfdGltZWAgaW50KDExKSB1bnNpZ25lZCBOT1QgTlVMTCBERUZBVUxUICcwJyBDT01NRU5UICforrDlvZXliKDpmaTml7bpl7QnLAogIGB1cGRhdGVfdGltZWAgaW50KDExKSB1bnNpZ25lZCBOT1QgTlVMTCBERUZBVUxUICcwJyBDT01NRU5UICforrDlvZXmm7TmlrDml7bpl7QnLAogIFBSSU1BUlkgS0VZIChgaWRgKSBVU0lORyBCVFJFRSwKICBVTklRVUUgS0VZIGBldWlkYCAoYGV1aWRgKSwKICBVTklRVUUgS0VZIGBjb21wYW55X2lkYCAoYGNvbXBhbnlfaWRgLGBpZF9jb2RlYCgxMzYpLGBkZWxldGVfdGltZWApIFVTSU5HIEJUUkVFLAogIEtFWSBgY29tcGFueV9pZF9uYW1lYCAoYGNvbXBhbnlfaWRgLGBuYW1lYCgxMzYpLGBkZWxldGVfdGltZWApIFVTSU5HIEJUUkVFLAogIEtFWSBgY29tcGFueV9pZF91c2VyX2lkYCAoYGNvbXBhbnlfaWRgLGB1c2VyX2lkYCgxMzYpLGBkZWxldGVfdGltZWApIFVTSU5HIEJUUkVFLAogIEtFWSBgY29tcGFueV9pZF9hZmZpbGlhdGVkX3BlcnNvbmAgKGBjb21wYW55X2lkYCksCiAgS0VZIGBjb21wYW55X2lkX3VpZGAgKGBjb21wYW55X2lkYCxgdWlkYCgxMzYpKSwKICBLRVkgYGFmZmlsaWF0ZWRfcGVyc29uYCAoYGFmZmlsaWF0ZWRfcGVyc29uYCkKKSBFTkdJTkU9SW5ub0RCICAgREVGQVVMVCBDSEFSU0VUPXV0ZjhtYjQgUk9XX0ZPUk1BVD1DT01QQUNUIENPTU1FTlQ9J+WFrOWPuOWRmOW3pSc="
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("failed to create table, parse ddl error: %v", err)
	}
	if tableName == "" {
		return "", fmt.Errorf("failed to create table, table name is empty")
	}
	createTableSQL := "CREATE TABLE `" + tableName + "` " + string(data)
	return createTableSQL, nil
}

func (e *Employee) GetCryptoFieldsMap() map[string]bool {
	// false 是不设置 padding 填充的意思
	return map[string]bool{
		"uid":     false,
		"user_id": false,
		"name":    false,
		"phone":   false,
		"email":   false,
		"email2":  false,
		"id_code": false,
	}
}
