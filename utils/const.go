package utils

// 同步状态标识
type SyncState = int64

// 同步类型
type SyncType = string

// 数据类型
type DataType = int

const (
        //工行响应结果中的 is_delete 字段：0表示未删除，1表示已删除，
        ICBC_ENABLE = iota
        ICBC_DISABLE
)
const (
        //管控state标识字段：1:未删除 2:已删除。
        ENABLE int8 = iota + 1
        DISABLE
)
const (
        // 同步状态
        STATE_DISABLE SyncState = iota
        STATE_ENABLE
)
const (
        SYNC_STATE              = "sync_state"              //同步状态标识
        SYNC_TASK_EXECUTED_TIME = "sync_task_executed_time" //定时任务执行时间
        SUB_QUEUE_NAME          = "subBackend"
        PUB_QUEUE_NAME          = "pubBackend"
)

const (
        //全量同步标识
        SYNC_ALL SyncType = "sync_all"
        //增量同步标识
        SYNC_FAST SyncType = "sync_fast"
)
const (
        SYNC_GROUP DataType = iota + 1
        SYNC_STAFF
)

// 通用配置
const (
        //分布式锁
        DISTRIBUTED_LOCK = "org_sync_lock"
        //数据初始化锁
        DATA_INIT_LOCK = "data_init_lock"
        //redis订阅频道名称-组织架构数据同步时间
        ORG_SYNC_TIME_CHANNEL = "org_sync_time"
)

// 客户个性化配置
const (
        REQ_HOST        = "https://dontes.picp.vip"
        THIRD_KEY_VALUE = "b55f32ae629e4ca3b61c2ca9f6e93403"
        //客户接口地址
        GROUP_URL = "/getOrganizationsByTime"
        STAFF_URL = "/getEmployeesByTime"
        //工行身份数据源
        ROOT_TAG      = "icbc"
        ROOT_NAME     = "中国工商银行"
        ROOT_GROUP_ID = "1"
        //etcd配置键名称
        THIRD_SLEEP_TIME = "console.third.sleep.time"
        THIRD_LIMIT      = "console.third.limit"
        THIRD_PATH       = "console.thirdImport1.path"
        THIRD_KEY        = "console.thirdImport1.key"
        THIRD_ROOT_NAME  = "console.third.root.name"
        THIRD_ROOT_TAG   = "console.third.root.tag"
        //客户免登录的唯一标识映射字段名称，工行客户是username，不同客户可能不同。
        THIRD_ID_KEY = "console.third.id.key"
        //第三方唯一标识，每个客户不同，工行使用的是 username 字段作为唯一标识。
        THIRD_ID_ICBC = "username"
)

// 同步设置类型 1:月 2:周 3:天 4:分钟
type SyncSetupType = int32

const (
        IDP_SYNC_DEFAULT SyncSetupType = iota //默认值无意义
        SYNC_TYPE_MONTH                       //按月同步
        SYNC_TYPE_WEEK                        //按周同步
        SYNC_TYPE_DAY                         //按天同步
        SYNC_TYPE_MINUTE                      //按分钟同步
)

// 同步任务类型
type SyncTaskType = int64

const (
        SyncTaskTypeCron SyncTaskType = 1 //定时任务执行的任务
        SyncTaskTypeUser SyncTaskType = 2 //用户手动执行的任务
)
