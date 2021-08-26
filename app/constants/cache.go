package constants

// CacheKey 缓存key
const (
	CachedStaffKey string = "cached_model:staff:%s"
	CachedRoleKey  string = "cached_model:role:%s"

	CacheMainStaffInfoKey       string = "cached_model:corp:%s:dept:%s:offset:%d:limit:%d"
	CacheMainStaffInfoKeyPrefix string = "cached_model:corp:%s*"

	CacheCustomerSummaryKey string = "cached_model:corp:%s:staff:%s"

	StaffIDConverterKey string = "staff_id_converter"
)

const (
	DelCacheMainStaffInfoKeyScripts string = `
	local key_list = redis.call(KEYS[1], ARGV[1])
	if #key_list > 0 then
		return (redis.call('DEL', unpack(key_list)))
	else
		return nil
	end`
)
